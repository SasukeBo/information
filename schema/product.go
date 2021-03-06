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
		"token":            &graphql.Field{Type: graphql.String, Description: "产品token"},
		"productor":        &graphql.Field{Type: graphql.String, Description: "生产负责人"},
		"productorContact": &graphql.Field{Type: graphql.String, Description: "生产负责人联系方式"},
		"customer":         &graphql.Field{Type: graphql.String, Description: "订货方"},
		"customerContact":  &graphql.Field{Type: graphql.String, Description: "订货方联系方式"},
		"total":            &graphql.Field{Type: graphql.Int, Description: "指标总量"},
		"orderNum":         &graphql.Field{Type: graphql.String, Description: "订单编号"},
		"createdAt":        &graphql.Field{Type: graphql.DateTime},
		"updatedAt":        &graphql.Field{Type: graphql.DateTime},
		"finishTime":       &graphql.Field{Type: graphql.DateTime},
		"currentCount":     &graphql.Field{Type: graphql.Int, Description: "当前产量", Resolve: resolver.CurrentProductInsCount},
		"detectItemsCount": &graphql.Field{Type: graphql.Int, Description: "检测项数", Resolve: resolver.ProductDetectItemsCount},
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

var productHistogramObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProductHistogramResponse",
	Description: "直方图数据对象",
	Fields: graphql.Fields{
		"xAxisData":  &graphql.Field{Type: graphql.NewList(graphql.String), Description: "x轴数据"},
		"seriesData": &graphql.Field{Type: graphql.NewList(graphql.Int), Description: "区间数据频度"},
	},
})

var productOverviewResponse = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProductOverviewResponse",
	Description: "产品数据总览",
	Fields: graphql.Fields{
		"deviceTotalCount":    &graphql.Field{Type: graphql.Int, Description: "产品生产设备总数"},
		"deviceProdCount":     &graphql.Field{Type: graphql.Int, Description: "产品生产设备运行中总数"},
		"instanceCount":       &graphql.Field{Type: graphql.Int, Description: "产量"},
		"qualifiedCount":      &graphql.Field{Type: graphql.Int, Description: "良品数量"},
		"todayInstanceCount":  &graphql.Field{Type: graphql.Int, Description: "今日产量"},
		"todayQualifiedCount": &graphql.Field{Type: graphql.Int, Description: "今日良品数量"},
	},
})

/* 				   query
------------------------------ */

var productOverview = &graphql.Field{
	Type: productOverviewResponse,
	Args: graphql.FieldConfigArgument{
		"id": GenArg(graphql.Int, "产品ID", false),
	},
	Description: "产品数据总览",
	Resolve:     resolver.ProductOverView,
}

var productDevicesGet = &graphql.Field{
	Type: graphql.NewList(deviceType),
	Args: graphql.FieldConfigArgument{
		"id": GenArg(graphql.Int, "产品ID", false),
	},
	Description: "获取产品的生产设备列表",
	Resolve:     resolver.GetProductDevices,
}
var productHistogram = &graphql.Field{
	Type: productHistogramObject,
	Args: graphql.FieldConfigArgument{
		"id":           GenArg(graphql.Int, "产品ID", false),
		"detectItemID": GenArg(graphql.Int, "产品检测项ID", false),
		"deviceID":     GenArg(graphql.Int, "生产设备ID"),
		"lowerTime":    GenArg(graphql.DateTime, "时间区间下限"),
		"upperTime":    GenArg(graphql.DateTime, "时间区间上限"),
	},
	Description: `获取产品某检测项检测数据直方图`,
	Resolve:     resolver.ProductHistogram,
}

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
		"self":        GenArg(graphql.Boolean, "只看自己", true, false),
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
