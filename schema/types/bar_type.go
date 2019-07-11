package types

import (
	"github.com/graphql-go/graphql"
)

type bar struct {
	Name string
}

var fieldBarType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Bar",
	Fields: graphql.Fields{
		"name": &graphql.Field{Type: graphql.String},
	},
})

// TestBar field bar
var TestBar = &graphql.Field{
	Type: fieldBarType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		type result struct {
			data interface{}
			err  error
		}
		ch := make(chan *result, 1)
		go func() {
			defer close(ch)
			b := &bar{Name: "Bar's name"}
			ch <- &result{data: b, err: nil}
		}()
		return func() (interface{}, error) {
			r := <-ch
			return r.data, r.err
		}, nil
	},
}
