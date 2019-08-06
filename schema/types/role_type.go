package types

import (
	"github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

// Role 用户类型
var Role graphql.Type

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
}
