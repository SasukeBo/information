package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/*							   types
------------------------------------------ */

// DeviceStatusLog 设备状态变更记录类型
var DeviceStatusLog = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeviceStatusLog",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"status":    &graphql.Field{Type: DeviceStatus, Description: "设备运行状态"},
		"duration":  &graphql.Field{Type: graphql.Int, Description: "持续时间（s）"},
		"createdAt": &graphql.Field{Type: graphql.DateTime, Description: "变更时间"},
	},
})

func init() {
	DeviceStatusLog.AddFieldConfig("device", &graphql.Field{Type: DeviceType, Description: "设备", Resolve: resolver.LoadDevice})
}

/*							query
----------------------------------- */

var deviceStatusLogList = &graphql.Field{
	Type: graphql.NewList(DeviceStatusLog),
	Args: graphql.FieldConfigArgument{
		"deviceUUID": GenArg(graphql.String, "设备UUID", false),
		"status":     GenArg(DeviceStatus, "运行状态"),
		"beforeTime": GenArg(graphql.DateTime, "开始时间"),
		"afterTime":  GenArg(graphql.DateTime, "结束时间"),
		"limit":      GenArg(graphql.Int, "最大返回条数"),
	},
	Description: "获取设备状态变更列表",
	Resolve:     resolver.ListDeviceStatusLog,
}

var deviceStatusDuration = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"deviceID": GenArg(graphql.Int, "设备ID", false),
		"status":   GenArg(DeviceStatus, "运行状态", false),
	},
	Description: "获取设备状态持续时间",
	Resolve:     resolver.DeviceStatusDuration,
}

/*						mutation
----------------------------------- */

// TODO: 重构subscription的resolver
var deviceStatusRefresh = &graphql.Field{
	Type: DeviceType,
	Args: graphql.FieldConfigArgument{
		"deviceID": GenArg(graphql.Int, "设备ID", false),
	},
	Description: "刷新设备状态",
	Resolve:     resolver.RefreshDeviceStatus,
}
