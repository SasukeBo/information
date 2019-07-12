package schema

import (
	"github.com/SasukeBo/information/schema/types"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
)

// GraphqlHander is graphql http hander
var GraphqlHander *handler.Handler

// QueryRoot is query root
var QueryRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"sayHello":    types.SayHello,
		"roleGetByID": types.RoleGetByID,
	},
})

// MutateRoot is mutation root
var MutateRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"userCreate": types.UserCreate,
		"roleCreate": types.RoleCreate,
		"roleUpdate": types.RoleUpdate,
	},
})

func init() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryRoot,
		Mutation: MutateRoot,
	})
	if err != nil {
		log.Fatal("failed to create new schema, err:", err)
	}

	GraphqlHander = handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: false,
	})
}
