package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/*							   types
------------------------------------------ */

var roleType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Role",
	Description: "角色类型",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"roleName":   &graphql.Field{Type: graphql.String, Description: "role name"},
		"status":     &graphql.Field{Type: BaseStatus, Description: "role status, can be default, publish, block and deleted"},
		"isAdmin":    &graphql.Field{Type: graphql.Boolean, Description: "是否为管理员角色，仅管理员角色才可以调用管理员API"},
		"createdAt":  &graphql.Field{Type: graphql.DateTime},
		"updatedAt":  &graphql.Field{Type: graphql.DateTime},
		"privileges": &graphql.Field{Type: graphql.NewList(privilegeType), Description: "role privileges", Resolve: resolver.LoadPrivilege},
	},
})
