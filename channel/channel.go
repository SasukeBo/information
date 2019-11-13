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

// channel is channel's data payload.
type channel struct {
	Subscriptions map[string]subscription
	MessageChan   chan Message
}

// Message channel subsctiption variables should have member, which key is IDName, and value is IDValue,
// Message push to these subscriptions.
// FIXME: this struct may not fit chat room data, rewrite it.
type Message struct {
	IDName  string
	IDValue int
	Payload map[string]interface{}
}

// subscription struct:
// Conn is connection object.
// Query is apollo query.
type subscription struct {
	RefID     string
	Conn      *websocket.Conn
	Query     string
	Variables map[string]interface{}
}

// subscribe add subscription to channel
func subscribe(conn *websocket.Conn, message *apolloMessage, session string) error {
	sub := subscription{
		RefID:     message.ID,
		Conn:      conn,
		Query:     message.Payload.Query,
		Variables: message.Payload.Variables,
	}
	c, err := getChannel(message.Payload.OperationName)
	if err != nil {
		return err
	}
	c.Subscriptions[session] = sub
	return nil
}

// unsubscribe remove subscription from channel
func unsubscribe(session, topic string) error {
	c, err := getChannel(topic)
	if err != nil {
		return err
	}
	delete(c.Subscriptions, session)
	return nil
}

// registerChannel create channel
func registerChannel(topic string) {
	c := &channel{
		Subscriptions: make(map[string]subscription),
		MessageChan:   make(chan Message),
	}

	channels[topic] = c
}

// Publish a message for topic
func Publish(topic string, message Message) error {
	c, err := getChannel(topic)
	if err != nil {
		return err
	}

	c.MessageChan <- message
	return nil
}

// get channel by topic, return error if not exist.
func getChannel(topic string) (*channel, error) {
	c := channels[topic]
	if c == nil {
		return nil, models.Error{Message: fmt.Sprintf("There is no channel for topic %s", topic)}
	}
	return c, nil
}
