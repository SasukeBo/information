package utils

import (
  "encoding/json"
  "fmt"

  "github.com/aliyun/alibaba-cloud-sdk-go/sdk"
  "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
  "github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/logs"
)

var client *sdk.Client
var env = beego.AppConfig.String

// SendSmsCode 发送短信验证码
// phone 手机号
// code 验证码
func SendSmsCode(phone, code string) (*responses.CommonResponse, error) {
  request := requests.NewCommonRequest()
  request.Method = "POST"
  request.Scheme = "https"
  request.Domain = "dysmsapi.aliyuncs.com"
  request.Version = "2017-05-25"
  request.ApiName = "SendSms"
  request.QueryParams["RegionId"] = "cn-hangzhou"
  request.QueryParams["PhoneNumbers"] = phone
  request.QueryParams["SignName"] = env("SignName")
  request.QueryParams["TemplateCode"] = env("templatecode")

  templateParam := struct {
    Code string `json:"code"`
  }{Code: code}

  tpJSON, _ := json.Marshal(templateParam)
  request.QueryParams["TemplateParam"] = fmt.Sprintf("%s", tpJSON)

  response, err := client.ProcessCommonRequest(request)

  return response, err
}

func init() {
  var err error
  client, err = sdk.NewClientWithAccessKey(
    "cn-hangzhou",
    env("accesskeyid"),
    env("accesskeysecret"),
  )
  if err != nil {
    logs.Error("package utils init func: ", err)
  }
}
