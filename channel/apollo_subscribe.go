package channel

import (
	"encoding/json"
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/websocket"
	"io"
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

// connSubscription store subscription id during one connection.
// Key is subscription id, value is topic(operationName).
// All subscriptions should be removed before connection closed.
type connSubscriptions map[string]string

// removeAll will remove all subscriptions in this connection.
func (s *connSubscriptions) removeAll(conn *websocket.Conn, token string) {
	for _, topic := range *s {
		unsubscribe(token, topic)
	}
}

// add subscription for this connection
// replace if subscription already exist for this topic(operationName)
func (s *connSubscriptions) add(message *apolloMessage, conn *websocket.Conn, token string) error {
	if message.Payload.OperationName == "" {
		return models.Error{Message: "Fail to subscribe: topic(operationName) cannnot be empty!"}
	}

	if ok, _ := s.hasTopic(message); ok {
		if err := unsubscribe(token, message.Payload.OperationName); err != nil {
			return err
		}
	}

	(*s)[message.ID] = message.Payload.OperationName

	if err := subscribe(conn, message, token); err != nil {
		return err
	}
	return nil
}

// remove subscription from this connection
// need message object, conn object
func (s *connSubscriptions) remove(message *apolloMessage, conn *websocket.Conn, token string) error {
	topic := (*s)[message.ID]

	if err := unsubscribe(token, topic); err != nil {
		return err
	}

	delete((*s), message.ID)
	return nil
}

// return true and id if subscriptions contain this topic
func (s *connSubscriptions) hasTopic(message *apolloMessage) (bool, string) {
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

	var subscriptions = make(connSubscriptions)
	var token = utils.GenRandomToken(8)
	defer subscriptions.removeAll(conn, token) // 连接释放时断开所有话题订阅

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
			err := subscriptions.add(&message, conn, token)
			if err != nil {
				logs.Error(err)
			}
		case "stop":
			subscriptions.remove(&message, conn, token) // 取消订阅
		case "data":
			logs.Info("data type message handler not implemented!")
		}
	}
}

func init() {
	registerChannel("deviceStatus") // 注册话题
	registerChannel("productIns")   // 注册话题

	go func() {
		for {
			select {
			case m := <-channels["deviceStatus"].MessageChan:
				BroadcastDeviceStatus(m)
			case m := <-channels["productIns"].MessageChan:
				BroadcastProductIns(m)
			}
		}
	}()
}
