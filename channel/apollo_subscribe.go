package channel

import (
	"encoding/json"
	// "fmt"
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/websocket"
	"io"
	// "strings"
)

var (
	channels = make(map[string]*channel)
)

// apolloMessage apollo subscription message struct
type apolloMessage struct {
	Type    string        `json:"type,omitempty"`
	Payload apolloPayload `json:"payload,omitempty"`
	ID      string        `json:"id,omitempty"`
}

// apolloPayload apollo subscription message payload struct
type apolloPayload struct {
	Variables     map[string]interface{} `json:"variables,omitempty"`
	Extensions    map[string]string      `json:"extensions,omitempty"`
	OperationName string                 `json:"operationName,omitempty"`
	Query         string                 `json:"query,omitempty"`
}

// sessionSubscription store subscription id during one connection.
// Key is subscription id, value is topic(operationName).
// All subscriptions should be removed before connection closed.
type sessionSubscription map[string]string

// removeAll will remove all subscriptions in this connection.
func (s *sessionSubscription) removeAll() {
	// TODO: finifsh it
}

// add subscription for this connection
// replace if subscription already exist for this topic(operationName)
func (s *sessionSubscription) add(message *apolloMessage, conn *websocket.Conn) error {
	if message.Payload.OperationName == "" {
		return models.Error{Message: "Fail to subscribe: topic(operationName) cannnot be empty!"}
	}

	if ok, id := s.hasTopic(message); ok {
		logs.Info("already subscribe this topic, id is", id)
		// TODO: unsubscribe
	}

	(*s)[message.ID] = message.Payload.OperationName
	// TODO: subscribe

	// Subscribe(conn, message)
	return nil
}

func (s *sessionSubscription) remove(message *apolloMessage, conn *websocket.Conn) error {
	topic := (*s)[message.ID]

	// TODO: unsubscribe
	delete((*s), message.ID)
	logs.Info("unsubscribe topic:", topic)

	return nil
}

// return true and id if subscriptions contain this topic
func (s *sessionSubscription) hasTopic(message *apolloMessage) (bool, string) {
	for id, topic := range *s {
		if topic == message.Payload.OperationName {
			return true, id
		}
	}

	return false, ""
}

// recv message from socket connection
// parse it into a apolloMessage struct
// return error if failed.
func recv(conn *websocket.Conn, m *apolloMessage) error {
	var bytes []byte
	if err := websocket.Message.Receive(conn, &bytes); err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, m); err != nil {
		return err
	}

	return nil
}

// ack handle apollo subscription ack process
// only if connection_init received, it send connection_ack back.
// otherwise return an error
func ack(conn *websocket.Conn) error {
	var message apolloMessage
	if err := recv(conn, &message); err != nil {
		return err
	}
	if message.Type != "connection_init" {
		return models.Error{Message: "ack failed."}
	}
	response := apolloMessage{Type: "connection_ack"}
	bytes, err := json.Marshal(response)
	if err != nil {
		return err
	}
	if err := websocket.Message.Send(conn, string(bytes)); err != nil {
		return err
	}
	return nil
}

// ApolloSubscribe handle apollo graphql subscribe connection
func ApolloSubscribe(conn *websocket.Conn) {
	if err := ack(conn); err != nil {
		logs.Error(err)
		return
	}

	var subscriptions = make(sessionSubscription)
	defer subscriptions.removeAll() // 连接释放时断开所有话题订阅

	for {
		var message apolloMessage
		if err := recv(conn, &message); err == io.EOF {
			conn.Close()
			return
		} else if err != nil {
			logs.Error(err)
			continue
		}

		switch message.Type {
		case "start":
			if err := subscriptions.add(&message, conn); err != nil {
				logs.Error(err)
			} // 订阅话题
		case "stop":
			subscriptions.remove(&message, conn) // 取消订阅
		}
	}
}

func init() {
	registerChannel("deviceStatus")
	registerChannel("productIns")

	go func() {
		for {
			select {
			case m := <-channels["deviceStatus"].MessageChan:
				logs.Info("Message incoming", m.IDName, m.IDValue)
			case m := <-channels["productIns"].MessageChan:
				logs.Info("Message incoming", m.IDName, m.IDValue)
			}
		}
	}()
}
