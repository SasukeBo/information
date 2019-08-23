package channel

import (
	"container/list"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"

	"github.com/SasukeBo/information/models/errors"
	// "github.com/astaxie/beego/logs"
)

// SubscribeList 订阅队列类型
type SubscribeList map[string]*list.List

// SocketMsg Socket message struct
type SocketMsg struct {
	Channel string                 `json:"channel"` // socket channel
	Event   string                 `json:"event"`   // 消息类型
	Topic   string                 `json:"topic"`   // socket topic
	Payload map[string]interface{} `json:"payload"` // 消息负载
}

// ParseMsg 解析 socket 消息
func (socketMsg *SocketMsg) ParseMsg(websocketMessage string) error {
	if err := json.Unmarshal([]byte(websocketMessage), socketMsg); err != nil {
		return errors.LogicError{
			Type:    "Controller",
			Message: "parse websocket message error",
			OriErr:  err,
		}
	}

	return nil
}

// Marshal socket 消息 to json
func (socketMsg *SocketMsg) Marshal() string {
	socketMsgB, err := json.Marshal(socketMsg)
	if err != nil {
		return fmt.Sprintf(`{"channel": "system", "topic": "error", "payload": {"error": %s, "type": "json.Marshal Error"}}`, err.Error())
	}

	return string(socketMsgB)
}

// Subscribe 订阅结构体，表示一次订阅
type Subscribe struct {
	UserUUID string          // 订阅者 uuid
	Topic    string          // 当前订阅的话题
	Socket   *websocket.Conn // 当前订阅者建立的 websocket 连接
}

// Channel 消息管道接口
type Channel interface {
	JoinTopic(conn *websocket.Conn, sm *SocketMsg, uuid string) // 订阅话题
	LeaveTopic(sm *SocketMsg, uuid string)                      // 取消订阅话题
	HandleReceive(sm *SocketMsg, uuid string)                   // 处理接收的消息
	Broadcast(sm SocketMsg)                                     // 广播消息
	hasTopic(sm *SocketMsg) bool                                // 话题是否存在
	GetSubscribes() SubscribeList
	SetSubscribes(sl SubscribeList)
}

func subscribe(ch Channel, sub Subscribe) {
	topic := sub.Topic
	subscribes := ch.GetSubscribes()
	for sub := subscribes[topic].Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscribe).Topic == topic {
			return
		}
	}
	subscribes[topic].PushBack(sub)
	ch.SetSubscribes(subscribes)
}

func unsubscribe(ch Channel, unsub Subscribe) {
	topic := unsub.Topic
	subscribes := ch.GetSubscribes()
	for sub := subscribes[topic].Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscribe).UserUUID == unsub.UserUUID {
			subscribes[topic].Remove(sub)
		}
	}
	ch.SetSubscribes(subscribes)
}

func pushMessage(ch Channel, sm *SocketMsg) {
	subs := ch.GetSubscribes()[sm.Topic]
	for item := subs.Front(); item != nil; item = item.Next() {
		if sub, ok := item.Value.(Subscribe); ok {
			websocket.Message.Send(sub.Socket, sm.Marshal())
		}
	}
}
