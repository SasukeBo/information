package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/* 					 types
------------------------------ */

// UserLogin 用户登录类型
var UserLogin *graphql.Object

func init() {
	UserLogin = graphql.NewObject(graphql.ObjectConfig{
		Name: "UserLogin",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":        &graphql.Field{Type: graphql.Int},
				"userAgent": &graphql.Field{Type: graphql.String, Description: "UA"},
				"user":      &graphql.Field{Type: User, Description: "用户", Resolve: resolver.LoadUser},
				"remoteIP":  &graphql.Field{Type: graphql.String, Description: "头像链接"},
				"logout":    &graphql.Field{Type: graphql.Boolean, Description: "用户角色"},
				"createdAt": &graphql.Field{Type: graphql.DateTime},
				"updatedAt": &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})
}

/* 					 fields
------------------------------ */

// UserLoginListField _
var UserLoginListField = &graphql.Field{
	Type: graphql.NewList(UserLogin),
	Args: graphql.FieldConfigArgument{
		"limit":      GenArg(graphql.Int, "单次查询最大返回条数"),
		"offset":     GenArg(graphql.Int, "返回条数的偏移量"),
		"beforeTime": GenArg(graphql.DateTime, "查询 beforeTime 之前的记录"),
		"afterTime":  GenArg(graphql.DateTime, "查询 afterTime 之后的记录"),
	},
	Description: "获取用户登录记录列表，按照时间倒序排列",
	Resolve:     resolver.ListUserLogin,
}

// UserLoginLastField _
var UserLoginLastField = &graphql.Field{
	Type:        UserLogin,
	Description: "获取用户非本session的最近一次登录记录。",
	Resolve:     resolver.LastUserLogin,
}

// UserLoginThisField _
var UserLoginThisField = &graphql.Field{
	Type:        UserLogin,
	Description: "获取用户此次登录记录。",
	Resolve:     resolver.ThisUserLogin,
}
