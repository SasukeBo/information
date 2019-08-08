package aliyun

import (
  "encoding/json"

  "github.com/SasukeBo/information/schema/resolvers"
  "github.com/SasukeBo/information/utils"
  "github.com/astaxie/beego"
  "github.com/graphql-go/graphql"
)

type sendSmsCodeResponse struct {
  Message   string `json:"Message"`
  RequestID string `json:"RequestID"`
  Code      string `json:"Code"`
  BizID     string `json:"BizID"`
}

// SendSmsCode is a gql resolver, send message code by aliyun service
func SendSmsCode(p graphql.ResolveParams) (interface{}, error) {
  rootValue := p.Info.RootValue.(map[string]interface{})
  var response sendSmsCodeResponse

  phone := p.Args["phone"].(string)
  // validate phone
  if err := resolvers.ValidatePhone(phone); err != nil {
    return nil, err
  }
  smsCode := utils.GenSmsCode()

  if disableSend, _ := beego.AppConfig.Bool("DisableSend"); disableSend {
    // 如果设置了阻止发送短信
    response.Code = "OK"
    response.Message = "测试短信验证码，不使用aliyun发送"
  } else {
    // 否则正常调用aliyun短信服务
    smsRsp, err := utils.SendSmsCode(phone, smsCode)
    if err != nil {
      return nil, err
    }

    err = json.Unmarshal([]byte(smsRsp.GetHttpContentString()), &response)
    if err != nil {
      return nil, err
    }
  }

  rootValue["phone"] = phone
  rootValue["smsCode"] = smsCode
  rootValue["setSession"] = []string{"phone", "smsCode"}

  return response, nil
}

// GetSmsCode is a gql resolver, get develop test smsCode
// 仅在测试环境下有效
func GetSmsCode(p graphql.ResolveParams) (interface{}, error) {
  if beego.AppConfig.String("runmode") != "dev" {
    return nil, utils.LogicError{
      Message: "this api only work on dev environment.",
    }
  }
  rootValue := p.Info.RootValue.(map[string]interface{})
  smsCode := rootValue["smsCode"]
  if smsCode == nil {
    return nil, utils.LogicError{
      Message: "smsCode not find.",
    }
  }

  return smsCode.(string), nil
}
