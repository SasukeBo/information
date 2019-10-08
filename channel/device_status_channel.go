package channel

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
	// "strconv"
)

// device status log channel
type dslChannelType struct {
	channelType
}

func (dsl *dslChannelType) Join(sm *SocketMessage) {
	join(dsl, sm)
}

func (dsl *dslChannelType) Leave(sm *SocketMessage) {
	leave(dsl, sm)
}

func (dsl *dslChannelType) HandleIn(sm *SocketMessage) {
	dsl.Messagechan <- *sm
}

func (dsl *dslChannelType) HandleOut(sm *SocketMessage) {
	o := orm.NewOrm()
	variables := sm.GetVariables()

	value, ok := variables["v"].(string)
	if !ok {
		logs.Error("variables value type assert string failed!")
		return
	}

	if value == "online" {
		value = "stop"
	}
	newStatus := schema.DeviceStatusMap[value].(int)

	remoteIP, ok := (*sm)["remoteIP"].(string)
	if !ok {
		logs.Error("variables remoteIP type assert string failed!")
		remoteIP = "unknown"
	}

	subTopic := sm.GetSubTopic()
	if subTopic == "" {
		logs.Error("missing subTopic")
		return
	}

	device := models.Device{Token: subTopic}
	if err := o.Read(&device, "token"); err != nil {
		logs.Error(err)
		return
	}

	oldStatus := device.Status

	now := time.Now()
	duration := now.Sub(device.StatusChangeAt).Truncate(time.Second)
	device.Status = newStatus
	device.RemoteIP = remoteIP
	device.StatusChangeAt = now
	if _, err := o.Update(&device, "status", "remote_ip", "status_change_at"); err != nil {
		logs.Error(err)
		return
	}

	validDuration := int(duration / 1e9)
	if validDuration <= 0 {
		return
	}

	statusLog := models.DeviceStatusLog{
		Device: &device,
		Status: oldStatus,
	}

	if _, err := o.Insert(&statusLog); err != nil {
		logs.Error(err)
		return
	}

	subs := dsl.Subscribers[subTopic]
	if subs == nil {
		return
	}

	for el := subs.Front(); el != nil; el = el.Next() {
		sub := el.Value.(Subscribe)
		query := sub.Payload.GetQuery()
		if query == "" {
			continue
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema.Root,
			RequestString:  query,
			VariableValues: map[string]interface{}{"deviceID": device.ID},
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
	dslChannel := dslChannelType{
		channelType{
			Subscribers: make(map[string]*list.List),
			Messagechan: make(chan SocketMessage, 10),
		},
	}

	go channel("dsl:*", &dslChannel)
}
