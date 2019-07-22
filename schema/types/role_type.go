package types

import (
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

// Role 用户类型
var Role graphql.Type

// RoleCreate create a role
var RoleCreate *graphql.Field

// RoleUpdate create a role
var RoleUpdate *graphql.Field

// RoleGetByID get role by id
var RoleGetByID *graphql.Field

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

	RoleGetByID = &graphql.Field{
		Type: Role,
		Args: graphql.FieldConfigArgument{"id": gNInt},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id := params.Args["id"].(int)
			role := &models.Role{ID: id}
			if err := role.GetByID(); err != nil {
				return nil, err
			}

			return role, nil
		},
	}

	RoleGetByName = &graphql.Field{
		Type: Role,
		Args: graphql.FieldConfigArgument{"name": gNString},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			name := params.Args["name"].(string)
			role := &models.Role{RoleName: name}
			if err := role.GetByName(); err != nil {
				return nil, err
			}

			return role, nil
		},
	}

	RoleCreate = &graphql.Field{
		Type: Role,
		Args: graphql.FieldConfigArgument{"name": gNString},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			role := &models.Role{
				RoleName: params.Args["name"].(string),
			}
			if err := role.Insert(); err != nil {
				return nil, err
			}
			return role, nil
		},
	}

	RoleUpdate = &graphql.Field{
		Type: Role,
		Args: graphql.FieldConfigArgument{
			"id":     gInt,
			"name":   gString,
			"status": gBaseStatus,
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			role := &models.Role{
				ID: params.Args["id"].(int),
			}
			attr := map[string]interface{}{
				"name":   params.Args["name"].(string),
				"status": params.Args["status"].(models.BaseStatus),
			}

			if err := role.Update(attr); err != nil {
				return nil, err
			}
			return role, nil
		},
	}
}
