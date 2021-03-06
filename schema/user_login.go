package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/* 					 types
------------------------------ */

var userLoginType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserLogin",
	Description: "用户登录类型",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"userAgent": &graphql.Field{Type: graphql.String, Description: "UA"},
		"remoteIP":  &graphql.Field{Type: graphql.String, Description: "头像链接"},
		"logout":    &graphql.Field{Type: graphql.Boolean, Description: "用户角色"},
		"createdAt": &graphql.Field{Type: graphql.DateTime},
		"updatedAt": &graphql.Field{Type: graphql.DateTime},
	},
})

func init() {
	userLoginType.AddFieldConfig("user", &graphql.Field{
		Type:        userType,
		Description: "用户",
		Resolve:     resolver.LoadUser,
	})
}

/* 					query
------------------------------ */

var userLoginList = &graphql.Field{
	Type: graphql.NewList(userLoginType),
	Args: graphql.FieldConfigArgument{
		"limit":      GenArg(graphql.Int, "单次查询最大返回条数"),
		"offset":     GenArg(graphql.Int, "返回条数的偏移量"),
		"beforeTime": GenArg(graphql.DateTime, "查询 beforeTime 之前的记录"),
		"afterTime":  GenArg(graphql.DateTime, "查询 afterTime 之后的记录"),
	},
	Description: "获取用户登录记录列表，按照时间倒序排列",
	Resolve:     resolver.ListUserLogin,
}

var userLoginLast = &graphql.Field{
	Type:        userLoginType,
	Description: "获取用户非本session的最近一次登录记录。",
	Resolve:     resolver.LastUserLogin,
}

var userLoginThis = &graphql.Field{
	Type:        userLoginType,
	Description: "获取用户此次登录记录。",
	Resolve:     resolver.ThisUserLogin,
}

/* 			mutation fields
------------------------------ */
