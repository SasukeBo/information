package types

import (
	"github.com/SasukeBo/information/resolvers/device"
	"github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

// Device 设备类型
var Device graphql.Type

// DeviceCreateType create a device
var DeviceCreateType *graphql.Field

// DeviceUpdateType update a device
var DeviceUpdateType *graphql.Field

// DeviceDeleteType update a device
var DeviceDeleteType *graphql.Field

// DeviceGetType get a device by uuid
var DeviceGetType *graphql.Field

// DeviceListType get a device list by options
var DeviceListType *graphql.Field

// DeviceBindType bind a device
var DeviceBindType *graphql.Field

func init() {
	Device = graphql.NewObject(graphql.ObjectConfig{
		Name: "Device",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"type":        &graphql.Field{Type: graphql.String, Description: "设备类型"},
				"name":        &graphql.Field{Type: graphql.String, Description: "设备名称"},
				"mac":         &graphql.Field{Type: graphql.String, Description: "设备Mac地址"},
				"token":       &graphql.Field{Type: graphql.String, Description: "设备token，用于数据加密"},
				"status":      &graphql.Field{Type: custom.BaseStatus, Description: "基础状态"},
				"id":          &graphql.Field{Type: graphql.Int},
				"uuid":        &graphql.Field{Type: graphql.String, Description: "设备UUID"},
				"user":        &graphql.Field{Type: User, Description: "注册人用户"},
				"description": &graphql.Field{Type: graphql.String, Description: "设备描述，备注"},
				"createdAt":   &graphql.Field{Type: graphql.DateTime},
				"updatedAt":   &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})

	DeviceListType = &graphql.Field{
		Type: graphql.NewList(Device),
		Args: graphql.FieldConfigArgument{
			"type":        GenArg(graphql.String, "设备类型"),
			"namePattern": GenArg(graphql.String, "设备名称模糊匹配"),
			"status":      GenArg(graphql.String, "设备状态"),
			"userUUID":    GenArg(graphql.String, "注册人uuid"),
		},
		Description: "根据条件筛选device列表",
		Resolve:     device.List,
	}

	DeviceGetType = &graphql.Field{
		Type: Device,
		Args: graphql.FieldConfigArgument{
			"uuid": GenArg(graphql.String, "设备UUID", false),
		},
		Description: "使用uuid获取device",
		Resolve:     device.Get,
	}

	DeviceCreateType = &graphql.Field{
		Type: Device,
		Args: graphql.FieldConfigArgument{
			"type":        GenArg(graphql.String, "设备类型", false),
			"name":        GenArg(graphql.String, "设备名称", false),
			"description": GenArg(graphql.String, "描述"),
		},
		Resolve: device.Create,
	}

	DeviceUpdateType = &graphql.Field{
		Type: Device,
		Args: graphql.FieldConfigArgument{
			"uuid":        GenArg(graphql.String, "设备UUID", false),
			"type":        GenArg(graphql.String, "设备类型"),
			"name":        GenArg(graphql.String, "设备名称"),
			"status":      GenArg(custom.BaseStatus, "设备状态"),
			"description": GenArg(graphql.String, "描述"),
		},
		Resolve: device.Update,
	}

	DeviceDeleteType = &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"uuid": GenArg(graphql.String, "设备UUID", false),
		},
		Resolve: device.Delete,
	}

	DeviceBindType = &graphql.Field{
		Type: Device,
		Args: graphql.FieldConfigArgument{
			"token": GenArg(graphql.String, "设备token，用于数据加密", false),
			"mac":   GenArg(graphql.String, "设备Mac地址", false),
		},
		Description: "绑定物理设备Mac地址",
		Resolve:     device.Bind,
	}
}
