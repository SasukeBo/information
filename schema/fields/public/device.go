package public

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/device"
	"github.com/SasukeBo/information/schema/scalars"
	"github.com/SasukeBo/information/schema/types"
)

// DeviceCreateField create a device
var DeviceCreateField = &graphql.Field{
	Type: types.Device,
	Args: graphql.FieldConfigArgument{
		"type":        fields.GenArg(graphql.String, "设备类型", false),
		"name":        fields.GenArg(graphql.String, "设备名称", false),
		"description": fields.GenArg(graphql.String, "描述"),
	},
	Resolve: device.Create,
}

// DeviceUpdateField update a device
var DeviceUpdateField = &graphql.Field{
	Type: types.Device,
	Args: graphql.FieldConfigArgument{
		"uuid":        fields.GenArg(graphql.String, "设备UUID", false),
		"type":        fields.GenArg(graphql.String, "设备类型"),
		"name":        fields.GenArg(graphql.String, "设备名称"),
		"status":      fields.GenArg(scalars.DeviceStatus, "设备状态 prod/stop/offline/online"),
		"description": fields.GenArg(graphql.String, "描述"),
	},
	Resolve: device.Update,
}

// DeviceDeleteField update a device
var DeviceDeleteField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"uuid": fields.GenArg(graphql.String, "设备UUID", false),
	},
	Resolve: device.Delete,
}

// DeviceGetField get a device by uuid
var DeviceGetField = &graphql.Field{
	Type: types.Device,
	Args: graphql.FieldConfigArgument{
		"uuid": fields.GenArg(graphql.String, "设备UUID", false),
	},
	Description: "使用uuid获取device",
	Resolve:     device.Get,
}

// DeviceTokenGetField get a device by token
var DeviceTokenGetField = &graphql.Field{
	Type: types.Device,
	Args: graphql.FieldConfigArgument{
		"token": fields.GenArg(graphql.String, "设备token", false),
	},
	Description: "使用token获取device",
	Resolve:     device.GetByToken,
}

// DeviceListField get a device list with options
var DeviceListField = &graphql.Field{
	Type: graphql.NewList(types.Device),
	Args: graphql.FieldConfigArgument{
		"type":        fields.GenArg(graphql.String, "设备类型"),
		"namePattern": fields.GenArg(graphql.String, "设备名称模糊匹配"),
		"status":      fields.GenArg(scalars.DeviceStatus, "设备状态 prod/stop/offline/online"),
		"userUUID":    fields.GenArg(graphql.String, "注册人uuid"),
		"ownership":   fields.GenArg(graphql.NewList(graphql.String), "物主身份，register、charger", true, []string{}),
	},
	Description: "查询device列表",
	Resolve:     device.List,
}

// DeviceBindField bind a device
var DeviceBindField = &graphql.Field{
	Type: types.Device,
	Args: graphql.FieldConfigArgument{
		"token": fields.GenArg(graphql.String, "设备token，用于数据加密", false),
		"mac":   fields.GenArg(graphql.String, "设备Mac地址", false),
	},
	Description: "绑定物理设备Mac地址",
	Resolve:     device.Bind,
}
