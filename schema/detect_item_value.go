package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/* 					 types
------------------------------ */

var detectItemValueType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "DetectItemValue",
	Description: "检测值",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.Int},
		"value": &graphql.Field{Type: graphql.Float, Description: "检测值"},
	},
})

func init() {
	detectItemValueType.AddFieldConfig("detectItem", &graphql.Field{
		Type:        detectItemType,
		Description: "检测项",
		Resolve:     resolver.LoadDetectItem,
	})
}
