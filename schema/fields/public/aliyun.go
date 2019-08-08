package public

import (
  "github.com/graphql-go/graphql"

  "github.com/SasukeBo/information/schema/fields"
  "github.com/SasukeBo/information/schema/resolvers/aliyun"
  "github.com/SasukeBo/information/schema/types"
)

// SendSmsCodeField 发送短信验证码调用接口
var SendSmsCodeField = &graphql.Field{
  Type: types.SendSmsCodeResponse,
  Args: graphql.FieldConfigArgument{
    "phone": fields.GenArg(graphql.String, "手机号", false),
  },
  Resolve:     aliyun.SendSmsCode,
  Description: "请求时需要加上 operationName",
}

// GetSmsCodeField 获取测试环境下未真正发送验证码短信时的短信验证码
var GetSmsCodeField = &graphql.Field{
  Type:        graphql.String,
  Description: `请求时需要加上 operationName，当配置中 DisableSend 设置为true时，发送短信功能将不会真的发送短信，可以调用该接口获取验证码，该功能只适用于测试环境。`,
  Resolve:     aliyun.GetSmsCode,
}
