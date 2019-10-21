package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/* 					 types
------------------------------ */

// DeviceStatisticsType 设备统计数据
var DeviceStatisticsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeviceStatistics",
	Fields: graphql.Fields{
		"activation": &graphql.Field{Type: graphql.Float, Description: "稼动率"},
		"yield":      &graphql.Field{Type: graphql.Float, Description: "良率"},
	},
})

var deviceType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Device",
	Description: "设备",
	Fields: graphql.Fields{
		"id":             &graphql.Field{Type: graphql.Int},
		"type":           &graphql.Field{Type: graphql.String, Description: "设备类型"},
		"name":           &graphql.Field{Type: graphql.String, Description: "设备名称"},
		"token":          &graphql.Field{Type: graphql.String, Description: "设备token，用于数据加密"},
		"status":         &graphql.Field{Type: DeviceStatus, Description: "基础状态"},
		"address":        &graphql.Field{Type: graphql.String, Description: "设备地址"},
		"number":         &graphql.Field{Type: graphql.String, Description: "设备编号"},
		"statusChangeAt": &graphql.Field{Type: graphql.DateTime, Description: "设备状态变更时间"},
		"statistics":     &graphql.Field{Type: DeviceStatisticsType, Description: "设备统计数据"},
		"createdAt":      &graphql.Field{Type: graphql.DateTime},
		"updatedAt":      &graphql.Field{Type: graphql.DateTime},
		"remoteIP":       &graphql.Field{Type: graphql.String},
	},
})

var devicePrivateFormInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "DevicePrivateFormInput",
	Description: "设备注册私有字段表单类型",
	Fields: graphql.InputObjectConfigFieldMap{
		"address": &graphql.InputObjectFieldConfig{Type: graphql.String, Description: "设备地址"},
		"number":  &graphql.InputObjectFieldConfig{Type: graphql.String, Description: "设备编号"},
	},
})

func init() {
	deviceType.AddFieldConfig("user", &graphql.Field{
		Type:        userType,
		Description: "注册人用户",
		Resolve:     resolver.LoadUser,
	})

	deviceType.AddFieldConfig("product", &graphql.Field{
		Type:        productType,
		Description: "生产产品",
		Resolve:     resolver.DeviceLoadProduct,
	})

	deviceType.AddFieldConfig("deviceChargers", &graphql.Field{
		Type:        graphql.NewList(deviceChargerType),
		Description: "设备负责人",
		Resolve:     resolver.LoadDeviceCharger,
	})
}

/* 					response
------------------------------ */

var deviceListResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeviceListResponse",
	Fields: graphql.Fields{
		"total":   &graphql.Field{Type: graphql.Int, Description: "当前筛选条件下记录数"},
		"devices": &graphql.Field{Type: graphql.NewList(deviceType), Description: "设备列表"},
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

var deviceMonthlyStatisticsResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeviceMonthlyStatisticsResponse",
	Fields: graphql.Fields{
		"runningTime": &graphql.Field{
			Type:        graphql.String,
			Description: "运行时间",
			Args:        graphql.FieldConfigArgument{"format": GenArg(graphql.String, "时间格式化模板")},
			Resolve:     resolver.MonthlyAnalyzeDeviceFormatTime,
		},
		"activation": &graphql.Field{Type: graphql.Float, Description: "稼动率"},
		"yieldRate":  &graphql.Field{Type: graphql.Float, Description: "良率"},
		"yield":      &graphql.Field{Type: graphql.Int, Description: "产量"},
	},
	Description: "设备月数据统计结果对象",
})

var realTimeStatisticsObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "RealTimeStatisticsObject",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int, Description: "检测值ID"},
		"sign":      &graphql.Field{Type: graphql.String, Description: "检测项名称"},
		"value":     &graphql.Field{Type: graphql.Float, Description: "检测值"},
		"createdAt": &graphql.Field{Type: graphql.DateTime, Description: "创建时间"},
	},
	Description: "设备生产实时数据对象",
})

/* 				   query
------------------------------ */

var realTimeStatistics = &graphql.Field{
	Type: graphql.NewList(realTimeStatisticsObject),
	Args: graphql.FieldConfigArgument{
		"deviceID":  GenArg(graphql.Int, "设备ID", false),
		"productID": GenArg(graphql.Int, "产品ID", false),
		"limit":     GenArg(graphql.Int, "数量限制", false),
		"afterTime": GenArg(graphql.DateTime, "起始时间"),
	},
	Description: `设备生产实时数据`,
	Resolve:     resolver.GetRealTimeStatistics,
}

var deviceMonthlyStatistics = &graphql.Field{
	Type:        deviceMonthlyStatisticsResponse,
	Args:        graphql.FieldConfigArgument{"id": GenArg(graphql.Int, "设备ID", false)},
	Description: `设备月数据统计`,
	Resolve:     resolver.MonthlyAnalyzeDevice,
}

var deviceStatusCount = &graphql.Field{
	Type: deviceStatusCountResponse,
	Args: graphql.FieldConfigArgument{
		"filter": GenArg(graphql.String, "条件筛选，默认值`all`，`register`统计当前用户注册的设备，`all`统计平台所有设备", true, "all"),
	},
	Description: `统计设备各状态数量`,
	Resolve:     resolver.CountDeviceStatus,
}

var deviceGet = &graphql.Field{
	Type: deviceType,
	Args: graphql.FieldConfigArgument{
		"id": GenArg(graphql.Int, "设备ID", false),
	},
	Description: "获取device",
	Resolve:     resolver.GetDevice,
}

var deviceTokenGet = &graphql.Field{
	Type: deviceType,
	Args: graphql.FieldConfigArgument{
		"token": GenArg(graphql.String, "设备token", false),
	},
	Description: "使用token获取device",
	Resolve:     resolver.GetDeviceByToken,
}

var deviceList = &graphql.Field{
	Type: deviceListResponse,
	Args: graphql.FieldConfigArgument{
		"limit":  GenArg(graphql.Int, "最大数量"),
		"offset": GenArg(graphql.Int, "数据偏移量"),
		"search": GenArg(graphql.String, "设备类型/设备名称/地址/设备编号模糊搜索"),
		"status": GenArg(DeviceStatus, "设备状态 prod/stop/offline"),
		"self":   GenArg(graphql.Boolean, "仅显示本人创建"),
	},
	Description: "获取device列表",
	Resolve:     resolver.ListDevice,
}

/* 					mutation
------------------------------ */

var deviceCreate = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"name":         GenArg(graphql.String, "设备名称", false),
		"type":         GenArg(graphql.String, "设备类型", false),
		"productID":    GenArg(graphql.Int, "生产产品ID"),
		"privateForms": GenArg(graphql.NewList(devicePrivateFormInputType), "设备私有字段", false),
	},
	Resolve: resolver.CreateDevice,
}

var deviceUpdate = &graphql.Field{
	Type: deviceType,
	Args: graphql.FieldConfigArgument{
		"id":      GenArg(graphql.Int, "设备ID", false),
		"address": GenArg(graphql.String, "设备地址"),
		"name":    GenArg(graphql.String, "设备名称"),
		"number":  GenArg(graphql.String, "设备编号"),
		"type":    GenArg(graphql.String, "设备类型"),
	},
	Resolve: resolver.UpdateDevice,
}

var deviceDelete = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"uuid": GenArg(graphql.String, "设备UUID", false),
	},
	Resolve: resolver.DeleteDevice,
}
