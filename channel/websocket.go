package channel

import (
	"fmt"
	"io"
	"time"

	"golang.org/x/net/websocket"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
	"github.com/astaxie/beego/logs"
)

// Connect handle websocket connection
func Connect(conn *websocket.Conn) {
	currentUser := models.User{}
	defer conn.Close()

	msgChan := make(chan string, 5)

	go func() {
		for {
			var msg string
			err := websocket.Message.Receive(conn, &msg)
			if err == io.EOF {
				conn.Close()
				return
			} else if err != nil {
				return
			}

			msgChan <- msg
		}
	}()

	for {
		select {
		case msg := <-msgChan:
			var socketMsg SocketMsg
			err := socketMsg.ParseMsg(msg)
			if err != nil {
				// 消息解析错误，忽略此条消息
				logs.Error("socketMsg", err.Error())
				continue
			}

			switch socketMsg.Channel {
			case "system":
				if err := systemChannelHandleEvent(conn, &SystemChannel, &socketMsg, &currentUser); err != nil {
					socketError := SocketMsg{
						Channel: "system",
						Topic:   "error",
						Payload: map[string]interface{}{
							"message": fmt.Sprintf("websocket system channel:%s error", socketMsg.Topic),
							"error":   err.Error(),
						},
					}

					websocket.Message.Send(conn, socketError.Marshal())
					return
				}
				defer SystemChannel.LeaveTopic(&socketMsg, currentUser.UUID)

			case "device": // device channel
				if err := deviceChannelHandleEvent(conn, &DeviceChannel, &socketMsg, &currentUser); err != nil {
					socketError := SocketMsg{
						Channel: "system",
						Topic:   "error",
						Payload: map[string]interface{}{
							"message": fmt.Sprintf("websocket device channel:%s error", socketMsg.Topic),
							"error":   err.Error(),
						},
					}

					websocket.Message.Send(conn, socketError.Marshal())
					return
				}
				defer DeviceChannel.LeaveTopic(&socketMsg, currentUser.UUID)

			case "heartbeat":
				websocket.Message.Send(conn, `{"channel": "heartbeat", "payload": {"message": "pang"}}`)
			}
		case <-time.After(31 * time.Second): // 31秒心跳超时时间
			return
		}
	}
}

func systemChannelHandleEvent(conn *websocket.Conn, channel Channel, socketMsg *SocketMsg, user *models.User) error {
	switch socketMsg.Event {
	case "join":
		channel.JoinTopic(conn, socketMsg, user.UUID)
	case "leave":
		channel.LeaveTopic(socketMsg, user.UUID)
	case "data":
		if socketMsg.Topic == "auth" {
			if uuid, ok := socketMsg.Payload["user_uuid"].(string); ok {
				user.UUID = uuid
				if err := user.GetBy("uuid"); err != nil {
					return err
				}

				return nil
			}

			return errors.LogicError{Type: "WebSocket", Field: "user_uuid", Message: "Invalid user_uuid."}
		}
		channel.HandleReceive(socketMsg, user.UUID)
	}

	return nil
}

func deviceChannelHandleEvent(conn *websocket.Conn, channel Channel, socketMsg *SocketMsg, user *models.User) error {
	// TODO: 权限验证
	switch socketMsg.Event {
	case "join":
		channel.JoinTopic(conn, socketMsg, user.UUID)
	case "leave":
		channel.LeaveTopic(socketMsg, user.UUID)
	case "data": // FIXME: 暂时用不上，从客户端接收消息，然后处理
		channel.HandleReceive(socketMsg, user.UUID)
	}

	return nil
}
