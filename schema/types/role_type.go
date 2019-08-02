package types

import (
	"github.com/SasukeBo/information/resolvers/role"
	"github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

// Role 用户类型
var Role graphql.Type

// RoleCreate create a role
var RoleCreate *graphql.Field

// RoleUpdate create a role
var RoleUpdate *graphql.Field

// RoleGet get role by id
var RoleGet *graphql.Field

// RoleGetByName get role by name
var RoleGetByName *graphql.Field

func init() {
	Role = graphql.NewObject(graphql.ObjectConfig{
		Name: "Role",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":        &graphql.Field{Type: graphql.Int},
				"roleName":  &graphql.Field{Type: graphql.String, Description: "role name"},
				"status":    &graphql.Field{Type: custom.BaseStatus, Description: "role status, can be default, publish, block and deleted"},
				"rolePrivs": &graphql.Field{Type: graphql.NewList(RolePriv), Description: "role and privilege relationship"},
				"createdAt": &graphql.Field{Type: graphql.DateTime},
				"updatedAt": &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})

	RoleCreate = &graphql.Field{
		Type: Role,
		Args: graphql.FieldConfigArgument{
			"roleName": GenArg(graphql.String, "角色名称", false),
		},
		Resolve: role.Create,
	}

	RoleUpdate = &graphql.Field{
		Type: Role,
		Args: graphql.FieldConfigArgument{
			"id":       GenArg(graphql.Int, "角色ID", false),
			"roleName": GenArg(graphql.String, "角色名称"),
			"status":   gBaseStatus,
		},
		Resolve: role.Update,
	}

	RoleGet = &graphql.Field{
		Type: Role,
		Args: graphql.FieldConfigArgument{
			"id": GenArg(graphql.Int, "角色ID", false),
		},
		Resolve: role.Get,
	}

	RoleGetByName = &graphql.Field{
		Type: Role,
		Args: graphql.FieldConfigArgument{
			"roleName": GenArg(graphql.String, "角色名称", false),
		},
		Resolve: role.GetByName,
	}
}
