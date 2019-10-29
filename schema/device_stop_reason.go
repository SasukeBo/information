package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/*							   types
------------------------------------------ */

var reasonType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ReasonType",
	Description: "停机类型",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.Int},
		"name": &graphql.Field{Type: graphql.String, Description: "原因类型名称"},
	},
})

var stopReasonType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "StopReason",
	Description: "停机原因",
	Fields: graphql.Fields{
		"id":      &graphql.Field{Type: graphql.Int},
		"content": &graphql.Field{Type: graphql.String, Description: "停机原因"},
		"code":    &graphql.Field{Type: graphql.String, Description: "停机代码"},
	},
})

func init() {
	reasonType.AddFieldConfig("reasons", &graphql.Field{
		Type:        graphql.NewList(stopReasonType),
		Description: "停机原因",
		Resolve:     resolver.ReasonTypeLoadReasons,
	})

	stopReasonType.AddFieldConfig("type", &graphql.Field{
		Type:        reasonType,
		Description: "停机类型",
		Resolve:     resolver.ReasonLoadType,
	})
}

/* 				   query
------------------------------ */

/* 					mutation
------------------------------ */

var reasonTypeCreate = &graphql.Field{
	Type: reasonType,
	Args: graphql.FieldConfigArgument{
		"name": GenArg(graphql.String, "类型名称", false),
	},
	Description: `创建停机类型`,
	Resolve:     resolver.CreateReasonType,
}

var reasonTypeDelete = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id": GenArg(graphql.Int, "类型ID", false),
	},
	Description: `删除停机类型`,
	Resolve:     resolver.DeleteReasonType,
}

var reasonTypeUpdate = &graphql.Field{
	Type: reasonType,
	Args: graphql.FieldConfigArgument{
		"id":   GenArg(graphql.Int, "类型ID", false),
		"name": GenArg(graphql.String, "类型名称", false),
	},
	Description: `更新停机类型`,
	Resolve:     resolver.UpdateReasonType,
}

var stopReasonCreate = &graphql.Field{
	Type: stopReasonType,
	Args: graphql.FieldConfigArgument{
		"content": GenArg(graphql.String, "停机原因", false),
		"code":    GenArg(graphql.String, "停机代码", false),
		"typeID":  GenArg(graphql.Int, "停机类型ID", false),
	},
	Description: `创建设备停机原因`,
	Resolve:     resolver.CreateStopReason,
}

var stopReasonUpdate = &graphql.Field{
	Type: stopReasonType,
	Args: graphql.FieldConfigArgument{
		"id":      GenArg(graphql.Int, "原因ID", false),
		"content": GenArg(graphql.String, "停机原因"),
		"code":    GenArg(graphql.String, "停机代码"),
	},
	Description: `更新设备停机原因`,
	Resolve:     resolver.UpdateStopReason,
}
