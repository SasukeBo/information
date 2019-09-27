package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/*							   types
------------------------------------------ */

// DeviceParamValue 设备参数值类型
var DeviceParamValue = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeviceParamValue",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"value":     &graphql.Field{Type: graphql.String, Description: "参数值字符串"},
		"createdAt": &graphql.Field{Type: graphql.DateTime, Description: "创建时间"},
	},
})

func init() {
	DeviceParamValue.AddFieldConfig(
		"deviceParam",
		&graphql.Field{Type: DeviceParam, Description: "设备参数", Resolve: resolver.LoadDeviceParam},
	)
}

/*							response
------------------------------------------ */

var deviceParamValueHistogramResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeviceParamValueHistogram",
	Fields: graphql.Fields{
		"category": &graphql.Field{Type: graphql.NewList(graphql.String), Description: "直方图x轴category"},
		"serie":    &graphql.Field{Type: graphql.NewList(graphql.Int), Description: "直方图serie data"},
	},
})

/*							 query fields
------------------------------------------ */

var deviceParamValueListField = &graphql.Field{
	Type: graphql.NewList(DeviceParamValue),
	Args: graphql.FieldConfigArgument{
		"limit":         GenArg(graphql.Int, "返回数量限制"),
		"offset":        GenArg(graphql.Int, "返回记录偏移量"),
		"deviceParamID": GenArg(graphql.Int, "参数ID", false),
		"beforeTime":    GenArg(graphql.DateTime, "开始时间"),
		"afterTime":     GenArg(graphql.DateTime, "结束时间"),
	},
	Resolve:     resolver.ListDeviceParamValue,
	Description: "获取设备参数值列表",
}

var deviceParamValueCountField = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"deviceUUID": GenArg(graphql.String, "设备UUID", false),
		"beforeTime": GenArg(graphql.DateTime, "开始时间"),
		"afterTime":  GenArg(graphql.DateTime, "结束时间"),
	},
	Resolve:     resolver.CountDeviceParamValue,
	Description: "查询时间段内设备参数值的最大记录数",
}

var deviceParamValueHistogramField = &graphql.Field{
	Type: deviceParamValueHistogramResponse,
	Args: graphql.FieldConfigArgument{
		"paramID":    GenArg(graphql.Int, "参数ID", false),
		"beforeTime": GenArg(graphql.DateTime, "开始时间"),
		"afterTime":  GenArg(graphql.DateTime, "结束时间"),
	},
	Resolve:     resolver.DeviceParamValueHistogram,
	Description: "获取时间段内参数值直方图数据",
}

/*						subscription fields
------------------------------------------ */

var deviceParamValueSubField = &graphql.Field{
	Type: DeviceParamValue,
	Args: graphql.FieldConfigArgument{
		"id": GenArg(graphql.Int, "值记录ID", false),
	},
	Resolve:     resolver.GetDeviceParamValue,
	Description: "设备参数值订阅",
}
