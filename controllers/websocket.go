package controllers

import (
	"encoding/json"
	"strings"
	// "net"

	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/websocket"

	"github.com/SasukeBo/information/channel"
	// "github.com/SasukeBo/information/models"
)

type apolloWSMessage struct {
	Event   string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
	ID      string                 `json:"id"`
}

// Connect handle websocket connection
func Connect(conn *websocket.Conn) {
	remote := getRemoteIP(conn.Request().Header.Get("X-Real-IP"))

	// TODO: leave defer
	defer logs.Warn("conn closed without leave topic")
	var msg string
	refTopic := make(map[string]map[string]interface{})

	if err := websocket.Message.Receive(conn, &msg); err != nil {
		logs.Error(err)
		return
	}

	var initMsg apolloWSMessage
	if err := json.Unmarshal([]byte(msg), &initMsg); err != nil {
		logs.Error(err)
		return
	}

	if initMsg.Event != "connection_init" {
		logs.Error("not connection_init")
		return
	}

	connectionACK, err := json.Marshal(map[string]string{"type": "connection_ack"})

	if err != nil {
		logs.Error(err)
		return
	}

	err = websocket.Message.Send(conn, string(connectionACK))
	if err != nil {
		logs.Error(err)
		return
	}

	for {
		var msg string
		var data apolloWSMessage
		if err := websocket.Message.Receive(conn, &msg); err != nil {
			logs.Error(err)
			return
		}

		if err := json.Unmarshal([]byte(msg), &data); err != nil {
			logs.Error(err)
			continue
		}

		// 由于 spollo ws stop event 不携带 topic 信息，
		// 需要存储当前 conn id 对应的话题信息，否则无法定位 subscribe 并取消订阅。
		if data.Event == "stop" {
			data.Payload = refTopic[data.ID]
		} else {
			refTopic[data.ID] = data.Payload
		}

		variables, ok := data.Payload["variables"].(map[string]interface{})
		if !ok {
			logs.Error("payload variables type assert map[string]interface{} failed!")
			continue
		}

		topic, ok := variables["t"].(string)
		if !ok {
			logs.Error("variables topic type assert string failed!")
			continue
		}

		variables["remoteIP"] = remote

		socketMessage := channel.SocketMessage{
			Topic:     topic,
			Event:     data.Event,
			Payload:   data.Payload,
			Ref:       data.ID,
			Socket:    conn,
			Variables: variables,
		}

		channel.PubSub(&socketMessage)
	}
}

func getRemoteIP(remoteAddr string) string {
	addr := strings.Split(remoteAddr, ":")
	return addr[0]
}
