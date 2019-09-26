package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/*							   fields
------------------------------------------ */

// CurrentUserField 获取当前登录的用户
var CurrentUserField = &graphql.Field{
	Type:        User,
	Resolve:     resolver.CurrentUser,
	Description: "返回当前用户信息",
}

// SignInField login by password
var SignInField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"phone":    GenArg(graphql.String, "手机号", false),
		"password": GenArg(graphql.String, "密码", false),
		"remember": GenArg(graphql.Boolean, "记住登录", true, true),
	},
	Resolve:     resolver.SignIn,
	Description: "请求时需要加上 operationName",
}

// SignOutField logout
var SignOutField = &graphql.Field{
	Type:        graphql.String,
	Resolve:     resolver.SignOut,
	Description: "请求时需要加上 operationName",
}
