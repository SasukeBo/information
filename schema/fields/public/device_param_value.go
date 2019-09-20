package public

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/device"
	"github.com/SasukeBo/information/schema/types"
)

// DeviceParamValueListField doc false
var DeviceParamValueListField = &graphql.Field{
	Type: graphql.NewList(types.DeviceParamValue),
	Args: graphql.FieldConfigArgument{
		"limit":         fields.GenArg(graphql.Int, "返回数量限制"),
		"offset":        fields.GenArg(graphql.Int, "返回记录偏移量"),
		"deviceParamID": fields.GenArg(graphql.Int, "参数ID", false),
		"beforeTime":    fields.GenArg(graphql.DateTime, "开始时间"),
		"afterTime":     fields.GenArg(graphql.DateTime, "结束时间"),
	},
	Resolve: device.ParamValueList,
}

// DeviceParamValueCountField return value count at an interval
var DeviceParamValueCountField = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"deviceUUID": fields.GenArg(graphql.String, "设备UUID", false),
		"beforeTime": fields.GenArg(graphql.DateTime, "开始时间"),
		"afterTime":  fields.GenArg(graphql.DateTime, "结束时间"),
	},
	Resolve:     device.ParamValueCount,
	Description: "查询时间段内设备参数值的最大记录数",
}

// DeviceParamValueAddField _
var DeviceParamValueAddField = &graphql.Field{
	Type: types.DeviceParamValue,
	Args: graphql.FieldConfigArgument{
		"id": fields.GenArg(graphql.Int, "参数值ID", false),
	},
	Resolve: device.ParamValueAdd,
}

// DeviceParamValueHistogramField _
var DeviceParamValueHistogramField = &graphql.Field{
	Type: types.DeviceParamValueHistogram,
	Args: graphql.FieldConfigArgument{
		"paramID":    fields.GenArg(graphql.Int, "参数ID", false),
		"beforeTime": fields.GenArg(graphql.DateTime, "开始时间"),
		"afterTime":  fields.GenArg(graphql.DateTime, "结束时间"),
	},
	Resolve:     device.ParamValueHistogram,
	Description: "获取时间段内参数值直方图数据",
}
