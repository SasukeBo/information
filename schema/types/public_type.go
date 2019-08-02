package types

import (
	"fmt"
	"time"

	"github.com/SasukeBo/information/models"
	// "github.com/SasukeBo/information/resolvers"
	"github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

var gBaseStatus = &graphql.ArgumentConfig{Type: custom.BaseStatus, Description: `
		基础状态
		- default 默认状态
		- publish 发布状态
		- block   屏蔽（禁用）状态
		- deleted 删除状态
	`}

// Response 消息体
var Response = graphql.NewObject(graphql.ObjectConfig{
	Name:        "response",
	Description: "测试graphql",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"uuid":    &graphql.Field{Type: graphql.String},
		"name":    &graphql.Field{Type: graphql.String},
		"phone":   &graphql.Field{Type: graphql.String},
	},
})

type whoAmIResponse struct {
	Message string
	UUID    string
	Name    string
	Phone   string
}

// WhoAmIType 测试获取context中存储的current_user
var WhoAmIType = &graphql.Field{
	Type: Response,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		rootValue := p.Info.RootValue.(map[string]interface{})
		currentUserUUID := rootValue["currentUserUUID"]
		if currentUserUUID == nil {
			return whoAmIResponse{Message: "not authenticated!"}, nil
		}

		user := models.User{UUID: currentUserUUID.(string)}
		if err := models.Repo.Read(&user, "uuid"); err != nil {
			return nil, err
		}

		if err := models.Repo.Read(user.UserExtend); err != nil {
			return nil, err
		}
		return whoAmIResponse{UUID: user.UUID, Name: user.UserExtend.Name, Phone: user.Phone}, nil
	},
}

// SayHelloType 测试接口
var SayHelloType = &graphql.Field{
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
		return struct{ Message string }{Message: message}, nil
	},
}

// GenArg 简化gql参数定义
func GenArg(gqlType graphql.Input, des string, opts ...interface{}) *graphql.ArgumentConfig {
	defaultValue := interface{}(nil)
	if len(opts) > 0 && !opts[0].(bool) {
		gqlType = graphql.NewNonNull(gqlType)
	}

	if len(opts) > 1 {
		defaultValue = opts[1]
	}

	return &graphql.ArgumentConfig{
		Type:         gqlType,
		Description:  des,
		DefaultValue: defaultValue,
	}
}
