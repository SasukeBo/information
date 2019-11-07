package resolver

import (
	"encoding/json"
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
	"github.com/astaxie/beego"
	"github.com/graphql-go/graphql"
	"regexp"
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

	// validate phone
	phone := p.Args["phone"].(string)
	pattern := `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[35678]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|66\d{2})\d{6}$`
	reg := regexp.MustCompile(pattern)
	if !reg.Match([]byte(phone)) {
		return nil, models.Error{Message: "invalid phone number."}
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
		return nil, models.Error{Message: "this api only work on dev environment."}
	}
	rootValue := p.Info.RootValue.(map[string]interface{})
	smsCode := rootValue["smsCode"]
	if smsCode == nil {
		return nil, models.Error{Message: "smsCode not found."}
	}

	return smsCode.(string), nil
}
