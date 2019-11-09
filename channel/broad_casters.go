package channel

import (
	"encoding/json"
	"github.com/SasukeBo/information/schema"
	"github.com/astaxie/beego/logs"
	"github.com/graphql-go/graphql"
	"golang.org/x/net/websocket"
	"strconv"
)

// BroadcastDeviceStatus broadcast message for Topic deviceStatus
func BroadcastDeviceStatus(message Message) {
	c, err := getChannel("deviceStatus")
	if err != nil {
		logs.Error(err)
	}

	logs.Info("deviceStatus channel subs:", c.Subscriptions)
	for _, sub := range c.Subscriptions {
		broadcast(sub, message)
	}
}

// BroadcastProductIns broadcast message for Topic productIns
func BroadcastProductIns(message Message) {
	c, err := getChannel("productIns")
	if err != nil {
		logs.Error(err)
	}

	for _, sub := range c.Subscriptions {
		broadcast(sub, message)
	}
}

// handle graphql subscription query broadcast
func broadcast(sub subscription, message Message) {
	if ok := isMatched(sub.Variables[message.IDName], message.IDValue); !ok {
		return
	}

	if sub.Query == "" {
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:         schema.Root,
		RequestString:  sub.Query,
		VariableValues: map[string]interface{}{"id": message.IDValue},
	})

	data, err := json.Marshal(map[string]interface{}{
		"type":    "data",
		"id":      sub.RefID,
		"payload": result,
	})

	if err != nil {
		logs.Error(err)
		return
	}

	websocket.Message.Send(sub.Conn, string(data))
}

// isMatched compare interface{} value with int value
// string 1 == int 1
// float32/64 1 == int 1
// int32/64 == int 1
// return true if equal
func isMatched(sv interface{}, id int) bool {
	switch v := sv.(type) {
	case int32:
	case int64:
	case float32:
	case float64:
		if int(v) == id {
			return true
		}
	case string:
		intV, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		if intV == id {
			return true
		}
	}

	return false
}
