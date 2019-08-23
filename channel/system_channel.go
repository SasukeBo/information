package channel

// system channel

import (
	"container/list"
	"golang.org/x/net/websocket"

	"github.com/astaxie/beego/logs"
)

// systemChannel 系统消息管道
type systemChannel struct {
	Subscribes    map[string]*list.List // 订阅队列
	SubChan       chan Subscribe        // 订阅添加 chan
	UnSubChan     chan Subscribe        // 订阅添加 chan
	SocketMsgChan chan SocketMsg        // 消息 chan
}

// SystemChannel 系统消息管道全局实例
var SystemChannel = systemChannel{
	Subscribes:    make(map[string]*list.List),
	SubChan:       make(chan Subscribe, 5),
	UnSubChan:     make(chan Subscribe, 5),
	SocketMsgChan: make(chan SocketMsg, 10),
}

func (sc *systemChannel) GetSubscribes() SubscribeList {
	return sc.Subscribes
}

func (sc *systemChannel) SetSubscribes(sl SubscribeList) {
	sc.Subscribes = sl
}

// hasTopic 话题是否存在
func (sc *systemChannel) hasTopic(sm *SocketMsg) bool {
	return sc.Subscribes[sm.Topic] != nil
}

func (sc *systemChannel) RegisterTopic(name string) {
	if sc.Subscribes[name] == nil {
		sc.Subscribes[name] = list.New()
	}
}

// JoinTopic 订阅话题
func (sc *systemChannel) JoinTopic(socket *websocket.Conn, sm *SocketMsg, uuid string) {
	if sc.hasTopic(sm) {
		sc.SubChan <- Subscribe{UserUUID: uuid, Topic: sm.Topic, Socket: socket}
	}
}

// LeaveTopic 取消订阅话题
func (sc *systemChannel) LeaveTopic(sm *SocketMsg, uuid string) {
	if sc.hasTopic(sm) {
		sc.UnSubChan <- Subscribe{UserUUID: uuid, Topic: sm.Topic}
	}
}

// Broadcast 广播消息
func (sc *systemChannel) Broadcast(sm SocketMsg) {
	if sc.hasTopic(&sm) {
		sc.SocketMsgChan <- sm
	}
}

// HandleData 处理消息
func (sc *systemChannel) HandleReceive(sm *SocketMsg, uuid string) {
	logs.Info("Socket handle %s data: %v, \n from user_uuid %s", sm.Topic, sm.Payload, uuid)
}

// systemChannelManager 处理system channel分发
func systemChannelManager() {
	for {
		select {
		case sub := <-SystemChannel.SubChan:
			subscribe(&SystemChannel, sub)
		case socketMsg := <-SystemChannel.SocketMsgChan:
			pushMessage(&SystemChannel, &socketMsg)
		case unsub := <-SystemChannel.UnSubChan:
			unsubscribe(&SystemChannel, unsub)
		}
	}
}

func init() {
	SystemChannel.RegisterTopic("auth")
	SystemChannel.RegisterTopic("error")

	go systemChannelManager()
}
