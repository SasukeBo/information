package types

import (
	"github.com/SasukeBo/information/models"
	"github.com/graphql-go/graphql"
)

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

type whoIAmResponse struct {
	Message string
	UUID    string
	Name    string
	Phone   string
}

// WhoIAmType 测试获取context中存储的current_user
var WhoIAmType = &graphql.Field{
	Type: Response,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		rootValue := p.Info.RootValue.(map[string]interface{})
		currentUserUUID := rootValue["currentUserUUID"]
		if currentUserUUID == nil {
			return whoIAmResponse{Message: "not authenticated!"}, nil
		}

		user := models.User{UUID: currentUserUUID.(string)}
		if err := models.Repo.Read(&user, "uuid"); err != nil {
			return nil, err
		}

		if err := models.Repo.Read(user.UserExtend); err != nil {
			return nil, err
		}
		return whoIAmResponse{UUID: user.UUID, Name: user.UserExtend.Name, Phone: user.Phone}, nil
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
