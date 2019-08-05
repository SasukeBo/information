package types

import (
	"github.com/SasukeBo/information/resolvers/device"
	"github.com/graphql-go/graphql"
)

// DeviceParamValue 设备参数值
var DeviceParamValue graphql.Type

// DeviceParamValueListType doc false
var DeviceParamValueListType *graphql.Field

func init() {
	DeviceParamValue = graphql.NewObject(graphql.ObjectConfig{
		Name: "DeviceParamValue",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":          &graphql.Field{Type: graphql.Int},
				"value":       &graphql.Field{Type: graphql.String, Description: "参数值字符串"},
				"deviceParam": &graphql.Field{Type: DeviceParam, Description: "设备参数"},
				"createdAt":   &graphql.Field{Type: graphql.DateTime, Description: "创建时间"},
			}
		}),
	})

	DeviceParamValueListType = &graphql.Field{
		Type: graphql.NewList(DeviceParamValue),
		Args: graphql.FieldConfigArgument{
			"deviceParamID": GenArg(graphql.Int, "参数ID", false),
			"beforeTime":    GenArg(graphql.DateTime, "开始时间"),
			"afterTime":     GenArg(graphql.DateTime, "结束时间"),
		},
		Resolve: device.ParamValueList,
	}
}
