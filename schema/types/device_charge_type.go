package types

import (
	"github.com/SasukeBo/information/resolvers/device"
	"github.com/graphql-go/graphql"
)

/*		object		*/

// DeviceCharge 设备类型
var DeviceCharge graphql.Type

/*		types		*/

// DeviceChargeCreateType bind a device
var DeviceChargeCreateType *graphql.Field

// DeviceChargeDeleteType bind a device
var DeviceChargeDeleteType *graphql.Field

// DeviceChargeUpdateType bind a device
var DeviceChargeUpdateType *graphql.Field

func init() {
	/* 								object begin								*/

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

	/* 								objects end								*/

	/* 								types begin								*/

	DeviceChargeCreateType = &graphql.Field{
		Type: DeviceCharge,
		Args: graphql.FieldConfigArgument{
			"uuid":     GenArg(graphql.String, "设备UUID", false),
			"userUuid": GenArg(graphql.String, "指派人UUID", false),
		},
		Resolve: device.Charge,
	}

	DeviceChargeDeleteType = &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"id": GenArg(graphql.Int, "设备指派ID", false),
		},
		Resolve: device.UNCharge,
	}

	DeviceChargeUpdateType = &graphql.Field{
		Type: DeviceCharge,
		Args: graphql.FieldConfigArgument{
			"id":       GenArg(graphql.Int, "设备指派ID", false),
			"userUuid": GenArg(graphql.String, "指派人UUID", false),
		},
		Resolve: device.RECharge,
	}

	/* 							types end								*/
}
