package channel

import (
	"container/list"
	"github.com/SasukeBo/information/models"
	// "github.com/astaxie/beego/logs"
	"golang.org/x/net/websocket"
	"strings"
)

// SocketMessage 消息结构体
type SocketMessage struct {
	Topic   string // 消息话题
	Event   string // 消息事件
	Payload map[string]interface{}
	Socket  *websocket.Conn
	Ref     string // apollo 订阅id
}

// IChannel 接口
type IChannel interface {
	Join(message *SocketMessage)      // 加入话题
	Leave(message *SocketMessage)     // 离开话题
	HandleIn(message *SocketMessage)  // 处理消息进入 channel
	HandleOut(message *SocketMessage) // 处理消息发出 channel
	getMessageChan() chan SocketMessage
}

// Subscribe 订阅结构体
type Subscribe struct {
	ID        string          // apollo 订阅id
	Topic     string          // 订阅话题
	UserUUID  string          // 用户 uuid
	SessionID string          // 当前连接 session_id
	Socket    *websocket.Conn // 当前连接 socket
	Payload   map[string]interface{}
}

// channelType struct type
type channelType struct {
	Subscribers map[string]*list.List
	Messagechan chan SocketMessage
}

// channel router
var channelRouter = make(map[string]IChannel)

// register channel route
func channel(topic string, c IChannel) {
	channelRouter[topic] = c
	messageChan := c.getMessageChan()

	for {
		select {
		case socketMessage := <-messageChan:
			switch socketMessage.Event {
			case "start":
				c.Join(&socketMessage)
			case "stop":
				c.Leave(&socketMessage)
			case "data":
				c.HandleOut(&socketMessage)
			}
		}
	}
}

// PubSub handleSocketMessage
func PubSub(sm *SocketMessage) {
	topics := strings.Split(sm.Topic, ":")
	topic := strings.Join([]string{topics[0], "*"}, ":")
	c := channelRouter[topic]
	if c == nil {
		return
	}

	c.HandleIn(sm)
}

func getSubTopic(topic string) string {
	topics := strings.Split(topic, ":")
	if len(topics) > 1 {
		return topics[1]
	}

	return "any"
}

func unsubscribe(subs *list.List, sessionID string, user *models.User) *list.List {
	for el := subs.Front(); el != nil; el = el.Next() {
		sub := el.Value.(Subscribe)
		if sub.UserUUID == user.UUID && sub.SessionID == sessionID {
			subs.Remove(el)
		}
	}

	return subs
}
