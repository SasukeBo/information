package channel

// device channel

import (
	"container/list"
	"fmt"
	"golang.org/x/net/websocket"

	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/logs"
)

// deviceChannel 设备消息管道
type deviceChannel struct {
	Subscribes    SubscribeList  // 订阅队列
	SubChan       chan Subscribe // 订阅添加 chan
	UnSubChan     chan Subscribe // 订阅添加 chan
	SocketMsgChan chan SocketMsg // 消息 chan
}

// DeviceChannel 设备消息管道实例
var DeviceChannel = deviceChannel{
	Subscribes:    make(SubscribeList),
	SubChan:       make(chan Subscribe, 5),
	UnSubChan:     make(chan Subscribe, 5),
	SocketMsgChan: make(chan SocketMsg, 10),
}

func (dc *deviceChannel) RegisterTopic(name string) {
	if dc.Subscribes[name] == nil {
		dc.Subscribes[name] = list.New()
	}
}

// hasTopic 话题是否存在
func (dc *deviceChannel) hasTopic(sm *SocketMsg) bool {
	return dc.Subscribes[sm.Topic] != nil
}

// JoinTopic 订阅话题
func (dc *deviceChannel) JoinTopic(socket *websocket.Conn, sm *SocketMsg, uuid string) {
	if dc.hasTopic(sm) {
		dc.SubChan <- Subscribe{UserUUID: uuid, Topic: sm.Topic, Socket: socket}
	}
}

// LeaveTopic 取消订阅话题
func (dc *deviceChannel) LeaveTopic(sm *SocketMsg, uuid string) {
	if dc.hasTopic(sm) {
		dc.UnSubChan <- Subscribe{UserUUID: uuid, Topic: sm.Topic}
	}
}

func (dc *deviceChannel) GetSubscribes() SubscribeList {
	return dc.Subscribes
}

func (dc *deviceChannel) SetSubscribes(sl SubscribeList) {
	dc.Subscribes = sl
}

// Broadcast 广播消息
func (dc *deviceChannel) Broadcast(sm SocketMsg) {
	if dc.hasTopic(&sm) {
		logs.Warn("send to socketMsgChan: ", sm.Payload)
		dc.SocketMsgChan <- sm
	}
}

// HandleData 处理消息
func (dc *deviceChannel) HandleReceive(sm *SocketMsg, uuid string) {
	logs.Info("Socket handle %s data: %v, \n from user_uuid %s", sm.Topic, sm.Payload, uuid)
}

// deviceChannelManager 处理device channel chan分发
func deviceChannelManager() {
	for {
		select {
		case sub := <-DeviceChannel.SubChan:
			subscribe(&DeviceChannel, sub)
		case socketMsg := <-DeviceChannel.SocketMsgChan:
			pushMessage(&DeviceChannel, &socketMsg)
		case unsub := <-DeviceChannel.UnSubChan:
			unsubscribe(&DeviceChannel, unsub)
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
