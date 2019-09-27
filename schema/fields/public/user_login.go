package public

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/userlogin"
	"github.com/SasukeBo/information/schema/types"
)

// UserLoginListField _
var UserLoginListField = &graphql.Field{
	Type: graphql.NewList(types.UserLogin),
	Args: graphql.FieldConfigArgument{
		"limit":      fields.GenArg(graphql.Int, "单次查询最大返回条数"),
		"offset":     fields.GenArg(graphql.Int, "返回条数的偏移量"),
		"beforeTime": fields.GenArg(graphql.DateTime, "查询 beforeTime 之前的记录"),
		"afterTime":  fields.GenArg(graphql.DateTime, "查询 afterTime 之后的记录"),
	},
	Description: "获取用户登录记录列表，按照时间倒序排列",
	Resolve:     userlogin.List,
}

// UserLoginLastField _
var UserLoginLastField = &graphql.Field{
	Type:        types.UserLogin,
	Description: "获取用户非本session的最近一次登录记录。",
	Resolve:     userlogin.Last,
}

// UserLoginThisField _
var UserLoginThisField = &graphql.Field{
	Type:        types.UserLogin,
	Description: "获取用户此次登录记录。",
	Resolve:     userlogin.This,
}
