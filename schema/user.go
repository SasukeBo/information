package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/* 					 types
------------------------------ */

// User 用户类型
var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.FieldsThunk(func() graphql.Fields {
		return graphql.Fields{
			"avatarURL": &graphql.Field{Type: graphql.String, Description: "头像链接"},
			"createdAt": &graphql.Field{Type: graphql.DateTime},
			"email":     &graphql.Field{Type: graphql.String, Description: "邮箱"},
			"id":        &graphql.Field{Type: graphql.Int},
			"name":      &graphql.Field{Type: graphql.String, Description: "姓名"},
			"phone":     &graphql.Field{Type: graphql.String, Description: "手机号"},
			"status":    &graphql.Field{Type: BaseStatus, Description: "基础状态"},
			"updatedAt": &graphql.Field{Type: graphql.DateTime},
			"uuid":      &graphql.Field{Type: graphql.String, Description: "通用唯一标识"},
		}
	}),
})

func init() {
	User.AddFieldConfig("role", &graphql.Field{Type: Role, Description: "用户角色", Resolve: resolver.LoadRole})
}

/* 					 fields
------------------------------ */

// SignUpField _
var SignUpField = &graphql.Field{
	Type: User,
	Args: graphql.FieldConfigArgument{
		"phone":    GenArg(graphql.String, "手机号", false),
		"name":     GenArg(graphql.String, "姓名", false),
		"password": GenArg(graphql.String, "密码", false),
		"smsCode":  GenArg(graphql.String, "验证码", false),
	},
	Resolve:     resolver.SignUp,
	Description: "请求时需要加上 operationName",
}

// ResetPasswordField _
var ResetPasswordField = &graphql.Field{
	Type: User,
	Args: graphql.FieldConfigArgument{
		"phone":    GenArg(graphql.String, "手机号", false),
		"password": GenArg(graphql.String, "密码", false),
		"smsCode":  GenArg(graphql.String, "验证码", false),
	},
	Resolve: resolver.ResetPassword,
	Description: `
    未登录状态下修改密码
    请求时需要加上 operationName
    `,
}

// UserGetField _
var UserGetField = &graphql.Field{
	Type: User,
	Args: graphql.FieldConfigArgument{
		"uuid": GenArg(graphql.String, "用户UUID", false),
	},
	Resolve:     resolver.GetUser,
	Description: "使用UUID获取用户",
}

// UserListField _
var UserListField = &graphql.Field{
	Type: graphql.NewList(User),
	Args: graphql.FieldConfigArgument{
		"namePattern": GenArg(graphql.String, "用户名称模糊匹配"),
		"phone":       GenArg(graphql.String, "用户手机号"),
		"email":       GenArg(graphql.String, "用户邮箱"),
	},
	Resolve:     resolver.ListUser,
	Description: "按条件查询用户列表，如果没有给出查询条件，返回空列表",
}

// UserUpdateField _
var UserUpdateField = &graphql.Field{
	Type: User,
	Args: graphql.FieldConfigArgument{
		"avatarURL":   GenArg(graphql.String, "头像链接"),
		"password":    GenArg(graphql.String, "当前密码"),
		"newPassword": GenArg(graphql.String, "新密码"),
		"newPhone":    GenArg(graphql.String, "新手机号"),
		"smsCode":     GenArg(graphql.String, "短信验证码"),
	},
	Resolve: resolver.UpdateUser,
	Description: `用户登录状态下修改账号信息，注意：
		1. 修改密码和修改手机号需要提供当前密码
		2. 修改手机号需要提供短信验证码
	`,
}
