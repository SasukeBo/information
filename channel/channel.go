package channel

import (
	"container/list"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/websocket"
	"strings"
	// "github.com/SasukeBo/information/models"
	// "github.com/SasukeBo/information/models/errors"
)

// SocketMessage socket message map
type SocketMessage map[string]interface{}

// GetSessionID _
func (sm *SocketMessage) GetSessionID() string {
	socket := sm.GetSocket()
	if socket == nil {
		return ""
	}

	session, err := socket.Request().Cookie(beego.AppConfig.String("SessionName"))
	if err != nil {
		return ""
	}

	return session.Value
}

// GetTopic _
func (sm *SocketMessage) GetTopic() string {
	topics := sm.getTopics()
	if len(topics) > 0 {
		return topics[0]
	}

	return ""
}

// GetSubTopic _
func (sm *SocketMessage) GetSubTopic() string {
	topics := sm.getTopics()
	if len(topics) > 1 {
		return topics[1]
	}

	return ""
}

func (sm *SocketMessage) getTopics() []string {
	topic, ok := (*sm)["topic"].(string)
	if ok {
		topics := strings.Split(topic, ":")
		return topics
	}

	return []string{}
}

// GetType _
func (sm *SocketMessage) GetType() string {
	messageType, ok := (*sm)["type"].(string)
	if ok {
		return messageType
	}

	return ""
}

// GetSocket _
func (sm *SocketMessage) GetSocket() *websocket.Conn {
	socket, ok := (*sm)["socket"].(*websocket.Conn)
	if ok {
		return socket
	}

	return nil
}

// GetRefID _
func (sm *SocketMessage) GetRefID() string {
	refID, ok := (*sm)["refID"].(string)
	if ok {
		return refID
	}

	return ""
}

// GetQuery _
func (sm *SocketMessage) GetQuery() string {
	query, ok := (*sm)["query"].(string)
	if ok {
		return query
	}

	return ""
}

// GetVariables _
func (sm *SocketMessage) GetVariables() map[string]interface{} {
	variables, ok := (*sm)["variables"].(map[string]interface{})
	if ok {
		return variables
	}

	return map[string]interface{}{}
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
	// ID        string          // apollo 订阅id
	// UserUUID  string          // 用户 uuid
	SessionID string          // 当前连接 session_id
	Socket    *websocket.Conn // 当前连接 socket
	Payload   *SocketMessage
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
		case sm := <-messageChan:
			switch sm.GetType() {
			case "start":
				c.Join(&sm)
			case "stop":
				c.Leave(&sm)
			case "data":
				c.HandleOut(&sm)
			}
		}
	}
}

// PubSub handleSocketMessage
func PubSub(sm *SocketMessage) {
	route := strings.Join([]string{sm.GetTopic(), "*"}, ":")
	c := channelRouter[route]
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

func unsubscribe(subs *list.List, sessionID string) *list.List {
	for el := subs.Front(); el != nil; el = el.Next() {
		sub := el.Value.(Subscribe)
		if sub.SessionID == sessionID {
			subs.Remove(el)
		}
	}

	return subs
}

// handle join
func join(ic IChannel, sm *SocketMessage) {
	socket := sm.GetSocket()
	if socket == nil {
		logs.Error("missing socket")
	}

	subTopic := sm.GetSubTopic()
	if subTopic == "" {
		logs.Error("missing subTopic")
		return
	}

	subs := ic.getSubscribers(subTopic)
	if subs == nil {
		subs = list.New()
	}

	sessionID := sm.GetSessionID()
	if sessionID == "" {
		logs.Error("missing sessionID")
		return
	}

	sub := Subscribe{
		SessionID: sessionID,
		Socket:    socket,
		Payload:   sm,
	}

	subs.PushBack(sub)
	ic.setSubscribers(subTopic, subs)
	// logs.Warn("user:%s join %s", user.(*models.User).Phone, msg.Topic)
}

// handle leave
func leave(ic IChannel, sm *SocketMessage) {
	socket := sm.GetSocket()
	if socket == nil {
		logs.Error("missing socket")
		return
	}

	subTopic := sm.GetSubTopic()
	sessionID := sm.GetSessionID()
	if sessionID == "" {
		logs.Error("missing sessionID")
		return
	}

	subs := ic.getSubscribers(subTopic)
	newSubs := unsubscribe(subs, sessionID)
	ic.setSubscribers(subTopic, newSubs)
	// logs.Warn("user:%s leave %s", user.(*models.User).Phone, msg.Topic)
}
