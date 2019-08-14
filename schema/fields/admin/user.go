package admin

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/admin/user"
	"github.com/SasukeBo/information/schema/scalars"
	"github.com/SasukeBo/information/schema/types"
)

// UserUpdateField update a user
var UserUpdateField = &graphql.Field{
	Type: types.User,
	Args: graphql.FieldConfigArgument{
		"uuid":      fields.GenArg(graphql.String, "用户UUID", false),
		"phone":     fields.GenArg(graphql.String, "用户手机号"),
		"avatarURL": fields.GenArg(graphql.String, "用户头像地址"),
		"roleID":    fields.GenArg(graphql.Int, "用户角色ID"),
		"status":    fields.GenArg(scalars.BaseStatus, "用户基础状态"),
	},
	Resolve: user.Update,
}

// UserDeleteField delete a user
var UserDeleteField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"uuid": fields.GenArg(graphql.String, "用户UUID", false),
	},
	Resolve: user.Delete,
}

// UserListField _
var UserListField = &graphql.Field{
	Type: graphql.NewList(types.User),
	Args: graphql.FieldConfigArgument{
		"limit":           fields.GenArg(graphql.Int, "返回最大条数"),
		"offset":          fields.GenArg(graphql.Int, "返回列表偏移量"),
		"namePattern":     fields.GenArg(graphql.String, "用户名称模糊匹配"),
		"phone":           fields.GenArg(graphql.String, "用户手机号"),
		"email":           fields.GenArg(graphql.String, "用户邮箱"),
		"status":          fields.GenArg(graphql.NewList(scalars.BaseStatus), "用户状态集合，可以写多个状态"),
		"roleNamePattern": fields.GenArg(graphql.String, "用户角色名称模糊匹配"),
		"roleID":          fields.GenArg(graphql.Int, "用户角色ID"),
	},
	Resolve:     user.List,
	Description: "默认按照账号注册时间倒序排列",
}
