package public

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/user"
	"github.com/SasukeBo/information/schema/types"
)

// UserCreateField _
var UserCreateField = &graphql.Field{
	Type: types.User,
	Args: graphql.FieldConfigArgument{
		"phone":    fields.GenArg(graphql.String, "手机号", false),
		"name":     fields.GenArg(graphql.String, "姓名", false),
		"password": fields.GenArg(graphql.String, "密码", false),
		"smsCode":  fields.GenArg(graphql.String, "验证码", false),
	},
	Resolve:     user.Create,
	Description: "请求时需要加上 operationName",
}

// ResetPasswordField _
var ResetPasswordField = &graphql.Field{
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

// UserGetField _
var UserGetField = &graphql.Field{
	Type: types.User,
	Args: graphql.FieldConfigArgument{
		"uuid": fields.GenArg(graphql.String, "用户UUID", false),
	},
	Resolve:     user.Get,
	Description: "使用UUID获取用户",
}

// UserListField _
var UserListField = &graphql.Field{
	Type: graphql.NewList(types.User),
	Args: graphql.FieldConfigArgument{
		"namePattern": fields.GenArg(graphql.String, "用户名称模糊匹配"),
		"phone":       fields.GenArg(graphql.String, "用户手机号"),
		"email":       fields.GenArg(graphql.String, "用户邮箱"),
	},
	Resolve:     user.List,
	Description: "按条件查询用户列表，如果没有给出查询条件，返回空列表",
}

// UserUpdateAvatarField _
var UserUpdateAvatarField = &graphql.Field{
	Type: types.User,
	Args: graphql.FieldConfigArgument{
		"avatarURL": fields.GenArg(graphql.String, "头像链接", false),
	},
	Resolve:     user.UpdateAvatar,
	Description: "用户修改头像",
}

// UserUpdatePasswordField _
var UserUpdatePasswordField = &graphql.Field{
	Type: types.User,
	Args: graphql.FieldConfigArgument{
		"oldPassword": fields.GenArg(graphql.String, "旧密码", false),
		"newPassword": fields.GenArg(graphql.String, "新密码", false),
	},
	Resolve:     user.UpdatePassword,
	Description: "登录状态下，用户修改密码",
}

// UserUpdatePhoneField _
var UserUpdatePhoneField = &graphql.Field{
	Type: types.User,
	Args: graphql.FieldConfigArgument{
		"password": fields.GenArg(graphql.String, "登录密码", false),
		"newPhone": fields.GenArg(graphql.String, "新手机号", false),
		"smsCode":  fields.GenArg(graphql.String, "短信验证码", false),
	},
	Resolve:     user.UpdatePhone,
	Description: "登录状态下用户修改手机号",
}
