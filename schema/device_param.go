package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/*							   types
------------------------------------------ */

// DeviceParam 设备参数类型
var DeviceParam = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeviceParam",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"name":      &graphql.Field{Type: graphql.String, Description: "参数名称"},
		"sign":      &graphql.Field{Type: graphql.String, Description: "参数签名"},
		"createdAt": &graphql.Field{Type: graphql.DateTime, Description: "创建时间"},
	},
})

func init() {
	DeviceParam.AddFieldConfig("type", &graphql.Field{Type: DeviceParamValueType, Description: "参数值类型"})
	DeviceParam.AddFieldConfig("device", &graphql.Field{Type: DeviceType, Description: "设备", Resolve: resolver.LoadDevice})
	DeviceParam.AddFieldConfig(
		"values",
		&graphql.Field{
			Type:        graphql.NewList(DeviceParamValue),
			Description: "参数值记录",
			Args: graphql.FieldConfigArgument{
				"limit":      GenArg(graphql.Int, "最大数量"),
				"offset":     GenArg(graphql.Int, "偏移量"),
				"beforeTime": GenArg(graphql.DateTime, "开始时间"),
				"afterTime":  GenArg(graphql.DateTime, "结束时间"),
			},
			Resolve: resolver.LoadDeviceParamValue,
		},
	)
}

/*							query fields
------------------------------------------ */

var deviceParamGetField = &graphql.Field{
	Type: DeviceParam,
	Args: graphql.FieldConfigArgument{
		"id": GenArg(graphql.Int, "ID", false),
	},
	Description: "ID获取设备参数",
	Resolve:     resolver.GetDeviceParam,
}

var deviceParamListField = &graphql.Field{
	Type: graphql.NewList(DeviceParam),
	Args: graphql.FieldConfigArgument{
		"deviceUUID":  GenArg(graphql.String, "设备UUID", false),
		"namePattern": GenArg(graphql.String, "参数名称模糊匹配"),
		"signPattern": GenArg(graphql.String, "参数签名模糊匹配"),
		"type":        GenArg(DeviceParamValueType, "参数值类型"),
	},
	Resolve:     resolver.ListDeviceParam,
	Description: "按条件查询某设备的参数",
}

/*						mutation fields
------------------------------------------ */

var deviceParamCreateField = &graphql.Field{
	Type: DeviceParam,
	Args: graphql.FieldConfigArgument{
		"deviceUUID": GenArg(graphql.String, "设备UUID", false),
		"name":       GenArg(graphql.String, "参数名称", false),
		"sign":       GenArg(graphql.String, "参数签名", false),
		"type":       GenArg(DeviceParamValueType, "参数值类型", false),
	},
	Resolve: resolver.CreateDeviceParam,
}

var deviceParamUpdateField = &graphql.Field{
	Type: DeviceParam,
	Args: graphql.FieldConfigArgument{
		"id":   GenArg(graphql.Int, "ID", false),
		"name": GenArg(graphql.String, "参数名称"),
		"sign": GenArg(graphql.String, "参数签名"),
		"type": GenArg(DeviceParamValueType, "参数值类型"),
	},
	Resolve: resolver.UpdateDeviceParam,
}

var deviceParamDeleteField = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"id": GenArg(graphql.Int, "ID", false),
	},
	Resolve: resolver.DeleteDeviceParam,
}
