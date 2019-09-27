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

/*							query fields
------------------------------------------ */

// DeviceStatusLogListField doc false
var deviceStatusLogListField = &graphql.Field{
	Type: graphql.NewList(DeviceStatusLog),
	Args: graphql.FieldConfigArgument{
		"deviceUUID": GenArg(graphql.String, "设备UUID", false),
		"status":     GenArg(DeviceStatus, "运行状态"),
		"beforeTime": GenArg(graphql.DateTime, "开始时间"),
		"afterTime":  GenArg(graphql.DateTime, "结束时间"),
		"limit":      GenArg(graphql.Int, "最大返回条数"),
	},
	Resolve: resolver.ListDeviceStatusLog,
}

// DeviceStatusDurationField _
var deviceStatusDurationField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"deviceID": GenArg(graphql.Int, "设备ID", false),
		"status":   GenArg(DeviceStatus, "运行状态", false),
	},
	Resolve: resolver.DeviceStatusDuration,
}

/*						mutation fields
------------------------------------------ */

// DeviceStatusRefreshField _
var deviceStatusRefreshField = &graphql.Field{
	Type: DeviceType,
	Args: graphql.FieldConfigArgument{
		"deviceID": GenArg(graphql.Int, "设备ID", false),
	},
	Resolve: resolver.RefreshDeviceStatus,
}
