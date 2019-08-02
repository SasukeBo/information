package types

import (
	// "github.com/SasukeBo/information/models"
	// "github.com/SasukeBo/information/schema/custom"
	// "github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// UserExtend 用户类型
var UserExtend graphql.Type

func init() {
	UserExtend = graphql.NewObject(graphql.ObjectConfig{
		Name: "UserExtend",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":    &graphql.Field{Type: graphql.Int},
				"user":  &graphql.Field{Type: User},
				"name":  &graphql.Field{Type: graphql.String},
				"email": &graphql.Field{Type: graphql.String},
			}
		}),
	})
}
