package channel

import (
	"container/list"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/websocket"
	"strings"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
)

// SocketMessage 消息结构体
type SocketMessage struct {
	Topic     string // 消息话题
	Event     string // 消息事件
	Payload   map[string]interface{}
	Socket    *websocket.Conn
	Ref       string // apollo 订阅id
	Variables map[string]interface{}
}

// IChannel 接口
type IChannel interface {
	Join(message *SocketMessage)      // 加入话题
	Leave(message *SocketMessage)     // 离开话题
	HandleIn(message *SocketMessage)  // 处理消息进入 channel
	HandleOut(message *SocketMessage) // 处理消息发出 channel
	getMessageChan() chan SocketMessage
	getSubscribers(subTopic string) *list.List
	setSubscribers(subTopic string, subs *list.List)
}

// Subscribe 订阅结构体
type Subscribe struct {
	ID        string          // apollo 订阅id
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

func (ct *channelType) getMessageChan() chan SocketMessage {
	return ct.Messagechan
}

func (ct *channelType) getSubscribers(subTopic string) *list.List {
	return ct.Subscribers[subTopic]
}

func (ct *channelType) setSubscribers(subTopic string, subs *list.List) {
	ct.Subscribers[subTopic] = subs
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

	sm.Payload["id"] = topics[1]

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

func fetchSessionIDAndUserUUID(conn *websocket.Conn) (interface{}, interface{}, error) {
	session, err := conn.Request().Cookie(beego.AppConfig.String("SessionName"))
	if err != nil {
		return nil, nil, err
	}

	userLogin := models.UserLogin{SessionID: session.Value}
	if err := userLogin.GetBy("session_id"); err != nil {
		return nil, nil, err
	}

	user, err := userLogin.LoadUser()
	if err != nil {
		return nil, nil, err
	}

	return session.Value, user, nil
}

// handle join
func join(ic IChannel, msg *SocketMessage) error {
	conn := msg.Socket
	if conn == nil {
		return errors.LogicError{
			Type:    "Channel",
			Message: "Can't create subscribe without socket object!",
		}
	}

	subTopic := getSubTopic(msg.Topic)
	subs := ic.getSubscribers(subTopic)
	if subs == nil {
		subs = list.New()
	}

	sessionID, user, err := fetchSessionIDAndUserUUID(conn)
	if err != nil {
		return err
	}

	sub := Subscribe{
		UserUUID:  user.(*models.User).UUID,
		SessionID: sessionID.(string),
		Socket:    conn,
		Payload:   msg.Payload,
		ID:        msg.Ref,
	}

	subs.PushBack(sub)
	ic.setSubscribers(subTopic, subs)
	logs.Warn("user:%s join %s", user.(*models.User).Phone, msg.Topic)

	return nil
}

// handle leave
func leave(ic IChannel, msg *SocketMessage) error {
	conn := msg.Socket
	if conn == nil {
		return errors.LogicError{
			Type:    "Channel",
			Message: "Can't unsubscribe without socket object!",
		}
	}

	subTopic := getSubTopic(msg.Topic)
	sessionID, user, err := fetchSessionIDAndUserUUID(conn)
	if err != nil {
		return err
	}

	subs := ic.getSubscribers(subTopic)
	newSubs := unsubscribe(subs, sessionID.(string), user.(*models.User))
	ic.setSubscribers(subTopic, newSubs)
	logs.Warn("user:%s leave %s", user.(*models.User).Phone, msg.Topic)

	return nil
}
