package types

import (
	"github.com/SasukeBo/information/resolvers/auth"
	// "github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

// LoginByPasswordType login by password
var LoginByPasswordType *graphql.Field

// LogoutType the system
var LogoutType *graphql.Field

func init() {
	LoginByPasswordType = &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"phone":    GenArg(graphql.String, "手机号", false),
			"password": GenArg(graphql.String, "密码", false),
			"remember": GenArg(graphql.Boolean, "记住登录", true, true),
		},
		Resolve:     auth.LoginByPassword,
		Description: "请求时需要加上 operationName",
	}

	LogoutType = &graphql.Field{
		Type:        graphql.String,
		Resolve:     auth.Logout,
		Description: "请求时需要加上 operationName",
	}
}
