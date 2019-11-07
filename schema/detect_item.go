package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/* 					 types
------------------------------ */

var detectItemType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "DetectItem",
	Description: "检测项",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"sign":       &graphql.Field{Type: graphql.String, Description: "检测项标识"},
		"upperLimit": &graphql.Field{Type: graphql.Float, Description: "值上限"},
		"lowerLimit": &graphql.Field{Type: graphql.Float, Description: "值下限"},
		"createdAt":  &graphql.Field{Type: graphql.DateTime},
		"updatedAt":  &graphql.Field{Type: graphql.DateTime},
	},
})

var detectItemInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "DetectItemInput",
	Description: "检测项参数对象",
	Fields: graphql.InputObjectConfigFieldMap{
		"sign":       &graphql.InputObjectFieldConfig{Type: graphql.String, Description: "检测项标识"},
		"upperLimit": &graphql.InputObjectFieldConfig{Type: graphql.Float, Description: "值上限"},
		"lowerLimit": &graphql.InputObjectFieldConfig{Type: graphql.Float, Description: "值下限"},
	},
})

func init() {
	detectItemType.AddFieldConfig("product", &graphql.Field{
		Type:        productType,
		Description: "检测项所述的产品分类",
		Resolve:     resolver.LoadProduct,
	})
}

/* 					response
------------------------------ */

var detectItemListResponse = graphql.NewObject(graphql.ObjectConfig{
	Name:        "DetectItemListResponse",
	Description: "检测项列表",
	Fields: graphql.Fields{
		"count":       &graphql.Field{Type: graphql.Int, Description: "总数"},
		"detectItems": &graphql.Field{Type: graphql.NewList(detectItemType), Description: "检测项列表"},
	},
})

/* 					 query
------------------------------ */

var detectItemGet = &graphql.Field{
	Type:        detectItemType,
	Args:        graphql.FieldConfigArgument{"id": GenArg(graphql.Int, "检测项ID", false)},
	Description: `查询产品检测项`,
	Resolve:     resolver.GetDetectItem,
}

var detectItemList = &graphql.Field{
	Type: detectItemListResponse,
	Args: graphql.FieldConfigArgument{
		"productID": GenArg(graphql.Int, "产品ID", false),
		"limit":     GenArg(graphql.Int, "最大返回条数"),
		"offset":    GenArg(graphql.Int, "查询偏移量"),
	},
	Description: `查询产品检测项`,
	Resolve:     resolver.ListDetectItem,
}

/* 					mutation
------------------------------ */

var detectItemCreate = &graphql.Field{
	Type: detectItemType,
	Args: graphql.FieldConfigArgument{
		"productID":  GenArg(graphql.Int, "产品ID", false),
		"sign":       GenArg(graphql.String, "检测项标识", false),
		"upperLimit": GenArg(graphql.Float, "上限值"),
		"lowerLimit": GenArg(graphql.Float, "下限值"),
	},
	Description: `增加产品检测项`,
	Resolve:     resolver.CreateDetectItem,
}

var detectItemUpdate = &graphql.Field{
	Type: detectItemType,
	Args: graphql.FieldConfigArgument{
		"id":         GenArg(graphql.Int, "检测项ID", false),
		"sign":       GenArg(graphql.String, "检测项标识"),
		"upperLimit": GenArg(graphql.Float, "上限值"),
		"lowerLimit": GenArg(graphql.Float, "下限值"),
	},
	Description: `更新产品检测项`,
	Resolve:     resolver.UpdateDetectItem,
}

var detectItemDelete = &graphql.Field{
	Type:        graphql.Int,
	Args:        graphql.FieldConfigArgument{"id": GenArg(graphql.Int, "检测项ID", false)},
	Description: `删除产品检测项`,
	Resolve:     resolver.DeleteDetectItem,
}
