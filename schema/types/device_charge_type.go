package types

import (
	"github.com/SasukeBo/information/resolvers/device"
	"github.com/graphql-go/graphql"
)

// DeviceCharge 设备类型
var DeviceCharge graphql.Type

// DeviceChargeCreateType doc false
var DeviceChargeCreateType *graphql.Field

// DeviceChargeDeleteType doc false
var DeviceChargeDeleteType *graphql.Field

// DeviceChargeUpdateType doc false
var DeviceChargeUpdateType *graphql.Field

// DeviceChargeGetType doc false
var DeviceChargeGetType *graphql.Field

// DeviceChargeListType doc false
var DeviceChargeListType *graphql.Field

func init() {
	DeviceCharge = graphql.NewObject(graphql.ObjectConfig{
		Name: "DeviceCharge",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":        &graphql.Field{Type: graphql.Int},
				"device":    &graphql.Field{Type: Device, Description: "设备"},
				"user":      &graphql.Field{Type: User, Description: "负责人"},
				"createdAt": &graphql.Field{Type: graphql.DateTime},
				"updatedAt": &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})

	DeviceChargeCreateType = &graphql.Field{
		Type: DeviceCharge,
		Args: graphql.FieldConfigArgument{
			"uuid":     GenArg(graphql.String, "设备UUID", false),
			"userUuid": GenArg(graphql.String, "指派人UUID", false),
		},
		Resolve: device.ChargeCreate,
	}

	DeviceChargeDeleteType = &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"id": GenArg(graphql.Int, "设备指派ID", false),
		},
		Resolve: device.ChargeDelete,
	}

	DeviceChargeUpdateType = &graphql.Field{
		Type: DeviceCharge,
		Args: graphql.FieldConfigArgument{
			"id":       GenArg(graphql.Int, "设备指派ID", false),
			"userUuid": GenArg(graphql.String, "指派人UUID", false),
		},
		Resolve: device.ChargeUpdate,
	}

	DeviceChargeGetType = &graphql.Field{
		Type: DeviceCharge,
		Args: graphql.FieldConfigArgument{
			"id": GenArg(graphql.Int, "设备负责关系ID", false),
		},
		Description: "通过id获取设备负责关系",
		Resolve:     device.ChargeGet,
	}

	DeviceChargeListType = &graphql.Field{
		Type: graphql.NewList(DeviceCharge),
		Args: graphql.FieldConfigArgument{
			"userUUID":   GenArg(graphql.String, "设备负责人uuid"),
			"deviceUUID": GenArg(graphql.String, "设备uuid"),
		},
		Description: "通过负责人uuid或设备uuid获取设备负责关系列表",
		Resolve:     device.ChargeList,
	}
}
