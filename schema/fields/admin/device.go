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
		"type":        fields.GenArg(graphql.String, "设备类型"),
		"namePattern": fields.GenArg(graphql.String, "设备名称模糊匹配"),
		"status":      fields.GenArg(graphql.NewList(scalars.BaseStatus), "设备状态列表"),
		// TODO: "userName":      fields.GenArg(graphql.NewList(scalars.BaseStatus), "设备状态列表"),
		// TODO: 管理员接口，设备接口，包含权限验证
	},
	Resolve: device.List,
}

// DeviceUpdateField _
var DeviceUpdateField = &graphql.Field{
	Type: graphql.NewList(types.Device),
	Args: graphql.FieldConfigArgument{
		// TODO:
		"todo": fields.GenArg(graphql.String, "todo"),
	},
	Resolve: device.Update,
}

// DeviceDeleteField _
var DeviceDeleteField = &graphql.Field{
	Type: graphql.NewList(types.Device),
	Args: graphql.FieldConfigArgument{
		// TODO:
		"todo": fields.GenArg(graphql.String, "todo"),
	},
	Resolve: device.Delete,
}
