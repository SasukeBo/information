package types

import (
	// "github.com/SasukeBo/information/models"
	// "github.com/SasukeBo/information/schema/custom"
	// "github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// UserProfileType 用户类型
var UserProfileType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfile",
	Fields: graphql.Fields{
		"uuid":     &graphql.Field{Type: graphql.String},
		"realName": &graphql.Field{Type: graphql.String},
		"phone":    &graphql.Field{Type: graphql.String},
		"email":    &graphql.Field{Type: graphql.String},
	},
})
