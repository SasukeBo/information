package channel

// 设备参数值channel
// 参照 phoenix channel，每个 channel 都是独立的线程

import (
	"container/list"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/graphql-go/graphql"
	"golang.org/x/net/websocket"
	"strconv"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/schema"
)

// dpvChannelType device param value channel type
type dpvChannelType struct {
	channelType
}

// dpvChannel device param value channel
var dpvChannel = dpvChannelType{
	channelType{
		Subscribers: make(map[string]*list.List),
		Messagechan: make(chan SocketMessage, 10),
	},
}

// Join a topic
func (dpv *dpvChannelType) Join(message *SocketMessage) {
	conn := message.Socket
	if conn == nil {
		logs.Error("Can't create subscribe without socket object!")
		return
	}

	subTopic := getSubTopic(message.Topic)
	subs := dpv.Subscribers[subTopic]
	if subs == nil {
		subs = list.New()
	}

	sessionID, user, err := fetchSessionIDAndUserUUID(conn)
	if err != nil {
		logs.Error(err)
		return
	}
	// TODO: validate privilege

	sub := Subscribe{
		Topic:     subTopic,
		UserUUID:  user.(*models.User).UUID,
		SessionID: sessionID.(string),
		Socket:    conn,
		Payload:   message.Payload,
		ID:        message.Ref,
	}

	subs.PushBack(sub)
	dpv.Subscribers[subTopic] = subs
	logs.Warn("user:%s join %s", user.(*models.User).Phone, message.Topic)
}

// Leave 取消订阅
func (dpv *dpvChannelType) Leave(message *SocketMessage) {
	conn := message.Socket
	if conn == nil {
		logs.Error("Can't unsubscribe without socket object!")
		return
	}

	subTopic := getSubTopic(message.Topic)
	sessionID, user, err := fetchSessionIDAndUserUUID(conn)
	if err != nil {
		logs.Error(err)
		return
	}

	subs := dpv.Subscribers[subTopic]
	newSubs := unsubscribe(subs, sessionID.(string), user.(*models.User))
	dpv.Subscribers[subTopic] = newSubs
	logs.Warn("user:%s leave %s", user.(*models.User).Phone, message.Topic)
}

// HandleIn 处理消息进入
func (dpv *dpvChannelType) HandleIn(message *SocketMessage) {
	dpv.Messagechan <- *message
}

// HandleOut 处理消息发出
func (dpv *dpvChannelType) HandleOut(message *SocketMessage) {
	subTopic := getSubTopic(message.Topic)
	subs := dpv.Subscribers[subTopic]
	paramIDStr, ok := message.Payload["paramID"].(string)
	if !ok {
		logs.Error("paramID type assert string failed")
		return
	}

	paramID, err := strconv.ParseInt(paramIDStr, 10, 0)
	if err != nil {
		logs.Error(err)
		return
	}

	paramValue := &models.DeviceParamValue{
		Value:       message.Payload["value"].(string),
		DeviceParam: &models.DeviceParam{ID: int(paramID)},
	}

	if err := paramValue.Insert(); err != nil {
		logs.Error(err)
		return
	}

	for el := subs.Front(); el != nil; el = el.Next() {
		sub := el.Value.(Subscribe)
		query := sub.Payload["query"].(string)
		result := graphql.Do(graphql.Params{
			Schema:         schema.PublicSchema,
			RequestString:  query,
			VariableValues: map[string]interface{}{"id": paramValue.ID},
		})

		message, err := json.Marshal(map[string]interface{}{
			"type":    "data",
			"id":      sub.ID,
			"payload": result,
		})
		if err != nil {
			logs.Error(err)
			continue
		}
		websocket.Message.Send(sub.Socket, string(message))
	}
}

func (dpv *dpvChannelType) getMessageChan() chan SocketMessage {
	return dpv.Messagechan
}

func fetchSessionIDAndUserUUID(conn *websocket.Conn) (interface{}, interface{}, error) {
	session, err := conn.Request().Cookie(beego.AppConfig.String("SessionName"))
	if err != nil {
		return nil, nil, err
	}

	userLogin := models.UserLogin{SessionID: session.Value}
	if err := userLogin.GetBy("session_id"); err != nil {
		return nil, nil, err
	}

	user, err := userLogin.LoadUser()
	if err != nil {
		return nil, nil, err
	}

	return session.Value, user, nil
}

func init() {
	go channel("device_param_value:*", &dpvChannel)
}
