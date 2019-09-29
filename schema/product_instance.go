package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/* 					 types
------------------------------ */

var productInsType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProductIns",
	Description: "产品实例",
	Fields: graphql.Fields{
		"id":               &graphql.Field{Type: graphql.Int},
		"detectItemValues": &graphql.Field{Type: graphql.NewList(detectItemValueType), Description: "各检测项对应的值"},
		"createdAt":        &graphql.Field{Type: graphql.DateTime},
	},
})

func init() {
	productInsType.AddFieldConfig("product", &graphql.Field{
		Type:        productType,
		Description: "产品信息",
		Resolve:     resolver.LoadProduct,
	})
}
