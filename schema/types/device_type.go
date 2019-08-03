package types

import (
	"github.com/SasukeBo/information/resolvers/device"
	"github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

/*		object		*/

// Device 设备类型
var Device graphql.Type

/*		types		*/

// DeviceCreateType create a device
var DeviceCreateType *graphql.Field

// DeviceUpdateType update a device
var DeviceUpdateType *graphql.Field

// DeviceBindType bind a device
// 设备注册后需要绑定物理机台，在客户端绑定机器
var DeviceBindType *graphql.Field

func init() {
	/* 								object begin								*/
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

	/* 								objects end								*/

	/* 								types begin								*/
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
			"status":      GenArg(graphql.String, "设备状态"),
			"description": GenArg(graphql.String, "描述"),
		},
		Resolve: device.Create,
	}

	DeviceBindType = &graphql.Field{
		Type: Device,
		Args: graphql.FieldConfigArgument{
			"token": GenArg(graphql.String, "设备token，用于数据加密", false),
			"mac":   GenArg(graphql.String, "设备Mac地址", false),
		},
		Resolve: device.Bind,
	}

	/* 							types end								*/
}
