package types

import (
	"github.com/SasukeBo/information/resolvers/aliyun"
	"github.com/graphql-go/graphql"
)

// SendSmsCode 发送短信验证码调用接口
var SendSmsCode *graphql.Field

// GetSmsCode 获取测试环境下未真正发送验证码短信时的短信验证码
var GetSmsCode *graphql.Field

// SendSmsCodeResponse response type of SendSmsCode
var SendSmsCodeResponse graphql.Type

func init() {
	// SendSmsCodeResponse 短信验证码response消息体
	SendSmsCodeResponse = graphql.NewObject(graphql.ObjectConfig{
		Name:        "SendSmsCodeResponse",
		Description: "短信验证码response消息体",
		Fields: graphql.Fields{
			"message":   &graphql.Field{Type: graphql.String, Description: "状态码的描述"},
			"requestID": &graphql.Field{Type: graphql.String, Description: "请求ID"},
			"code": &graphql.Field{Type: graphql.String, Description: `
			请求状态码。
			· 返回OK代表请求成功
			· 其他错误码见阿里云短信服务错误码列表
			`},
			"bizID": &graphql.Field{Type: graphql.String, Description: "状态码的描述"},
		},
	})

	SendSmsCode = &graphql.Field{
		Type: SendSmsCodeResponse,
		Args: graphql.FieldConfigArgument{
			"phone": GenArg(graphql.String, "手机号", false),
		},
		Resolve: aliyun.SendSmsCode,
	}

	GetSmsCode = &graphql.Field{
		Type:        graphql.String,
		Description: `当配置中 DisableSend 设置为true时，发送短信功能将不会真的发送短信，可以调用该接口获取验证码，该功能只适用于测试环境。`,
		Resolve:     aliyun.GetSmsCode,
	}
}
