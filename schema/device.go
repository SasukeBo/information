package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/* 					 types
------------------------------ */

// Device 设备类型
var Device = graphql.NewObject(graphql.ObjectConfig{
	Name: "Device",
	Fields: graphql.FieldsThunk(func() graphql.Fields {
		return graphql.Fields{
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
			// TODO: "statistics": 		&graphql.Field{Type: graphql.String},
		}
	}),
})

func init() {
	Device.AddFieldConfig("user", &graphql.Field{Type: User, Description: "注册人用户", Resolve: resolver.LoadUser})
	Device.AddFieldConfig("params", &graphql.Field{Type: graphql.NewList(DeviceParam), Description: "设备参数", Resolve: resolver.LoadDeviceParam})
	Device.AddFieldConfig("deviceChargers", &graphql.Field{Type: graphql.NewList(DeviceCharger), Description: "设备负责人", Resolve: resolver.LoadDeviceCharger})
}

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
			"devices": &graphql.Field{Type: graphql.NewList(Device), Description: "设备列表"},
		}
	}),
})

/* 					 fields
------------------------------ */

// DeviceStatusCountField _
var DeviceStatusCountField = &graphql.Field{
	Type: DeviceStatusCount,
	Args: graphql.FieldConfigArgument{
		"filter": GenArg(graphql.String, "条件筛选，'register'为请求当前用户创建的，'all'为请求所有设备的统计数据，默认为'all'", true, "all"),
	},
	Resolve: resolver.CountDeviceStatus,
}

// DeviceCreateField create a device
var DeviceCreateField = &graphql.Field{
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
var DeviceUpdateField = &graphql.Field{
	Type: Device,
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
var DeviceDeleteField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"uuid": GenArg(graphql.String, "设备UUID", false),
	},
	Resolve: resolver.DeleteDevice,
}

// DeviceGetField get a device by uuid
var DeviceGetField = &graphql.Field{
	Type: Device,
	Args: graphql.FieldConfigArgument{
		"uuid": GenArg(graphql.String, "设备UUID", false),
	},
	Description: "使用uuid获取device",
	Resolve:     resolver.GetDevice,
}

// DeviceTokenGetField get a device by token
var DeviceTokenGetField = &graphql.Field{
	Type: Device,
	Args: graphql.FieldConfigArgument{
		"token": GenArg(graphql.String, "设备token", false),
	},
	Description: "使用token获取device",
	Resolve:     resolver.GetDeviceByToken,
}

// DeviceListField get a device list with options
var DeviceListField = &graphql.Field{
	Type: DeviceList,
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
