package schema

import (
	"github.com/SasukeBo/information/schema/types"
	"github.com/graphql-go/graphql"
	"log"
)

// QueryRoot is query root
var QueryRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"sayHello":      types.SayHello,
		"whoAmI":        types.WhoAmI,
		"roleGet":       types.RoleGet,
		"roleGetByName": types.RoleGetByName,
	},
})

// MutateRoot is mutation root
var MutateRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"register":        types.UserCreate,
		"loginByPassword": types.LoginByPassword,
		"roleCreate":      types.RoleCreate,
		"roleUpdate":      types.RoleUpdate,
		"sendSmsCode":     types.SendSmsCode,
	},
})

// Schema is graphql schema
var Schema graphql.Schema

func init() {
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryRoot,
		Mutation: MutateRoot,
	})
	if err != nil {
		log.Fatal("failed to create new schema, err: ", err)
	}
}
