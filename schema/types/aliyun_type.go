package types

import (
	"encoding/json"
	"github.com/SasukeBo/information/utils"
	"github.com/graphql-go/graphql"
)

// SendSmsCode 发送短信验证码调用接口
var SendSmsCode *graphql.Field

// SendSmsCodeResponse response type of SendSmsCode
var SendSmsCodeResponse graphql.Type

type sendSmsCodeResponse struct {
	Message   string `json:"Message"`
	RequestID string `json:"RequestID"`
	Code      string `json:"Code"`
	BizID     string `json:"BizID"`
}

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
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			rootValue := p.Info.RootValue.(map[string]interface{})
			var response sendSmsCodeResponse

			phone := p.Args["phone"].(string)
			smsCode := utils.GenSmsCode()
			smsRsp, err := utils.SendSmsCode(phone, smsCode)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal([]byte(smsRsp.GetHttpContentString()), &response)
			if err != nil {
				return nil, err
			}

			rootValue["phone"] = phone
			rootValue["smsCode"] = smsCode
			rootValue["setSession"] = []string{"phone", "smsCode"}

			return response, nil
		},
	}
}
