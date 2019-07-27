package controllers

import (
	"container/list"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"time"
)

// subscribe 订阅结构体，表示一次订阅
// UUID 唯一标示订阅者
// WS 为当前订阅者建立的 websocket 连接
type subscribe struct {
	UUID  string
	Topic string
	WS    *websocket.Conn
}

// Event 消息事件
// Topic 为该消息所属话题
// Content 为消息内容
type Event struct {
	Topic   string
	Content string
}

var (
	// wsTopics map存储话题的订阅队列
	wsTopics = make(map[string]interface{})

	// wsSubChan 订阅channel
	wsSubChan = make(chan subscribe, 10)

	// wsDisSubChan 取消订阅channel
	wsDisSubChan = make(chan subscribe, 10)

	// wsEveChan 传递将要推送的消息channel
	wsEveChan = make(chan Event, 10)
)

// Send message event
func (eve Event) Send() {
	wsEveChan <- eve
}

// newTopic 新建话题
// 创建一个subscribe队列，存储在 wsTopics 中
func newTopic(topic string) {
	subs := list.New()
	wsTopics[topic] = subs
}

// join 加入话题
func join(ws *websocket.Conn, topic, UUID string) {
	wsSubChan <- subscribe{UUID: UUID, Topic: topic, WS: ws}
}

// leave 离开话题
func leave(topic, UUID string) {
	wsDisSubChan <- subscribe{UUID: UUID, Topic: topic}
}

// isTopicExist 判断话题存在性
func isTopicExist(topic string) bool {
	if wsTopics[topic] != nil {
		return true
	}
	return false
}

// isSubExist 判断订阅存在性
func isSubExist(l interface{}, uuid string) bool {
	list := l.(*list.List)
	for sub := list.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(subscribe).UUID == uuid {
			return true
		}
	}
	return false
}

// pushMessage 推送消息
func (eve Event) pushMessage() {
	subs := wsTopics[eve.Topic].(*list.List)
	for sub := subs.Front(); sub != nil; sub = sub.Next() {
		websocket.Message.Send(sub.Value.(subscribe).WS, eve.Content)
	}
}

// WsManager 处理所有的 websocket 连接
func WsManager() {
	for {
		select {
		case sub := <-wsSubChan:
			if !isTopicExist(sub.Topic) {
				websocket.Message.Send(sub.WS, "illegal topic")
				sub.WS.Close()
				break
			} else if !isSubExist(wsTopics[sub.Topic], sub.UUID) {
				subs := wsTopics[sub.Topic].(*list.List)
				subs.PushBack(sub)
			}
		case eve := <-wsEveChan:
			eve.pushMessage()
		case unsub := <-wsDisSubChan:
			subs := wsTopics[unsub.Topic].(*list.List)
			for sub := subs.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(subscribe).UUID == unsub.UUID {
					subs.Remove(sub)
					wsTopics[unsub.Topic] = subs
				}
			}
		}
	}
}

// Connect handle websocket connection
func Connect(ws *websocket.Conn) {
	defer ws.Close()
	var topic string

	for {
		if err := websocket.Message.Receive(ws, &topic); err != io.EOF {
			if !isTopicExist(topic) {
				websocket.Message.Send(ws, "illegal topic")
				return
			}

			uuid := time.Now().String()
			log.Println("join topic")
			join(ws, topic, uuid)
			defer leave(topic, uuid)
		} else {
			return
		}
	}
}

func init() {
	// 在初始化时创建话题
	newTopic("fakeData")
	go WsManager()
}
