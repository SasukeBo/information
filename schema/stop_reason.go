package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/*							   types
------------------------------------------ */

var stopReasonType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "StopReason",
	Description: "停机原因",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"wordIndex": &graphql.Field{Type: graphql.Int, Description: "第几个字节"},
		"bitPos":    &graphql.Field{Type: graphql.Int, Description: "第几个bit位"},
		"content":   &graphql.Field{Type: graphql.String, Description: "停机原因"},
	},
})

/* 				   query
------------------------------ */

var logStopReasonsGet = &graphql.Field{
	Type: graphql.NewList(stopReasonType),
	Args: graphql.FieldConfigArgument{
		"logID": GenArg(graphql.Int, "停机日志ID", false),
	},
	Description: `获取停机日志关联的停机原因`,
	Resolve:     resolver.GetLogStopReasons,
}

/* 					mutation
------------------------------ */

var stopReasonCreate = &graphql.Field{
	Type: stopReasonType,
	Args: graphql.FieldConfigArgument{
		"content":   GenArg(graphql.String, "停机原因", false),
		"wordIndex": GenArg(graphql.Int, "第几个字", false),
		"bitPos":    GenArg(graphql.Int, "第几个二进制位", false),
		"deviceID":  GenArg(graphql.Int, "设备ID", false),
	},
	Description: `创建设备停机原因`,
	Resolve:     resolver.CreateStopReason,
}

var stopReasonUpdate = &graphql.Field{
	Type: stopReasonType,
	Args: graphql.FieldConfigArgument{
		"id":        GenArg(graphql.Int, "原因ID", false),
		"content":   GenArg(graphql.String, "停机原因"),
		"wordIndex": GenArg(graphql.Int, "第几个字"),
		"bitPos":    GenArg(graphql.Int, "第几个二进制位"),
	},
	Description: `更新设备停机原因`,
	Resolve:     resolver.UpdateStopReason,
}

var stopReasonDelete = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id": GenArg(graphql.Int, "原因ID", false),
	},
	Description: `删除设备停机原因`,
	Resolve:     resolver.DeleteStopReason,
}
