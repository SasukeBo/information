package controllers

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego/logs"
	"golang.org/x/net/websocket"

	"github.com/SasukeBo/information/channel"
)

// Message websocket message
type Message struct {
	Type    string                `json:"type"`
	Payload channel.SocketMessage `json:"payload"`
	ID      string                `json:"id"`
}

// Connect handle websocket connection
func Connect(conn *websocket.Conn) {
	// TODO: leave defer
	defer logs.Warn("conn closed without leave topic")
	remote := getRemoteIP(conn.Request().Header.Get("X-Real-IP"))

	var msg string
	refTopic := make(map[string]string)

	if err := websocket.Message.Receive(conn, &msg); err != nil {
		logs.Error(err)
		return
	}

	var ack Message
	if err := json.Unmarshal([]byte(msg), &ack); err != nil {
		logs.Error(err)
		return
	}

	if ack.Type != "connection_init" {
		logs.Error("connection ack failed.")
		conn.Close()
		return
	}

	connectionACK, _ := json.Marshal(map[string]string{"type": "connection_ack"})
	if err := websocket.Message.Send(conn, string(connectionACK)); err != nil {
		logs.Error(err)
		return
	}

	for {
		var receive string
		var message Message
		socketMsg := make(channel.SocketMessage)
		if err := websocket.Message.Receive(conn, &receive); err != nil {
			logs.Error(err)
			return
		}

		if err := json.Unmarshal([]byte(receive), &message); err != nil {
			logs.Error(err)
			continue
		}

		if message.Payload != nil {
			socketMsg = channel.SocketMessage(message.Payload)
		}

		variables := socketMsg.GetVariables()
		topic, ok := variables["t"].(string)
		if ok && message.ID != "" && refTopic[message.ID] == "" {
			refTopic[message.ID] = topic
		}

		if message.ID == "" {
			socketMsg["topic"] = topic
		} else {
			socketMsg["topic"] = refTopic[message.ID]
		}
		socketMsg["remoteIP"] = remote
		socketMsg["refID"] = message.ID
		socketMsg["type"] = message.Type
		socketMsg["socket"] = conn

		channel.PubSub(&socketMsg)
	}
}

func getRemoteIP(remoteAddr string) string {
	addr := strings.Split(remoteAddr, ":")
	return addr[0]
}
