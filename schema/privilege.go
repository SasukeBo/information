package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/* 					 types
------------------------------ */

// Privilege 权限类型
var Privilege = graphql.NewObject(graphql.ObjectConfig{
	Name: "Privilege",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.Int},
		"name":     &graphql.Field{Type: graphql.String, Description: "权限名称"},
		"sign":     &graphql.Field{Type: graphql.String, Description: "权限签名"},
		"privType": &graphql.Field{Type: graphql.Int, Description: "权限类型"},
	},
})

/* 				query fields
------------------------------ */

// PrivilegeListField get list of privilege
var privilegeListField = &graphql.Field{
	Type: graphql.NewList(Privilege),
	Args: graphql.FieldConfigArgument{
		"privType":    GenArg(PrivType, "权限类型"),
		"namePattern": GenArg(graphql.String, "权限名称模糊匹配"),
	},
	Resolve: resolver.ListPrivilege,
}
