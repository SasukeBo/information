package types

import (
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

// RoleType 用户类型
var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.Int},
		"name":   &graphql.Field{Type: graphql.String, Description: "role name"},
		"status": &graphql.Field{Type: custom.BaseStatus, Description: "role status, can be default, publish, block and deleted"},
	},
})

// RoleCreate create a role
var RoleCreate = &graphql.Field{
	Type: RoleType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		role := &models.Role{
			Name: params.Args["name"].(string),
		}
		if err := role.Insert(); err != nil {
			return nil, err
		}
		return role, nil
	},
}

// RoleUpdate create a role
var RoleUpdate = &graphql.Field{
	Type: RoleType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"status": &graphql.ArgumentConfig{
			Type: custom.BaseStatus,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		role := &models.Role{
			Id: params.Args["id"].(int),
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

// RoleGetByID get role by id
var RoleGetByID = &graphql.Field{
	Type: RoleType,
	Args: graphql.FieldConfigArgument{"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)}},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Args["id"].(int)
		role := &models.Role{Id: id}
		if err := role.GetByID(); err != nil {
			return nil, err
		}

		return role, nil
	},
}
