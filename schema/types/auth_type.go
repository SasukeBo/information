package types

import (
	"github.com/SasukeBo/information/resolvers/auth"
	// "github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

// LoginByPassword login by password
var LoginByPassword *graphql.Field

// Logout the system
var Logout *graphql.Field

func init() {
	LoginByPassword = &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"phone":    GenArg(graphql.String, "手机号", false),
			"password": GenArg(graphql.String, "密码", false),
			"remember": GenArg(graphql.Boolean, "记住登录", true, true),
		},
		Resolve: auth.LoginByPassword,
	}

	Logout = &graphql.Field{
		Type:    graphql.String,
		Resolve: auth.Logout,
	}
}
