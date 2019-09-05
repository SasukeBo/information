package channel

// 设备参数值channel
// 参照 phoenix channel，每个 channel 都是独立的线程

import (
	"container/list"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/graphql-go/graphql"
	"golang.org/x/net/websocket"
	"strconv"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/schema"
)

// dpvChannelType device param value channel type
type dpvChannelType struct {
	channelType
}

// Join a topic
func (dpv *dpvChannelType) Join(msg *SocketMessage) {
	if err := join(dpv, msg); err != nil {
		logs.Error(err)
	}
}

// Leave 取消订阅
func (dpv *dpvChannelType) Leave(msg *SocketMessage) {
	if err := leave(dpv, msg); err != nil {
		logs.Error(err)
	}
}

// HandleIn 处理消息进入
func (dpv *dpvChannelType) HandleIn(msg *SocketMessage) {
	dpv.Messagechan <- *msg
}

// HandleOut 处理消息发出
func (dpv *dpvChannelType) HandleOut(msg *SocketMessage) {
	subTopic := getSubTopic(msg.Topic)
	subs := dpv.Subscribers[subTopic]
	if subs == nil {
		return
	}

	paramIDStr, ok := msg.Payload["paramID"].(string)
	if !ok {
		logs.Error("paramID type assert string failed")
		return
	}

	paramID, err := strconv.ParseInt(paramIDStr, 10, 0)
	if err != nil {
		logs.Error(err)
		return
	}

	paramValue := &models.DeviceParamValue{
		Value:       msg.Payload["value"].(string),
		DeviceParam: &models.DeviceParam{ID: int(paramID)},
	}

	if err := paramValue.Insert(); err != nil {
		logs.Error(err)
		return
	}

	for el := subs.Front(); el != nil; el = el.Next() {
		sub := el.Value.(Subscribe)
		query := sub.Payload["query"].(string)
		result := graphql.Do(graphql.Params{
			Schema:         schema.PublicSchema,
			RequestString:  query,
			VariableValues: map[string]interface{}{"id": paramValue.ID},
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
	dpvChannel := dpvChannelType{
		channelType{
			Subscribers: make(map[string]*list.List),
			Messagechan: make(chan SocketMessage, 10),
		},
	}

	go channel("device_param_value:*", &dpvChannel)
}
