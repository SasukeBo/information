package admin

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/admin/device"
	"github.com/SasukeBo/information/schema/scalars"
	"github.com/SasukeBo/information/schema/types"
)

// DeviceGetField get a device
var DeviceGetField = &graphql.Field{
	Type: types.Device,
	Args: graphql.FieldConfigArgument{
		"uuid": fields.GenArg(graphql.String, "设备UUID", false),
	},
	Resolve: device.Get,
}

// DeviceListField _
var DeviceListField = &graphql.Field{
	Type: graphql.NewList(types.Device),
	Args: graphql.FieldConfigArgument{
		"limit":       fields.GenArg(graphql.Int, "最大获取条数"),
		"offset":      fields.GenArg(graphql.Int, "偏移量"),
		"type":        fields.GenArg(graphql.String, "设备类型"),
		"namePattern": fields.GenArg(graphql.String, "设备名称模糊匹配"),
		"status":      fields.GenArg(graphql.NewList(scalars.BaseStatus), "设备状态列表"),
		"userName":    fields.GenArg(graphql.String, "设备注册人姓名"),
		"userUUID":    fields.GenArg(graphql.String, "设备创建人UUID"),
		"chargerName": fields.GenArg(graphql.String, "设备负责人姓名"),
		"chargerUUID": fields.GenArg(graphql.String, "设备负责人UUID"),
	},
	Description: `
	管理员获取设备列表，返回记录列表按照设备创建时间倒序排列，
	按照时间的升降序排列控制交由前端处理。
	`,
	Resolve: device.List,
}

// DeviceUpdateField _
var DeviceUpdateField = &graphql.Field{
	Type: types.Device,
	Args: graphql.FieldConfigArgument{
		"uuid":        fields.GenArg(graphql.String, "设备UUID", false),
		"name":        fields.GenArg(graphql.String, "设备名称"),
		"mac":         fields.GenArg(graphql.String, "设备Mac地址"),
		"status":      fields.GenArg(scalars.BaseStatus, "设备状态"),
		"description": fields.GenArg(graphql.String, "设备描述"),
	},
	Resolve: device.Update,
}

// DeviceDeleteField _
var DeviceDeleteField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"uuid": fields.GenArg(graphql.String, "设备UUID", false),
	},
	Resolve: device.Delete,
}
