package channel

// 产品channel
// 参照 phoenix channel，每个 channel 都是独立的线程

import (
	"container/list"
	// "encoding/json"
	// "github.com/astaxie/beego/logs"
	// "golang.org/x/net/websocket"
	// "time"
	// "strconv"
	// "github.com/graphql-go/graphql"
	// "github.com/SasukeBo/information/models"
	// "github.com/SasukeBo/information/schema"
)

// productChan product channel
type productChan struct {
	channelType
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

func (pc *productChan) HandleOut(sm *SocketMessage) {
	return
}

/*
// HandleOut 处理消息发出
func (pc *productChan) HandleOut(sm *SocketMessage) {
	variables := sm.GetVariables()
	subTopic := sm.GetSubTopic()
	if subTopic == "" {
		logs.Error("missing subTopic")
	}

	// value, ok := variables["v"].(string)
	// if !ok {
		// logs.Error("variables value type assert string failed!")
		// return
	// }

	timeStr, ok := variables["time"].(string)
	if !ok {
		logs.Error("variables time type assert string failed!")
		return
	}

	timeValue, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		logs.Error(err)
		return
	}

	// param := &models.DeviceParam{Sign: subTopic}
	// if err := param.GetBy("sign"); err != nil {
	// logs.Error(err)
	// return
	// }

	// paramValue := &models.DeviceParamValue{
	// Value:       value,
	// DeviceParam: param,
	// CreatedAt:   timeValue,
	// }

	// if err := paramValue.Insert(); err != nil {
	// logs.Error(err)
	// return
	// }

	subs := pc.Subscribers[subTopic]
	if subs == nil {
		return
	}

	for el := subs.Front(); el != nil; el = el.Next() {
		sub := el.Value.(Subscribe)
		query := sub.Payload.GetQuery()
		if query == "" {
			continue
		}

		// result := graphql.Do(graphql.Params{
		// Schema:         schema.Root,
		// RequestString:  query,
		// VariableValues: map[string]interface{}{"id": paramValue.ID},
		// })

		message, err := json.Marshal(map[string]interface{}{
			"type": "data",
			"id":   sub.Payload.GetRefID(),
			// "payload": result,
			"payload": "",
		})

		if err != nil {
			logs.Error(err)
			continue
		}

		websocket.Message.Send(sub.Socket, string(message))
	}
}
*/

func init() {
	pc := productChan{
		channelType{
			Subscribers: make(map[string]*list.List),
			Messagechan: make(chan SocketMessage, 10),
		},
	}

	go channel("pc:*", &pc)
}
