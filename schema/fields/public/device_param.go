package public

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/device"
	"github.com/SasukeBo/information/schema/scalars"
	"github.com/SasukeBo/information/schema/types"
)

// DeviceParamCreateField doc false
var DeviceParamCreateField = &graphql.Field{
	Type: types.DeviceParam,
	Args: graphql.FieldConfigArgument{
		"deviceUUID": fields.GenArg(graphql.String, "设备UUID", false),
		"name":     fields.GenArg(graphql.String, "参数名称", false),
		"sign":     fields.GenArg(graphql.String, "参数签名", false),
		"type":     fields.GenArg(scalars.DeviceParamValueType, "参数值类型", false),
	},
	Resolve: device.ParamCreate,
}

// DeviceParamUpdateField doc false
var DeviceParamUpdateField = &graphql.Field{
	Type: types.DeviceParam,
	Args: graphql.FieldConfigArgument{
		"id":   fields.GenArg(graphql.Int, "ID", false),
		"name": fields.GenArg(graphql.String, "参数名称"),
		"sign": fields.GenArg(graphql.String, "参数签名"),
		"type": fields.GenArg(scalars.DeviceParamValueType, "参数值类型"),
	},
	Resolve: device.ParamUpdate,
}

// DeviceParamDeleteField doc false
var DeviceParamDeleteField = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"id": fields.GenArg(graphql.Int, "ID", false),
	},
	Resolve: device.ParamDelete,
}

// DeviceParamGetField doc false
var DeviceParamGetField = &graphql.Field{
	Type: types.DeviceParam,
	Args: graphql.FieldConfigArgument{
		"id": fields.GenArg(graphql.Int, "ID", false),
	},
	Description: "ID获取设备参数",
	Resolve:     device.ParamGet,
}

// DeviceParamListField doc false
var DeviceParamListField = &graphql.Field{
	Type: graphql.NewList(types.DeviceParam),
	Args: graphql.FieldConfigArgument{
		"deviceUUID":  fields.GenArg(graphql.String, "设备UUID", false),
		"namePattern": fields.GenArg(graphql.String, "参数名称模糊匹配"),
		"signPattern": fields.GenArg(graphql.String, "参数签名模糊匹配"),
		"type":        fields.GenArg(scalars.DeviceParamValueType, "参数值类型"),
		"userUUID":    fields.GenArg(graphql.String, "创建人UUID"),
	},
	Resolve:     device.ParamList,
	Description: "按条件查询某设备的参数",
}
