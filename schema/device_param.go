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

/*							query
------------------------------------ */

var deviceParamGet = &graphql.Field{
	Type: DeviceParam,
	Args: graphql.FieldConfigArgument{
		"id": GenArg(graphql.Int, "ID", false),
	},
	Description: "获取设备参数信息",
	Resolve:     resolver.GetDeviceParam,
}

var deviceParamList = &graphql.Field{
	Type: graphql.NewList(DeviceParam),
	Args: graphql.FieldConfigArgument{
		"deviceUUID":  GenArg(graphql.String, "设备UUID", false),
		"namePattern": GenArg(graphql.String, "参数名称模糊匹配"),
		"signPattern": GenArg(graphql.String, "参数签名模糊匹配"),
		"type":        GenArg(DeviceParamValueType, "参数值类型"),
	},
	Description: "获取设备参数列表",
	Resolve:     resolver.ListDeviceParam,
}

/*						mutation
------------------------------------ */

var deviceParamCreate = &graphql.Field{
	Type: DeviceParam,
	Args: graphql.FieldConfigArgument{
		"deviceUUID": GenArg(graphql.String, "设备UUID", false),
		"name":       GenArg(graphql.String, "参数名称", false),
		"sign":       GenArg(graphql.String, "参数签名", false),
		"type":       GenArg(DeviceParamValueType, "参数值类型", false),
	},
	Description: "增加设备参数",
	Resolve:     resolver.CreateDeviceParam,
}

var deviceParamUpdate = &graphql.Field{
	Type: DeviceParam,
	Args: graphql.FieldConfigArgument{
		"id":   GenArg(graphql.Int, "ID", false),
		"name": GenArg(graphql.String, "参数名称"),
		"sign": GenArg(graphql.String, "参数签名"),
		"type": GenArg(DeviceParamValueType, "参数值类型"),
	},
	Description: "设备参数更新",
	Resolve:     resolver.UpdateDeviceParam,
}

var deviceParamDelete = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"id": GenArg(graphql.Int, "ID", false),
	},
	Description: "删除设备参数",
	Resolve:     resolver.DeleteDeviceParam,
}
