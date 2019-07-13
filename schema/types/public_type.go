package types

import (
	"fmt"
	"time"

	"github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

var (
	// 减少重复代码
	gBaseStatus = &graphql.ArgumentConfig{Type: custom.BaseStatus}
	gNString    = &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)}
	gString     = &graphql.ArgumentConfig{Type: graphql.String}
	gNInt       = &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)}
	gInt        = &graphql.ArgumentConfig{Type: graphql.Int}
)

type response struct {
	Message string
}

// Response 消息体
var Response = graphql.NewObject(graphql.ObjectConfig{
	Name:        "response",
	Description: "测试graphql",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
	},
})

// SayHello 测试接口
var SayHello = &graphql.Field{
	Type: Response,
	Args: graphql.FieldConfigArgument{
		"your_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		name := params.Args["your_name"]
		now := time.Now()
		message := fmt.Sprintf(
			"你好%s! 现在是：%d年%d月%d日 %d:%d:%d",
			name,
			now.Year(),
			int(now.Month()),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second(),
		)
		return &response{Message: message}, nil
	},
}
