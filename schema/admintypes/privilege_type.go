package admintypes

import (
	"github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

// Privilege 用户类型
var Privilege graphql.Type

func init() {
	Privilege = graphql.NewObject(graphql.ObjectConfig{
		Name: "Privilege",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":        &graphql.Field{Type: graphql.Int},
				"privName":  &graphql.Field{Type: graphql.String},
				"privType":  &graphql.Field{Type: graphql.Int},
				"status":    &graphql.Field{Type: custom.BaseStatus, Description: "基础状态"},
				"rolePrivs": &graphql.Field{Type: graphql.NewList(RolePriv), Description: "role and privilege relationship"},
				"createdAt": &graphql.Field{Type: graphql.DateTime},
				"updatedAt": &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})
}
