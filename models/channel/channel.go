package channel

import (
  "container/list"
  "golang.org/x/net/websocket"
)

// Subscribe 订阅结构体，表示一次订阅
// UserUUID 订阅者 uuid
// WS 为当前订阅者建立的 websocket 连接
type Subscribe struct {
  UserUUID string
  Topic    string
  WS       *websocket.Conn
}

// Event 消息事件
// Topic 为该消息所属话题
// Content 为消息内容
type Event struct {
  Topic   string
  Content string
}

var (
  // wsTopics 使用 map 存储话题的订阅队列
  wsTopics = make(map[string]interface{})

  // wsSubChan 订阅 channel
  wsSubChan = make(chan Subscribe, 10)

  // wsDisSubChan 取消订阅 channel
  wsDisSubChan = make(chan Subscribe, 10)

  // wsEveChan 传递将要推送的消息 channel
  wsEveChan = make(chan Event, 10)
)

// newTopic 注册话题
// 创建一个队列，用于存储 Subscribes
func newTopic(topic string) {
  subs := list.New()
  wsTopics[topic] = subs
}

// IsTopicExist 判断话题是否存在
func IsTopicExist(topic string) bool {
  if wsTopics[topic] != nil {
    return true
  }
  return false
}

// IsSubExist 判断订阅是否存在
func IsSubExist(l interface{}, userUUID string) bool {
  list := l.(*list.List)
  for sub := list.Front(); sub != nil; sub = sub.Next() {
    if sub.Value.(Subscribe).UserUUID == userUUID {
      return true
    }
  }
  return false
}

// pushMessage 推送消息
func (eve Event) pushMessage() {
  subs := wsTopics[eve.Topic].(*list.List)
  for sub := subs.Front(); sub != nil; sub = sub.Next() {
    websocket.Message.Send(sub.Value.(Subscribe).WS, eve.Content)
  }
}

// Send message event
func (eve Event) Send() {
  wsEveChan <- eve
}

// Join 加入话题
func Join(ws *websocket.Conn, topic, userUUID string) {
  wsSubChan <- Subscribe{UserUUID: userUUID, Topic: topic, WS: ws}
}

// Leave 离开话题
func Leave(topic, userUUID string) {
  wsDisSubChan <- Subscribe{UserUUID: userUUID, Topic: topic}
}

// WsManager 处理所有的 websocket 连接
func WsManager() {
  for {
    select {
    case sub := <-wsSubChan:
      if !IsTopicExist(sub.Topic) {
        websocket.Message.Send(sub.WS, "illegal topic")
        sub.WS.Close()
        break
      } else if !IsSubExist(wsTopics[sub.Topic], sub.UserUUID) {
        subs := wsTopics[sub.Topic].(*list.List)
        subs.PushBack(sub)
      }
    case eve := <-wsEveChan:
      eve.pushMessage()
    case unsub := <-wsDisSubChan:
      subs := wsTopics[unsub.Topic].(*list.List)
      for sub := subs.Front(); sub != nil; sub = sub.Next() {
        if sub.Value.(Subscribe).UserUUID == unsub.UserUUID {
          subs.Remove(sub)
          wsTopics[unsub.Topic] = subs
        }
      }
    }
  }
}

func init() {
  // 在初始化时创建话题
  newTopic("fakeData")
  go WsManager()
}
