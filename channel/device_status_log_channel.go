package channel

import (
	"container/list"
	"encoding/json"
	"time"
	// "github.com/astaxie/beego"
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/schema"
	"github.com/SasukeBo/information/schema/scalars"
	"github.com/astaxie/beego/logs"
	"github.com/graphql-go/graphql"
	"golang.org/x/net/websocket"
	"strconv"
)

// device status log channel
type dslChannelType struct {
	channelType
}

func (dsl *dslChannelType) Join(msg *SocketMessage) {
	if err := join(dsl, msg); err != nil {
		logs.Error(err)
	}
}

func (dsl *dslChannelType) Leave(msg *SocketMessage) {
	if err := leave(dsl, msg); err != nil {
		logs.Error(err)
	}
}

func (dsl *dslChannelType) HandleIn(msg *SocketMessage) {
	dsl.Messagechan <- *msg
}

func (dsl *dslChannelType) HandleOut(msg *SocketMessage) {
	value, ok := msg.Variables["v"].(string)
	if !ok {
		logs.Error("variables value type assert string failed!")
		return
	}

	if value == "online" {
		value = "stop"
	}
	newStatus := scalars.DeviceStatusMap[value].(int)

	remoteIP, ok := msg.Variables["remoteIP"].(string)
	if !ok {
		logs.Error("variables remoteIP type assert string failed!")
		remoteIP = "unknown"
	}

	subTopic := getSubTopic(msg.Topic)
	deviceID64, err := strconv.ParseInt(subTopic, 10, 0)
	if err != nil {
		logs.Error(err)
		return
	}
	deviceID := int(deviceID64)

	device := models.Device{ID: deviceID}
	if err := device.GetBy("id"); err != nil {
		logs.Error(err)
		return
	}
	oldStatus := device.Status

	now := time.Now()
	duration := now.Sub(device.StatusChangeAt)
	device.Status = newStatus
	device.RemoteIP = remoteIP
	device.StatusChangeAt = now
	if err := device.Update("status", "remote_ip", "status_change_at"); err != nil {
		logs.Error(err)
		return
	}

	statusLog := models.DeviceStatusLog{
		Device:   &device,
		Duration: int64(duration),
		Status:   oldStatus,
	}

	if err := statusLog.Insert(); err != nil {
		logs.Error(err)
		return
	}

	subs := dsl.Subscribers[subTopic]
	if subs == nil {
		return
	}

	for el := subs.Front(); el != nil; el = el.Next() {
		sub := el.Value.(Subscribe)
		query := sub.Payload["query"].(string)
		result := graphql.Do(graphql.Params{
			Schema:         schema.PublicSchema,
			RequestString:  query,
			VariableValues: map[string]interface{}{"deviceID": int(deviceID)},
		})

		message, err := json.Marshal(map[string]interface{}{
			"type":    "data",
			"id":      sub.ID,
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
