package controllers

import (
	"encoding/json"

	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/websocket"
)

type apolloWSMessage struct {
	Type    string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
	ID      string                 `json:"id"`
}

type topicMap map[string]struct {
	OprationName string
	ID           int
}

// NewConnect handle websocket connection
func NewConnect(conn *websocket.Conn) {
	sessionID, err := conn.Request().Cookie(beego.AppConfig.String("SessionName"))
	if err != nil {
		logs.Error(err)
		return
	}

	userLogin := models.UserLogin{SessionID: sessionID.Value}
	if err := userLogin.GetBy("session_id"); err != nil {
		logs.Error(err)
		return
	}

	user, err := userLogin.LoadUser()
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Info("ws for user uuid: ", user.UUID)

	var msg string
	if err := websocket.Message.Receive(conn, &msg); err != nil {
		logs.Error(err)
		return
	}

	var initMsg apolloWSMessage
	if err := json.Unmarshal([]byte(msg), &initMsg); err != nil {
		logs.Error(err)
		return
	}

	if initMsg.Type != "connection_init" {
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

	// TODO: 完成 websocket
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

		switch data.Type {
		case "start":
			logs.Warn(data.Payload)
		case "stop":
			logs.Warn(data)
		}
	}
}
