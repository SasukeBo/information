package channel

import (
	"golang.org/x/net/websocket"

	"github.com/astaxie/beego/logs"
)

// Connect handle websocket connection
func Connect(conn *websocket.Conn) {
	userUUID := "" // TODO: 获取用户uuid
	defer conn.Close()

	for {
		var msg string
		if err := websocket.Message.Receive(conn, &msg); err == nil {
			var socketMsg SocketMsg
			err := socketMsg.ParseMsg(msg)
			if err != nil {
				// 消息解析错误，忽略此条消息
				logs.Error("socketMsg", err.Error())
				continue
			}

			switch socketMsg.Channel {
			case "device": // device channel
				handleEvent(conn, &DeviceChannel, &socketMsg, userUUID)
				defer DeviceChannel.LeaveTopic(&socketMsg, userUUID)

			case "chat":
				break
			}
		}
	}
}

func handleEvent(conn *websocket.Conn, channel Channel, socketMsg *SocketMsg, userUUID string) {
	switch socketMsg.Event {
	case "join":
		channel.JoinTopic(conn, socketMsg, userUUID)
	case "leave":
		channel.LeaveTopic(socketMsg, userUUID)
	case "data": // FIXME: 暂时用不上，从客户端接收消息，然后处理
		channel.HandleReceive(socketMsg, userUUID)
	}
}
