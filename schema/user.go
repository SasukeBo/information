package schema

import (
	// "fmt"
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/* 					 types
------------------------------ */

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "用户",
	Fields: graphql.Fields{
		"avatarURL": &graphql.Field{Type: graphql.String, Description: "头像链接"},
		"createdAt": &graphql.Field{Type: graphql.DateTime},
		"email":     &graphql.Field{Type: graphql.String, Description: "邮箱"},
		"id":        &graphql.Field{Type: graphql.Int},
		"name":      &graphql.Field{Type: graphql.String, Description: "姓名"},
		"phone":     &graphql.Field{Type: graphql.String, Description: "手机号"},
		"status":    &graphql.Field{Type: BaseStatus, Description: "基础状态"},
		"updatedAt": &graphql.Field{Type: graphql.DateTime},
		"role":      &graphql.Field{Type: roleType, Description: "用户角色", Resolve: resolver.LoadRole},
	},
})

/* 				 	query
------------------------------ */

var userGet = &graphql.Field{
	Type: userType,
	Args: graphql.FieldConfigArgument{
		"uuid": GenArg(graphql.String, "用户UUID", false),
	},
	Resolve:     resolver.GetUser,
	Description: "使用UUID获取用户",
}

var userList = &graphql.Field{
	Type: graphql.NewList(userType),
	Args: graphql.FieldConfigArgument{
		"namePattern": GenArg(graphql.String, "用户名称模糊匹配"),
		"phone":       GenArg(graphql.String, "用户手机号"),
		"email":       GenArg(graphql.String, "用户邮箱"),
	},
	Resolve:     resolver.ListUser,
	Description: "按条件查询用户列表，如果没有给出查询条件，返回空列表",
}

/* 					mutation
------------------------------ */

var signUp = &graphql.Field{
	Type: userType,
	Args: graphql.FieldConfigArgument{
		"phone":    GenArg(graphql.String, "手机号", false),
		"name":     GenArg(graphql.String, "姓名", false),
		"password": GenArg(graphql.String, "密码", false),
		"smsCode":  GenArg(graphql.String, "验证码", false),
	},
	Description: `#### 注册账号
**注意** 需要提供operationName`,
	Resolve: resolver.SignUp,
}

var resetPassword = &graphql.Field{
	Type: userType,
	Args: graphql.FieldConfigArgument{
		"phone":    GenArg(graphql.String, "手机号", false),
		"password": GenArg(graphql.String, "密码", false),
		"smsCode":  GenArg(graphql.String, "验证码", false),
	},
	Description: `#### 找回密码
非登录状态下修改密码
**注意** 需要提供operationName`,
	Resolve: resolver.ResetPassword,
}

var userUpdate = &graphql.Field{
	Type: userType,
	Args: graphql.FieldConfigArgument{
		"avatarURL":   GenArg(graphql.String, "头像链接"),
		"password":    GenArg(graphql.String, "当前密码"),
		"newPassword": GenArg(graphql.String, "新密码"),
		"newPhone":    GenArg(graphql.String, "新手机号"),
		"smsCode":     GenArg(graphql.String, "短信验证码"),
	},
	Description: `#### 修改用户信息
登录状态下修改用户账号信息。
**注意**
- 修改密码和修改手机号需要提供当前密码
- 修改手机号需要提供短信验证码`,
	Resolve: resolver.UpdateUser,
}
