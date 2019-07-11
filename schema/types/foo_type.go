package types

import (
	"github.com/graphql-go/graphql"
)

type foo struct {
	Name string
}

var fieldFooType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Foo",
	Fields: graphql.Fields{
		"name": &graphql.Field{Type: graphql.String},
	},
})

// TestFoo fields foo
var TestFoo = &graphql.Field{
	Type: fieldFooType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// var f = foo{Name: "Foo's name"}
		// return func() (interface{}, error) {
		// return &f, nil
		// }, nil
		return &foo{Name: "SasukeBo"}, nil
	},
}
