package channel

import (
	"encoding/json"
	"golang.org/x/net/websocket"

	"github.com/SasukeBo/information/models/errors"
	// "github.com/astaxie/beego/logs"
)

// SocketMsg Socket message struct
type SocketMsg struct {
	Event   string                 `json:"event"`   // 消息类型
	Channel string                 `json:"channel"` // socket channel
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
}
