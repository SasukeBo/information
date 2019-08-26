package public

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/device"
	"github.com/SasukeBo/information/schema/types"
)

// DeviceChargeCreateField doc false
var DeviceChargeCreateField = &graphql.Field{
	Type: types.DeviceCharge,
	Args: graphql.FieldConfigArgument{
		"uuid":     fields.GenArg(graphql.String, "设备UUID", false),
		"userUUID": fields.GenArg(graphql.String, "指派人UUID", false),
		"privIDs":  fields.GenArg(graphql.NewList(graphql.Int), "权限ids", false),
	},
	Resolve: device.ChargeCreate,
}

// DeviceChargeDeleteField doc false
var DeviceChargeDeleteField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id": fields.GenArg(graphql.Int, "设备指派ID", false),
	},
	Resolve: device.ChargeDelete,
}

// DeviceChargeUpdateField doc false
var DeviceChargeUpdateField = &graphql.Field{
	Type: types.DeviceCharge,
	Args: graphql.FieldConfigArgument{
		"id":       fields.GenArg(graphql.Int, "设备指派ID", false),
		"userUUID": fields.GenArg(graphql.String, "指派人UUID", false),
	},
	Resolve: device.ChargeUpdate,
}

// DeviceChargeGetField doc false
var DeviceChargeGetField = &graphql.Field{
	Type: types.DeviceCharge,
	Args: graphql.FieldConfigArgument{
		"id": fields.GenArg(graphql.Int, "设备负责关系ID", false),
	},
	Description: "通过id获取设备负责关系",
	Resolve:     device.ChargeGet,
}

// DeviceChargeListField doc false
var DeviceChargeListField = &graphql.Field{
	Type: graphql.NewList(types.DeviceCharge),
	Args: graphql.FieldConfigArgument{
		"deviceUUID": fields.GenArg(graphql.String, "设备uuid"),
	},
	Description: `
	查询本人负责的设备或创建的设备的设备负责关系列表，
	可通过设备uuid指定某台设备，但必须是当前用户可访问的设备
	`,
	Resolve: device.ChargeList,
}
