package admin

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/admin/ulogin"
	"github.com/SasukeBo/information/schema/types"
)

// UserLoginListField _
var UserLoginListField = &graphql.Field{
	Type: graphql.NewList(types.UserLogin),
	Args: graphql.FieldConfigArgument{
		"limit":      fields.GenArg(graphql.Int, "最大返回条数"),
		"offset":     fields.GenArg(graphql.Int, "返回记录偏移量"),
		"userUUID":   fields.GenArg(graphql.String, "用户UUID"),
		"remoteIP":   fields.GenArg(graphql.String, "登录IP，格式如 127.0.0.1"),
		"beforeTime": fields.GenArg(graphql.DateTime, "返回beforeTime时间前的记录"),
		"afterTime":  fields.GenArg(graphql.DateTime, "返回afterTime时间后的记录"),
	},
	Description: "获取用户登录记录，按照时间倒序返回列表",
	Resolve:     ulogin.List,
}
