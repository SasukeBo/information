package types

import (
	"github.com/SasukeBo/information/resolvers/device"
	"github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

// DeviceStatusLog 设备状态变更记录
var DeviceStatusLog graphql.Type

// DeviceStatusLogListType doc false
var DeviceStatusLogListType *graphql.Field

func init() {
	DeviceStatusLog = graphql.NewObject(graphql.ObjectConfig{
		Name: "DeviceStatusLog",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":       &graphql.Field{Type: graphql.Int},
				"status":   &graphql.Field{Type: custom.DeviceStatus, Description: "设备运行状态"},
				"device":   &graphql.Field{Type: Device, Description: "设备"},
				"changeAt": &graphql.Field{Type: graphql.DateTime, Description: "变更时间"},
			}
		}),
	})

	DeviceStatusLogListType = &graphql.Field{
		Type: graphql.NewList(DeviceStatusLog),
		Args: graphql.FieldConfigArgument{
			"deviceID":   GenArg(graphql.Int, "设备ID", false),
			"status":     GenArg(custom.DeviceStatus, "运行状态"),
			"beforeTime": GenArg(graphql.DateTime, "开始时间"),
			"afterTime":  GenArg(graphql.DateTime, "结束时间"),
		},
		Resolve: device.StatusLogList,
	}
}
