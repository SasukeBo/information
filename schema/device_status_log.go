package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/*							   types
------------------------------------------ */

var deviceStatusLogType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "DeviceStatusLog",
	Description: "设备状态变更记录类型，考虑到列表查询是Load停机原因会造成很大时间开销，不推荐直接Load",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.Int},
		"status":   &graphql.Field{Type: DeviceStatus, Description: "设备运行状态"},
		"beginAt":  &graphql.Field{Type: graphql.DateTime, Description: "开始时间"},
		"finishAt": &graphql.Field{Type: graphql.DateTime, Description: "结束时间"},
	},
})

func init() {
	deviceStatusLogType.AddFieldConfig("reasons", &graphql.Field{
		Type:        graphql.NewList(stopReasonType),
		Description: "停机原因",
		Resolve:     resolver.LogLoadStopReason,
	})
}

/*						response
----------------------------------- */

var deviceStopLogListResponse = graphql.NewObject(graphql.ObjectConfig{
	Name:        "DeviceStopLogListResponse",
	Description: "停机日志分页查询返回对象",
	Fields: graphql.Fields{
		"total": &graphql.Field{Type: graphql.Int, Description: "总数量"},
		"logs":  &graphql.Field{Type: graphql.NewList(deviceStatusLogType), Description: "设备运行状态"},
	},
})

var typeCount = graphql.NewObject(graphql.ObjectConfig{
	Name:        "TypeCount",
	Description: "类型数量",
	Fields: graphql.Fields{
		"name":    &graphql.Field{Type: graphql.String, Description: "类型名称"},
		"numbers": &graphql.Field{Type: graphql.NewList(graphql.Int), Description: "统计数列表，对应日期列表顺序"},
	},
})

var deviceStopTypeCountResponse = graphql.NewObject(graphql.ObjectConfig{
	Name:        "DeviceStopTypeCountResponse",
	Description: "停机类型次数统计结果",
	Fields: graphql.Fields{
		"days":   &graphql.Field{Type: graphql.NewList(graphql.String), Description: "日期"},
		"types":  &graphql.Field{Type: graphql.NewList(graphql.String), Description: "停机类型名称列表"},
		"counts": &graphql.Field{Type: graphql.NewList(typeCount), Description: "设备运行状态"},
	},
})

/*							query
----------------------------------- */

var deviceStopLogList = &graphql.Field{
	Type: deviceStopLogListResponse,
	Args: graphql.FieldConfigArgument{
		"deviceID":  GenArg(graphql.Int, "设备ID", false),
		"beginTime": GenArg(graphql.DateTime, "开始时间"),
		"endTime":   GenArg(graphql.DateTime, "结束时间"),
		"offset":    GenArg(graphql.Int, "分页查询偏移量"),
		"limit":     GenArg(graphql.Int, "分页最大返回条数"),
	},
	Description: "获取设备停机日志列表",
	Resolve:     resolver.ListDeviceStopLogs,
}

var deviceStopTypeCount = &graphql.Field{
	Type: deviceStopTypeCountResponse,
	Args: graphql.FieldConfigArgument{
		"deviceID":  GenArg(graphql.Int, "设备ID", false),
		"beginTime": GenArg(graphql.DateTime, "开始时间", false),
		"endTime":   GenArg(graphql.DateTime, "结束时间", false),
	},
	Description: "获取设备停机类型次数统计",
	Resolve:     resolver.CountDeviceStopType,
}
