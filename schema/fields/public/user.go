package public

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/user"
	"github.com/SasukeBo/information/schema/types"
)

// UserCreateField create a user
var UserCreateField *graphql.Field

// ResetPasswordField reset user password
var ResetPasswordField *graphql.Field

// UserGetField doc false
var UserGetField *graphql.Field

// UserListField doc false
var UserListField *graphql.Field

func init() {

	UserCreateField = &graphql.Field{
		Type: types.User,
		Args: graphql.FieldConfigArgument{
			"phone":    fields.GenArg(graphql.String, "手机号", false),
			"password": fields.GenArg(graphql.String, "密码", false),
			"smsCode":  fields.GenArg(graphql.String, "验证码", false),
		},
		Resolve:     user.Create,
		Description: "请求时需要加上 operationName",
	}

	ResetPasswordField = &graphql.Field{
		Type: types.User,
		Args: graphql.FieldConfigArgument{
			"phone":    fields.GenArg(graphql.String, "手机号", false),
			"password": fields.GenArg(graphql.String, "密码", false),
			"smsCode":  fields.GenArg(graphql.String, "验证码", false),
		},
		Resolve: user.ResetPassword,
		Description: `
		未登录状态下修改密码
		请求时需要加上 operationName
		`,
	}

	UserGetField = &graphql.Field{
		Type: types.User,
		Args: graphql.FieldConfigArgument{
			"uuid": fields.GenArg(graphql.String, "用户UUID", false),
		},
		Resolve:     user.Get,
		Description: "使用UUID获取用户",
	}

	/*
		UserListField = &graphql.Field{
			Type: types.User,
			Args: graphql.FieldConfigArgument{
				"namePattern":
				"uuid": GenArg(graphql.String, "用户UUID", false),
			},
			Resolve:     user.List,
			Description: "使用UUID获取用户",
		}
	*/
}
