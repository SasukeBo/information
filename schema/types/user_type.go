package types

import (
	"github.com/SasukeBo/information/resolvers/user"
	"github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

// User 用户类型
var User graphql.Type

// UserCreate create a user
var UserCreate *graphql.Field

// LoginByPassword login by password
var LoginByPassword *graphql.Field

func init() {
	User = graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":         &graphql.Field{Type: graphql.Int},
				"uuid":       &graphql.Field{Type: graphql.String, Description: "通用唯一标识"},
				"phone":      &graphql.Field{Type: graphql.String, Description: "手机号"},
				"role":       &graphql.Field{Type: Role, Description: "用户角色"},
				"userExtend": &graphql.Field{Type: UserExtend, Description: "用户拓展信息"},
				"status":     &graphql.Field{Type: custom.BaseStatus, Description: "基础状态"},
				"createdAt":  &graphql.Field{Type: graphql.DateTime},
				"updatedAt":  &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})

	UserCreate = &graphql.Field{
		Type: User,
		Args: graphql.FieldConfigArgument{
			"phone":    GenArg(graphql.String, "手机号", false),
			"password": GenArg(graphql.String, "密码", false),
			"smsCode":  GenArg(graphql.String, "验证码", false),
		},
		Resolve: user.Create,
	}

	LoginByPassword = &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"phone":    GenArg(graphql.String, "手机号", false),
			"password": GenArg(graphql.String, "密码", false),
		},
		Resolve: user.LoginByPassword,
	}
}
