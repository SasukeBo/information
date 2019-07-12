package types

import (
	"fmt"

	"github.com/SasukeBo/information/models"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// RoleType 用户类型
var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.Int},
		"name":   &graphql.Field{Type: graphql.String, Description: "role name"},
		"status": &graphql.Field{Type: roleStatus, Description: "role status, can be default, publish, block and deleted"},
	},
})

var roleStatus = graphql.NewScalar(graphql.ScalarConfig{
	Name: "roleStatus",
	Description: `roleStatus is represent role current status,
	it convert string to int for DB,
	and convert int to string for output`,
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case models.RoleStatus:
			fmt.Println("Serialize value:", value)
			return int(value)
		case *models.RoleStatus:
			fmt.Println("Serialize value:", value)
			return int(*value)
		}
		return nil
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			fmt.Println("ParseValue value:", value)
			return value
		case *string:
			fmt.Println("ParseValue value:", value)
			return value
		}
		return nil
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			fmt.Println("valueAST.Value:", valueAST.Value)
			return valueAST.Value
		}
		return nil
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
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		role := &models.Role{
			Id: params.Args["id"].(int),
		}
		attr := map[string]interface{}{
			"name":   params.Args["name"].(string),
			"status": params.Args["status"].(int),
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
