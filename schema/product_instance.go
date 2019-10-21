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
		"id": &graphql.Field{Type: graphql.Int},
		"detectItemValues": &graphql.Field{
			Type:        graphql.NewList(detectItemValueType),
			Description: "各检测项对应的值",
			Resolve:     resolver.ProductInsLoadDetectItemValues,
		},
		"createdAt": &graphql.Field{Type: graphql.DateTime},
	},
})

func init() {
	productInsType.AddFieldConfig("product", &graphql.Field{
		Type:        productType,
		Description: "产品信息",
		Resolve:     resolver.LoadProduct,
	})
}

/* 				  response
------------------------------ */

var productInsListResponse = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProductInsListResponse",
	Description: "产品实例列表",
	Fields: graphql.Fields{
		"count":      &graphql.Field{Type: graphql.Int},
		"productIns": &graphql.Field{Type: graphql.NewList(productInsType), Description: "产品实例列表"},
	},
})

/* 				   query
------------------------------ */
