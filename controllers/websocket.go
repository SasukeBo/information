package controllers

import (
	"github.com/SasukeBo/information/models/channel"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/websocket"
	"io"
	"time"
)

// Connect handle websocket connection
func Connect(ws *websocket.Conn) {
	logs.Info(ws.Request().Cookies())
	defer ws.Close()
	var topic string

	for {
		if err := websocket.Message.Receive(ws, &topic); err != io.EOF {
			if !channel.IsTopicExist(topic) {
				websocket.Message.Send(ws, "illegal topic")
				return
			}

			uuid := time.Now().String()
			logs.Info("join topic")
			channel.Join(ws, topic, uuid)
			defer channel.Leave(topic, uuid)
		} else {
			return
		}
	}
}
