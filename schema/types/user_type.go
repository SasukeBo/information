package types

import (
	"github.com/SasukeBo/information/resolvers/user"
	"github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

// User 用户类型
var User graphql.Type

// UserCreateType create a user
var UserCreateType *graphql.Field

// ResetPasswordType reset user password
var ResetPasswordType *graphql.Field

// UserGet doc false
var UserGet *graphql.Field

// UserList doc false
var UserList *graphql.Field

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

	UserCreateType = &graphql.Field{
		Type: User,
		Args: graphql.FieldConfigArgument{
			"phone":    GenArg(graphql.String, "手机号", false),
			"password": GenArg(graphql.String, "密码", false),
			"smsCode":  GenArg(graphql.String, "验证码", false),
		},
		Resolve:     user.Create,
		Description: "请求时需要加上 operationName",
	}

	ResetPasswordType = &graphql.Field{
		Type: User,
		Args: graphql.FieldConfigArgument{
			"phone":    GenArg(graphql.String, "手机号", false),
			"password": GenArg(graphql.String, "密码", false),
			"smsCode":  GenArg(graphql.String, "验证码", false),
		},
		Resolve: user.ResetPassword,
		Description: `
		未登录状态下修改密码
		请求时需要加上 operationName
		`,
	}

	UserGet = &graphql.Field{
		Type: User,
		Args: graphql.FieldConfigArgument{
			"uuid": GenArg(graphql.String, "用户UUID", false),
		},
		Resolve:     user.Get,
		Description: "使用UUID获取用户",
	}

	/*
		UserList = &graphql.Field{
			Type: User,
			Args: graphql.FieldConfigArgument{
				"namePattern":
				"uuid": GenArg(graphql.String, "用户UUID", false),
			},
			Resolve:     user.List,
			Description: "使用UUID获取用户",
		}
	*/
}
