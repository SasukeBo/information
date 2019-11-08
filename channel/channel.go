package channel

import (
	// "container/list"
	// "github.com/astaxie/beego"
	"github.com/SasukeBo/information/models"
	// "github.com/astaxie/beego/logs"
	"golang.org/x/net/websocket"
	// "strings"
	"fmt"
)

// channel message channel，handle messages for a topic.
type channel struct {
	Subscriptions map[string]subscription
	MessageChan   chan channelMessage
}

// channelMessage
// channel subsctiption variables should have member, which key is IDName, and value is IDValue,
// channelMessage push to these subscriptions.
type channelMessage struct {
	IDName  string
	IDValue string
}

// subscription struct:
// Conn is connection object.
// Query is apollo query.
type subscription struct {
	Conn      *websocket.Conn
	Query     string
	Variables map[string]interface{}
}

// subscribe add subscription to channel
func subscribe(conn *websocket.Conn, message *apolloMessage, session string) error {
	sub := subscription{
		Conn:      conn,
		Query:     message.Payload.Query,
		Variables: message.Payload.Variables,
	}
	c := channels[message.Payload.OperationName]
	if c == nil {
		return models.Error{Message: fmt.Sprintf("There is no channel for topic %s", message.Payload.OperationName)}
	}

	c.Subscriptions[session] = sub
	return nil
}

// unsubscribe remove subscription from channel
func unsubscribe(session, topic string) error {
	c := channels[topic]
	if c == nil {
		return models.Error{Message: fmt.Sprintf("There is no channel for topic %s", topic)}
	}
	delete(c.Subscriptions, session)
	return nil
}

// registerChannel create channel
func registerChannel(topic string) {
	c := &channel{
		Subscriptions: make(map[string]subscription),
		MessageChan:   make(chan channelMessage),
	}

	channels[topic] = c
}
