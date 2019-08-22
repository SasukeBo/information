package channel

// device channel

import (
	"container/list"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"

	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/logs"
)

// deviceChannel 设备消息管道
type deviceChannel struct {
	Subscribes    map[string]*list.List // 订阅队列
	SubChan       chan Subscribe        // 订阅添加 chan
	UnSubChan     chan Subscribe        // 订阅添加 chan
	SocketMsgChan chan SocketMsg        // 消息 chan
}

// DeviceChannel 设备消息管道实例
var DeviceChannel = deviceChannel{
	Subscribes:    make(map[string]*list.List),
	SubChan:       make(chan Subscribe, 5),
	UnSubChan:     make(chan Subscribe, 5),
	SocketMsgChan: make(chan SocketMsg, 10),
}

// hasTopic 话题是否存在
func (dc *deviceChannel) hasTopic(sm *SocketMsg) bool {
	return dc.Subscribes[sm.Topic] != nil
}

// JoinTopic 订阅话题
func (dc *deviceChannel) JoinTopic(socket *websocket.Conn, sm *SocketMsg, uuid string) {
	if !dc.hasTopic(sm) {
		return
	}

	dc.SubChan <- Subscribe{
		UserUUID: uuid,
		Topic:    sm.Topic,
		Socket:   socket,
	}
}

func (dc *deviceChannel) subscribe(sub Subscribe) {
	logs.Info("join topic ", sub.Topic)
	dc.Subscribes[sub.Topic].PushBack(sub)
}

// LeaveTopic 取消订阅话题
func (dc *deviceChannel) LeaveTopic(sm *SocketMsg, uuid string) {
	if !dc.hasTopic(sm) {
		return
	}

	dc.UnSubChan <- Subscribe{
		UserUUID: uuid,
		Topic:    sm.Topic,
	}
}

func (dc *deviceChannel) unsubscribe(unsub Subscribe) {
	topic := unsub.Topic
	subs := dc.Subscribes[topic]
	for sub := subs.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscribe).UserUUID == unsub.UserUUID {
			subs.Remove(sub)
			dc.Subscribes[topic] = subs
		}
	}
}

// Broadcast 广播消息
func (dc *deviceChannel) Broadcast(sm SocketMsg) {
	if !dc.hasTopic(&sm) {
		return
	}

	dc.SocketMsgChan <- sm
}

func (dc *deviceChannel) pushMessage(sm *SocketMsg) {
	subs := dc.Subscribes[sm.Topic]
	for el := subs.Front(); el != nil; el = el.Next() {
		subscribe, ok := el.Value.(Subscribe)
		if !ok {
			return
		}

		jsonBytes, err := json.Marshal(sm)
		if err != nil {
			logs.Error("json.Marshal sm error", err.Error())
		}

		err = websocket.Message.Send(subscribe.Socket, string(jsonBytes))
		if err != nil {
			logs.Error("WebSocket Send jsonBytes error", err.Error())
		}
	}
}

// HandleData 处理消息
func (dc *deviceChannel) HandleReceive(sm *SocketMsg, uuid string) {
	logs.Info("Socket handle %s data: %v, \n from user_uuid %s", sm.Topic, sm.Payload, uuid)
}

func (dc *deviceChannel) RegisterTopic(name string) {
	if dc.Subscribes[name] == nil {
		dc.Subscribes[name] = list.New()
	}
}

// deviceChannelManager 处理device channel chan分发
func deviceChannelManager() {
	for {
		select {
		case sub := <-DeviceChannel.SubChan:
			DeviceChannel.subscribe(sub)
		case socketMsg := <-DeviceChannel.SocketMsgChan:
			DeviceChannel.pushMessage(&socketMsg)
		case unsub := <-DeviceChannel.UnSubChan:
			DeviceChannel.unsubscribe(unsub)
		}
	}
}

func init() {
	var devices []models.Device
	rawSeter := models.Repo.Raw("SELECT id FROM device")
	_, err := rawSeter.QueryRows(&devices)
	if err != nil {
		logs.Error("register device topic error ", err.Error())
	}

	for _, device := range devices {
		DeviceChannel.RegisterTopic(fmt.Sprintf("device_%d", device.ID))
	}

	go deviceChannelManager()
}
