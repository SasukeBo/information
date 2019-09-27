package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/* 					 types
------------------------------ */

// DeviceType 设备类型
var DeviceType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Device",
	Fields: graphql.Fields{
		"type":           &graphql.Field{Type: graphql.String, Description: "设备类型"},
		"name":           &graphql.Field{Type: graphql.String, Description: "设备名称"},
		"token":          &graphql.Field{Type: graphql.String, Description: "设备token，用于数据加密"},
		"status":         &graphql.Field{Type: DeviceStatus, Description: "基础状态"},
		"address":        &graphql.Field{Type: graphql.String, Description: "设备地址"},
		"number":         &graphql.Field{Type: graphql.String, Description: "设备编号"},
		"id":             &graphql.Field{Type: graphql.Int},
		"uuid":           &graphql.Field{Type: graphql.String, Description: "设备UUID"},
		"statusChangeAt": &graphql.Field{Type: graphql.DateTime, Description: "设备状态变更时间"},
		"description":    &graphql.Field{Type: graphql.String, Description: "设备描述，备注"},
		"createdAt":      &graphql.Field{Type: graphql.DateTime},
		"updatedAt":      &graphql.Field{Type: graphql.DateTime},
		"remoteIP":       &graphql.Field{Type: graphql.String},
		// "statistics": 		&graphql.Field{Type: graphql.String},
	},
})

func init() {
	DeviceType.AddFieldConfig("user", &graphql.Field{Type: User, Description: "注册人用户", Resolve: resolver.LoadUser})
	DeviceType.AddFieldConfig("params", &graphql.Field{Type: graphql.NewList(DeviceParam), Description: "设备参数", Resolve: resolver.LoadDeviceParam})
	DeviceType.AddFieldConfig("deviceChargers", &graphql.Field{Type: graphql.NewList(DeviceCharger), Description: "设备负责人", Resolve: resolver.LoadDeviceCharger})
}

/* 					response
------------------------------ */

var deviceListResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeviceListResponse",
	Fields: graphql.Fields{
		"total":   &graphql.Field{Type: graphql.Int, Description: "当前筛选条件下记录数"},
		"devices": &graphql.Field{Type: graphql.NewList(DeviceType), Description: "设备列表"},
	},
	Description: "设备列表对象",
})

var deviceStatusCountResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeviceStatusCountResponse",
	Fields: graphql.Fields{
		"offline": &graphql.Field{Type: graphql.Int, Description: "离线状态下的设备数量"},
		"prod":    &graphql.Field{Type: graphql.Int, Description: "生产状态下的设备数量"},
		"stop":    &graphql.Field{Type: graphql.Int, Description: "停机状态下的设备数量"},
	},
	Description: "设备状态数量对象",
})

/* 				query fields
------------------------------ */

var deviceStatusCountField = &graphql.Field{
	Type: deviceStatusCountResponse,
	Args: graphql.FieldConfigArgument{
		"filter": GenArg(graphql.String, `
条件筛选，默认值'all'
- 'register' 统计当前用户注册的设备
- 'all' 统计平台所有设备`,
			true, "all"),
	},
	Description: `统计设备各状态数量`,
	Resolve:     resolver.CountDeviceStatus,
}

// deviceGetField get a device by uuid
var deviceGetField = &graphql.Field{
	Type: DeviceType,
	Args: graphql.FieldConfigArgument{
		"uuid": GenArg(graphql.String, "设备UUID", false),
	},
	Description: "使用uuid获取device",
	Resolve:     resolver.GetDevice,
}

// deviceTokenGetField get a device by token
var deviceTokenGetField = &graphql.Field{
	Type: DeviceType,
	Args: graphql.FieldConfigArgument{
		"token": GenArg(graphql.String, "设备token", false),
	},
	Description: "使用token获取device",
	Resolve:     resolver.GetDeviceByToken,
}

// DeviceListField get a device list with options
var deviceListField = &graphql.Field{
	Type: deviceListResponse,
	Args: graphql.FieldConfigArgument{
		"limit":      GenArg(graphql.Int, "最大数量"),
		"offset":     GenArg(graphql.Int, "数据偏移量"),
		"pattern":    GenArg(graphql.String, "设备类型/设备名称模糊搜索"),
		"status":     GenArg(DeviceStatus, "设备状态 prod/stop/offline"),
		"isRegister": GenArg(graphql.Boolean, "是否为创建人"),
	},
	Description: "查询device列表",
	Resolve:     resolver.ListDevice,
}

/* 			mutation fields
------------------------------ */

// DeviceCreateField create a device
var deviceCreateField = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"count":       GenArg(graphql.Int, "创建数量", false, 1),
		"type":        GenArg(graphql.String, "设备类型", false),
		"name":        GenArg(graphql.String, "设备名称", false),
		"address":     GenArg(graphql.String, "设备地址", false),
		"description": GenArg(graphql.String, "描述"),
	},
	Resolve: resolver.CreateDevice,
}

// DeviceUpdateField update a device
var deviceUpdateField = &graphql.Field{
	Type: DeviceType,
	Args: graphql.FieldConfigArgument{
		"uuid":        GenArg(graphql.String, "设备UUID", false),
		"type":        GenArg(graphql.String, "设备类型"),
		"name":        GenArg(graphql.String, "设备名称"),
		"status":      GenArg(DeviceStatus, "设备状态 prod/stop/offline/online"),
		"description": GenArg(graphql.String, "描述"),
	},
	Resolve: resolver.UpdateDevice,
}

// DeviceDeleteField update a device
var deviceDeleteField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"uuid": GenArg(graphql.String, "设备UUID", false),
	},
	Resolve: resolver.DeleteDevice,
}
