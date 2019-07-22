package types

import (
	"github.com/graphql-go/graphql"
)

// RolePriv 用户类型
var RolePriv graphql.Type

func init() {
	RolePriv = graphql.NewObject(graphql.ObjectConfig{
		Name: "RolePriv",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"role":      &graphql.Field{Type: Role},
				"privilege": &graphql.Field{Type: Privilege},
			}
		}),
	})
}
