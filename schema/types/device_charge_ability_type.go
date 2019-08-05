package types

import (
	"github.com/SasukeBo/information/resolvers/device"
	"github.com/graphql-go/graphql"
)

// DeviceChargeAbility 设备负责人权限
var DeviceChargeAbility graphql.Type

// DeviceChargePrivGetType doc false
var DeviceChargePrivGetType *graphql.Field

// DeviceChargePrivListType doc false
var DeviceChargePrivListType *graphql.Field

// DeviceChargePrivCreateType doc false
var DeviceChargePrivCreateType *graphql.Field

// DeviceChargePrivDeleteType doc false
var DeviceChargePrivDeleteType *graphql.Field

func init() {
	DeviceChargeAbility = graphql.NewObject(graphql.ObjectConfig{
		Name: "DeviceChargeAbility",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":           &graphql.Field{Type: graphql.Int},
				"deviceCharge": &graphql.Field{Type: DeviceCharge, Description: "设备负责人关系"},
				"privilege":    &graphql.Field{Type: Privilege, Description: "权限"},
			}
		}),
	})

	DeviceChargePrivGetType = &graphql.Field{
		Type: DeviceChargeAbility,
		Args: graphql.FieldConfigArgument{
			"id": GenArg(graphql.Int, "负责人权限ID", false),
		},
		Resolve: device.ChargePrivGet,
	}

	DeviceChargePrivListType = &graphql.Field{
		Type: graphql.NewList(DeviceChargeAbility),
		Args: graphql.FieldConfigArgument{
			"deviceChargeID": GenArg(graphql.Int, "负责人关系ID", false),
		},
		Description: "根据设备负责人关系获取权限list",
		Resolve:     device.ChargePrivList,
	}

	DeviceChargePrivCreateType = &graphql.Field{
		Type: DeviceChargeAbility,
		Args: graphql.FieldConfigArgument{
			"deviceChargeID": GenArg(graphql.Int, "负责人关系ID", false),
			"privilegeID":    GenArg(graphql.Int, "权限ID", false),
		},
		Description: "为设备负责人添加权限",
		Resolve:     device.ChargePrivCreate,
	}

	DeviceChargePrivDeleteType = &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"id": GenArg(graphql.Int, "设备负责人权限ID", false),
		},
		Description: "删除设备负责人的权限",
		Resolve:     device.ChargePrivDelete,
	}
}
