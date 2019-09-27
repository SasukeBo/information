package schemaadmin

import (
	resolver "github.com/SasukeBo/information/resolveradmin"
	"github.com/SasukeBo/information/schema"
	"github.com/graphql-go/graphql"
)

// DeviceGetField get a device
var DeviceGetField = &graphql.Field{
	Type: schema.DeviceType,
	Args: graphql.FieldConfigArgument{
		"uuid": schema.GenArg(graphql.String, "设备UUID", false),
	},
	Resolve: resolver.GetDevice,
}

// DeviceListField _
var DeviceListField = &graphql.Field{
	Type: graphql.NewList(schema.DeviceType),
	Args: graphql.FieldConfigArgument{
		"limit":       schema.GenArg(graphql.Int, "最大获取条数"),
		"offset":      schema.GenArg(graphql.Int, "偏移量"),
		"type":        schema.GenArg(graphql.String, "设备类型"),
		"namePattern": schema.GenArg(graphql.String, "设备名称模糊匹配"),
		"status":      schema.GenArg(graphql.NewList(schema.BaseStatus), "设备状态列表"),
		"userName":    schema.GenArg(graphql.String, "设备注册人姓名"),
		"userUUID":    schema.GenArg(graphql.String, "设备创建人UUID"),
		"chargerName": schema.GenArg(graphql.String, "设备负责人姓名"),
		"chargerUUID": schema.GenArg(graphql.String, "设备负责人UUID"),
	},
	Description: `
	管理员获取设备列表，返回记录列表按照设备创建时间倒序排列，
	按照时间的升降序排列控制交由前端处理。
	`,
	Resolve: resolver.ListDevice,
}

// DeviceUpdateField _
var DeviceUpdateField = &graphql.Field{
	Type: schema.DeviceType,
	Args: graphql.FieldConfigArgument{
		"uuid":        schema.GenArg(graphql.String, "设备UUID", false),
		"name":        schema.GenArg(graphql.String, "设备名称"),
		"mac":         schema.GenArg(graphql.String, "设备Mac地址"),
		"status":      schema.GenArg(schema.BaseStatus, "设备状态"),
		"description": schema.GenArg(graphql.String, "设备描述"),
	},
	Resolve: resolver.UpdateDevice,
}

// DeviceDeleteField _
var DeviceDeleteField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"uuid": schema.GenArg(graphql.String, "设备UUID", false),
	},
	Resolve: resolver.DeleteDevice,
}
