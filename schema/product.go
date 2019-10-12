package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/* 					 types
------------------------------ */

var productType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Product",
	Description: "产品类型",
	Fields: graphql.Fields{
		"id":               &graphql.Field{Type: graphql.Int},
		"name":             &graphql.Field{Type: graphql.String, Description: "产品名称"},
		"productor":        &graphql.Field{Type: graphql.String, Description: "生产负责人"},
		"productorContact": &graphql.Field{Type: graphql.String, Description: "生产负责人联系方式"},
		"customer":         &graphql.Field{Type: graphql.String, Description: "订货方"},
		"customerContact":  &graphql.Field{Type: graphql.String, Description: "订货方联系方式"},
		"total":            &graphql.Field{Type: graphql.Int, Description: "指标总量"},
		"orderNum":         &graphql.Field{Type: graphql.String, Description: "订单编号"},
		"createdAt":        &graphql.Field{Type: graphql.DateTime},
		"updatedAt":        &graphql.Field{Type: graphql.DateTime},
		"finishTime":       &graphql.Field{Type: graphql.DateTime},
	},
})

func init() {
	productType.AddFieldConfig("detectItems", &graphql.Field{
		Type:        graphql.NewList(detectItemType),
		Description: "产品检测项",
		Resolve:     resolver.LoadDetectItem,
	})

	productType.AddFieldConfig("instances", &graphql.Field{
		Type:        graphql.NewList(productInsType),
		Description: "产品实例",
		Resolve:     resolver.LoadProductIns,
	})

	productType.AddFieldConfig("devices", &graphql.Field{
		Type:        graphql.NewList(deviceType),
		Description: "生产设备",
		Resolve:     resolver.LoadDevice,
	})

	productType.AddFieldConfig("register", &graphql.Field{
		Type:        userType,
		Description: "注册人",
		Resolve:     resolver.LoadUser,
	})
}

/* 				  response
------------------------------ */

var productListResponse = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProductListResponse",
	Description: "产品列表",
	Fields: graphql.Fields{
		"count":    &graphql.Field{Type: graphql.Int},
		"products": &graphql.Field{Type: graphql.NewList(productType), Description: "产品列表"},
	},
})

/* 				   query
------------------------------ */

var productGet = &graphql.Field{
	Type:        productType,
	Args:        graphql.FieldConfigArgument{"id": GenArg(graphql.Int, "产品ID", false)},
	Description: `查询产品`,
	Resolve:     resolver.GetProduct,
}

var productList = &graphql.Field{
	Type: productListResponse,
	Args: graphql.FieldConfigArgument{
		"namePattern": GenArg(graphql.String, "产品名称模糊匹配"),
		"offset":      GenArg(graphql.Int, "列表偏移量"),
		"limit":       GenArg(graphql.Int, "列表最大值"),
	},
	Description: `获取产品列表`,
	Resolve:     resolver.ListProduct,
}

/* 				 	mutation
------------------------------ */

var productCreate = &graphql.Field{
	Type: productType,
	Args: graphql.FieldConfigArgument{
		"name":             GenArg(graphql.String, "产品名称", false),
		"detectItems":      GenArg(graphql.NewList(detectItemInputType), "产品检测值", false),
		"productor":        GenArg(graphql.String, "生产负责人"),
		"productorContact": GenArg(graphql.String, "生产负责人联系方式"),
		"finishTime":       GenArg(graphql.DateTime, "预计完成订单时间"),
		"total":            GenArg(graphql.Int, "计划生产产品总量"),
		"orderNum":         GenArg(graphql.String, "订单编号"),
		"customer":         GenArg(graphql.String, "订货方"),
		"customerContact":  GenArg(graphql.String, "订货方联系方式"),
	},
	Description: `注册产品信息`,
	Resolve:     resolver.CreateProduct,
}

var productDelete = &graphql.Field{
	Type:        graphql.Int,
	Args:        graphql.FieldConfigArgument{"id": GenArg(graphql.Int, "产品ID", false)},
	Description: `删除产品`,
	Resolve:     resolver.DeleteProduct,
}

var productUpdate = &graphql.Field{
	Type: productType,
	Args: graphql.FieldConfigArgument{
		"id":               GenArg(graphql.Int, "产品ID", false),
		"name":             GenArg(graphql.String, "产品名称"),
		"productor":        GenArg(graphql.String, "生产负责人"),
		"productorContact": GenArg(graphql.String, "生产负责人联系方式"),
		"finishTime":       GenArg(graphql.DateTime, "预计完成订单时间"),
		"total":            GenArg(graphql.Int, "计划生产产品总量"),
		"orderNum":         GenArg(graphql.String, "订单编号"),
		"customer":         GenArg(graphql.String, "订货方"),
		"customerContact":  GenArg(graphql.String, "订货方联系方式"),
	},
	Description: `修改产品`,
	Resolve:     resolver.UpdateProduct,
}
