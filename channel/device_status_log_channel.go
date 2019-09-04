package channel

import (
	"container/list"
	// "encoding/json"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	// "github.com/graphql-go/graphql"
	// "golang.org/x/net/websocket"
	// "strconv"
	// "github.com/SasukeBo/information/models"
	// "github.com/SasukeBo/information/schema"
)

// device status log channel
type dslChannelType struct {
	channelType
}

func (dsl *dslChannelType) Join(msg *SocketMessage) {
	if err := join(dsl, msg); err != nil {
		logs.Error(err)
	}
}

func (dsl *dslChannelType) Leave(msg *SocketMessage) {
	if err := leave(dsl, msg); err != nil {
		logs.Error(err)
	}
}

func (dsl *dslChannelType) HandleIn(msg *SocketMessage) {
	dsl.Messagechan <- *msg
}

func (dsl *dslChannelType) HandleOut(msg *SocketMessage) {

}

func init() {
	dslChannel := dslChannelType{
		channelType{
			Subscribers: make(map[string]*list.List),
			Messagechan: make(chan SocketMessage, 10),
		},
	}

	go channel("device_status_log:*", &dslChannel)
}
