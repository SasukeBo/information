package public

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/device"
	"github.com/SasukeBo/information/schema/scalars"
	"github.com/SasukeBo/information/schema/types"
)

/* 					 types
------------------------------ */

// DeviceStatusCount graphql object contian counts of each device status
var DeviceStatusCount = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeviceStatusCount",
	Fields: graphql.FieldsThunk(func() graphql.Fields {
		return graphql.Fields{
			"offline": &graphql.Field{Type: graphql.Int, Description: "离线状态下的设备数量"},
			"prod":    &graphql.Field{Type: graphql.Int, Description: "生产状态下的设备数量"},
			"stop":    &graphql.Field{Type: graphql.Int, Description: "停机状态下的设备数量"},
		}
	}),
})

// DeviceList _
var DeviceList = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeviceList",
	Fields: graphql.FieldsThunk(func() graphql.Fields {
		return graphql.Fields{
			"total":   &graphql.Field{Type: graphql.Int, Description: "当前筛选条件下记录数"},
			"devices": &graphql.Field{Type: graphql.NewList(types.Device), Description: "设备列表"},
		}
	}),
})

/* 					 fields
------------------------------ */

// DeviceStatusCountField _
var DeviceStatusCountField = &graphql.Field{
	Type: DeviceStatusCount,
	Args: graphql.FieldConfigArgument{
		"filter": fields.GenArg(
			graphql.String,
			"条件筛选，'register'为请求当前用户创建的，'all'为请求所有设备的统计数据，默认为'all'",
			true,
			"all",
		),
	},
	Resolve: device.CountDeviceStatus,
}

// DeviceCreateField create a device
var DeviceCreateField = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"count":       fields.GenArg(graphql.Int, "创建数量", false, 1),
		"type":        fields.GenArg(graphql.String, "设备类型", false),
		"name":        fields.GenArg(graphql.String, "设备名称", false),
		"address":     fields.GenArg(graphql.String, "设备地址", false),
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
	Type: DeviceList,
	Args: graphql.FieldConfigArgument{
		"limit":      fields.GenArg(graphql.Int, "最大数量"),
		"offset":     fields.GenArg(graphql.Int, "数据偏移量"),
		"pattern":    fields.GenArg(graphql.String, "设备类型/设备名称模糊搜索"),
		"status":     fields.GenArg(scalars.DeviceStatus, "设备状态 prod/stop/offline"),
		"isRegister": fields.GenArg(graphql.Boolean, "是否为创建人"),
	},
	Description: "查询device列表",
	Resolve:     device.List,
}
