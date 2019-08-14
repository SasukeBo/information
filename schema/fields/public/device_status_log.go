package public

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/device"
	"github.com/SasukeBo/information/schema/scalars"
	"github.com/SasukeBo/information/schema/types"
)

// DeviceStatusLogListField doc false
var DeviceStatusLogListField = &graphql.Field{
	Type: graphql.NewList(types.DeviceStatusLog),
	Args: graphql.FieldConfigArgument{
		"deviceUUID": fields.GenArg(graphql.String, "设备UUID", false),
		"status":     fields.GenArg(scalars.DeviceStatus, "运行状态"),
		"beforeTime": fields.GenArg(graphql.DateTime, "开始时间"),
		"afterTime":  fields.GenArg(graphql.DateTime, "结束时间"),
	},
	Resolve: device.StatusLogList,
}
