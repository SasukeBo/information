package channel

// 产品channel
// 参照 phoenix channel，每个 channel 都是独立的线程

import (
	"container/list"
	"encoding/json"
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/schema"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
	"golang.org/x/net/websocket"
	"time"
)

// productChan product channel
type productChan struct {
	channelType
}

type detectItem struct {
	Sign  string
	Value float64
}

// Join a topic
func (pc *productChan) Join(msg *SocketMessage) {
	join(pc, msg)
}

// Leave 取消订阅
func (pc *productChan) Leave(sm *SocketMessage) {
	leave(pc, sm)
}

// HandleIn 处理消息进入
func (pc *productChan) HandleIn(sm *SocketMessage) {
	pc.Messagechan <- *sm
}

// HandleOut 处理消息发出
func (pc *productChan) HandleOut(sm *SocketMessage) {
	logs.Warn(sm)
	variables := sm.GetVariables()
	subTopic := sm.GetSubTopic()
	if subTopic == "" {
		logs.Error("missing subTopic")
	}

	jsonb, ok := variables["v"].(string)
	if !ok {
		logs.Error("get value failed.")
		return
	}

	token, ok := variables["d"].(string)
	if !ok {
		logs.Error("get token failed.")
		return
	}

	timeStr, ok := variables["time"].(string)
	if !ok {
		logs.Error("get time failed.")
		return
	}

	var items []detectItem
	if err := json.Unmarshal([]byte(jsonb), &items); err != nil {
		logs.Error("unmarshal value failed")
		return
	}

	timeValue, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		logs.Error(err)
		return
	}

	o := orm.NewOrm()
	o.Begin()

	device := models.Device{Token: token}
	if err := o.Read(&device, "token"); err != nil {
		o.Rollback()
		return
	}

	product := models.Product{Name: subTopic}
	if err := o.Read(&product, "name"); err != nil {
		o.Rollback()
		return
	}

	deviceProductShip := models.DeviceProductShip{Device: &device, Product: &product}
	if err := o.Read(&deviceProductShip, "device_id", "product_id"); err != nil {
		o.Rollback()
		return
	}

	productIns := models.ProductIns{DeviceProductShip: &deviceProductShip, CreatedAt: timeValue}
	if _, err := o.Insert(&productIns); err != nil {
		o.Rollback()
		return
	}

	for _, item := range items {
		detectItem := models.DetectItem{Sign: item.Sign}
		if err := o.Read(&detectItem, "sign"); err != nil {
			o.Rollback()
			return
		}

		detectItemValue := models.DetectItemValue{DetectItem: &detectItem, ProductIns: &productIns, Value: item.Value}
		if _, err := o.Insert(&detectItemValue); err != nil {
			o.Rollback()
			return
		}
	}

	o.Commit()

	subs := pc.Subscribers[subTopic]
	if subs == nil {
		return
	}

	for el := subs.Front(); el != nil; el = el.Next() {
		sub := el.Value.(Subscribe)
		subToken, ok := sub.Payload.GetVariables()["d"].(string)
		if !ok || subToken != token {
			continue
		}

		query := sub.Payload.GetQuery()
		if query == "" {
			continue
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema.Root,
			RequestString:  query,
			VariableValues: map[string]interface{}{"id": productIns.ID},
		})

		message, err := json.Marshal(map[string]interface{}{
			"type":    "data",
			"id":      sub.Payload.GetRefID(),
			"payload": result,
		})

		if err != nil {
			logs.Error(err)
			continue
		}

		websocket.Message.Send(sub.Socket, string(message))
	}
}

func init() {
	pc := productChan{
		channelType{
			Subscribers: make(map[string]*list.List),
			Messagechan: make(chan SocketMessage, 10),
		},
	}

	go channel("pc:*", &pc)
}
