package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/*							   types
------------------------------------------ */

// RolePriv 角色权限关系类型
var RolePriv = graphql.NewObject(graphql.ObjectConfig{
	Name: "RolePriv",
	Fields: graphql.FieldsThunk(func() graphql.Fields {
		return graphql.Fields{
			"role":      &graphql.Field{Type: Role, Resolve: resolver.LoadRole},
			"privilege": &graphql.Field{Type: Privilege, Resolve: resolver.LoadPrivilege},
		}
	}),
})

// Role 角色类型
var Role = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.FieldsThunk(func() graphql.Fields {
		return graphql.Fields{
			"id":        &graphql.Field{Type: graphql.Int},
			"roleName":  &graphql.Field{Type: graphql.String, Description: "role name"},
			"status":    &graphql.Field{Type: BaseStatus, Description: "role status, can be default, publish, block and deleted"},
			"isAdmin":   &graphql.Field{Type: graphql.Boolean, Description: "是否为管理员角色，仅管理员角色才可以调用管理员API"},
			"createdAt": &graphql.Field{Type: graphql.DateTime},
			"updatedAt": &graphql.Field{Type: graphql.DateTime},
		}
	}),
})

func init() {
	Role.AddFieldConfig("privileges", &graphql.Field{
		Type:        graphql.NewList(Privilege),
		Description: "role privileges",
		Resolve:     resolver.LoadPrivilege,
	})
}
