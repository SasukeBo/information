package types

import (
	// "github.com/SasukeBo/information/models"
	"github.com/graphql-go/graphql"
)

// UserType 用户类型
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"uuid":         &graphql.Field{Type: graphql.String},
		"account":      &graphql.Field{Type: graphql.String},
		"password":     &graphql.Field{Type: graphql.String},
		"user_profile": &graphql.Field{Type: graphql.String},
		"role":         &graphql.Field{Type: RoleType},
		"status":       &graphql.Field{Type: graphql.Int},
		"created_at":   &graphql.Field{Type: graphql.DateTime},
		"upadted_at":   &graphql.Field{Type: graphql.DateTime},
	},
})

// UserCreate create a user
var UserCreate = &graphql.Field{
	Type: UserType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}
