package public

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/device"
	"github.com/SasukeBo/information/schema/types"
)

// DeviceChargerCreateField doc false
var DeviceChargerCreateField = &graphql.Field{
	Type: types.DeviceCharger,
	Args: graphql.FieldConfigArgument{
		"deivceUUID": fields.GenArg(graphql.String, "设备UUID", false),
		"name":       fields.GenArg(graphql.String, "负责人姓名", false),
		"phone":      fields.GenArg(graphql.String, "手机号"),
		"department": fields.GenArg(graphql.String, "部门名称"),
		"jobNumber":  fields.GenArg(graphql.String, "工号"),
	},
	Resolve: device.ChargerCreate,
}

// DeviceChargerDeleteField doc false
var DeviceChargerDeleteField = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"id": fields.GenArg(graphql.Int, "设备负责人ID", false),
	},
	Resolve: device.ChargerDelete,
}

// DeviceChargerUpdateField doc false
var DeviceChargerUpdateField = &graphql.Field{
	Type: types.DeviceCharger,
	Args: graphql.FieldConfigArgument{
		"id":         fields.GenArg(graphql.Int, "设备负责人ID", false),
		"name":       fields.GenArg(graphql.String, "负责人姓名", false),
		"phone":      fields.GenArg(graphql.String, "手机号"),
		"department": fields.GenArg(graphql.String, "部门名称"),
		"jobNumber":  fields.GenArg(graphql.String, "工号"),
	},
	Resolve: device.ChargerUpdate,
}

// DeviceChargerGetField doc false
var DeviceChargerGetField = &graphql.Field{
	Type: types.DeviceCharger,
	Args: graphql.FieldConfigArgument{
		"id": fields.GenArg(graphql.Int, "设备负责人ID", false),
	},
	Description: "通过id获取设备负责人",
	Resolve:     device.ChargerGet,
}

// DeviceChargerListField doc false
var DeviceChargerListField = &graphql.Field{
	Type: graphql.NewList(types.DeviceCharger),
	Args: graphql.FieldConfigArgument{
		"deviceUUID": fields.GenArg(graphql.String, "设备uuid", false),
	},
	Description: `
	查询本人负责的设备或创建的设备的设备负责人列表，
	可通过设备uuid指定某台设备，但必须是当前用户可访问的设备
	`,
	Resolve: device.ChargerList,
}
