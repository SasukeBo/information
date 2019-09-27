package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/*							  response
------------------------------------------ */

var sendSmsCodeResponse = graphql.NewObject(graphql.ObjectConfig{
	Name:        "SendSmsCodeResponse",
	Description: "短信验证码response消息体",
	Fields: graphql.Fields{
		"message":   &graphql.Field{Type: graphql.String, Description: "状态码的描述"},
		"requestID": &graphql.Field{Type: graphql.String, Description: "请求ID"},
		"code":      &graphql.Field{Type: graphql.String, Description: "请求状态码，`OK`代表请求成功，其他错误码见阿里云短信服务错误码列表"},
		"bizID":     &graphql.Field{Type: graphql.String, Description: "状态码的描述"},
	},
})

/*					mutation
-------------------------------- */

var sendSmsCode = &graphql.Field{
	Type: sendSmsCodeResponse,
	Args: graphql.FieldConfigArgument{
		"phone": GenArg(graphql.String, "手机号", false),
	},
	Resolve: resolver.SendSmsCode,
	Description: `
#### 发送短信验证码

**注意** 需要提供 operationName
`,
}

/*						query
-------------------------------- */

var getSmsCode = &graphql.Field{
	Type: graphql.String,
	Description: `#### 获取session对话中存储的短信验证码。
当app.conf中DisableSend设置为true时，发送短信功能将不会真的发送短信，可以调用此接口获取短信验证码，该接口仅用于测试环境。

**注意** 请求时需要加上 operationName`,
	Resolve: resolver.GetSmsCode,
}
