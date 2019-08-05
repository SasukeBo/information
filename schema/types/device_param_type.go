package types

import (
	"github.com/SasukeBo/information/resolvers/device"
	"github.com/graphql-go/graphql"
)

// DeviceParam 设备参数
var DeviceParam graphql.Type

// DeviceParamCreateType doc false
var DeviceParamCreateType *graphql.Field

// DeviceParamUpdateType doc false
var DeviceParamUpdateType *graphql.Field

// DeviceParamDeleteType doc false
var DeviceParamDeleteType *graphql.Field

// DeviceParamGetType doc false
var DeviceParamGetType *graphql.Field

// DeviceParamListType doc false
var DeviceParamListType *graphql.Field

func init() {
	DeviceParam = graphql.NewObject(graphql.ObjectConfig{
		Name: "DeviceParam",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":        &graphql.Field{Type: graphql.Int},
				"name":      &graphql.Field{Type: graphql.String, Description: "参数名称"},
				"sign":      &graphql.Field{Type: graphql.String, Description: "参数签名"},
				"type":      &graphql.Field{Type: graphql.String, Description: "参数值类型"},
				"author":    &graphql.Field{Type: User, Description: "创建人"},
				"createdAt": &graphql.Field{Type: graphql.DateTime, Description: "创建时间"},
			}
		}),
	})

	DeviceParamCreateType = &graphql.Field{
		Type: DeviceParam,
		Args: graphql.FieldConfigArgument{
			"name": GenArg(graphql.String, "参数名称", false),
			"sign": GenArg(graphql.String, "参数签名", false),
			"type": GenArg(graphql.String, "参数值类型", false),
		},
		Resolve: device.ParamCreate,
	}

	DeviceParamUpdateType = &graphql.Field{
		Type: DeviceParam,
		Args: graphql.FieldConfigArgument{
			"id":   GenArg(graphql.Int, "ID", false),
			"name": GenArg(graphql.String, "参数名称"),
			"sign": GenArg(graphql.String, "参数签名"),
			"type": GenArg(graphql.String, "参数值类型，可以为 number/string/boolean"),
		},
		Resolve: device.ParamUpdate,
	}

	DeviceParamDeleteType = &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"id": GenArg(graphql.Int, "ID", false),
		},
		Resolve: device.ParamDelete,
	}

	DeviceParamGetType = &graphql.Field{
		Type: DeviceParam,
		Args: graphql.FieldConfigArgument{
			"id": GenArg(graphql.Int, "ID", false),
		},
		Description: "ID获取设备参数",
		Resolve:     device.ParamGet,
	}

	DeviceParamListType = &graphql.Field{
		Type: graphql.NewList(DeviceParam),
		Args: graphql.FieldConfigArgument{
			"namePattern": GenArg(graphql.String, "参数名称模糊匹配"),
			"signPattern": GenArg(graphql.String, "参数签名模糊匹配"),
			"type":        GenArg(graphql.String, "参数值类型"),
			"userUUID":    GenArg(graphql.String, "创建人UUID"),
		},
		Resolve: device.ParamList,
	}
}
