package types

import (
	"github.com/SasukeBo/information/utils"
	"github.com/graphql-go/graphql"
)

// SendSmsCode 发送短信验证码调用接口
var SendSmsCode *graphql.Field

// SendSmsCodeResponse response type of SendSmsCode
var SendSmsCodeResponse graphql.Type

func init() {
	SendSmsCode = &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"phone": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), Description: "手机号"},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			phone := p.Args["phone"].(string)
			response, err := utils.SendSmsCode(phone, "12345")
			return response.GetHttpContentString(), err
		},
	}
}
