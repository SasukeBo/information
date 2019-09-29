package schemaadmin

/*
import (
	resolver "github.com/SasukeBo/information/resolveradmin"
	"github.com/SasukeBo/information/schema"
	"github.com/graphql-go/graphql"
)

// UserUpdateField update a user
var UserUpdateField = &graphql.Field{
	Type: schema.User,
	Args: graphql.FieldConfigArgument{
		"uuid":      schema.GenArg(graphql.String, "用户UUID", false),
		"phone":     schema.GenArg(graphql.String, "用户手机号"),
		"avatarURL": schema.GenArg(graphql.String, "用户头像地址"),
		"roleID":    schema.GenArg(graphql.Int, "用户角色ID"),
		"status":    schema.GenArg(schema.BaseStatus, "用户基础状态"),
	},
	Resolve: resolver.UpdateUser,
}

// UserDeleteField delete a user
var UserDeleteField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"uuid": schema.GenArg(graphql.String, "用户UUID", false),
	},
	Resolve: resolver.DeleteUser,
}

// UserListField _
var UserListField = &graphql.Field{
	Type: graphql.NewList(schema.User),
	Args: graphql.FieldConfigArgument{
		"limit":           schema.GenArg(graphql.Int, "返回最大条数"),
		"offset":          schema.GenArg(graphql.Int, "返回列表偏移量"),
		"namePattern":     schema.GenArg(graphql.String, "用户名称模糊匹配"),
		"phone":           schema.GenArg(graphql.String, "用户手机号"),
		"email":           schema.GenArg(graphql.String, "用户邮箱"),
		"status":          schema.GenArg(graphql.NewList(schema.BaseStatus), "用户状态集合，可以写多个状态"),
		"roleNamePattern": schema.GenArg(graphql.String, "用户角色名称模糊匹配"),
		"roleID":          schema.GenArg(graphql.Int, "用户角色ID"),
	},
	Description: "默认按照账号注册时间倒序排列",
	Resolve:     resolver.ListUser,
}
*/
