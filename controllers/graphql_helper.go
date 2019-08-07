package controllers

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"

	"github.com/SasukeBo/information/utils"
)

// graphiqlData is the page data structure of the rendered GraphiQL page
type graphiqlData struct {
	GraphiqlVersion string
	QueryString     string
	VariablesString string
	OperationName   string
	ResultString    string
}

type gqlRootObject map[string]interface{}

// 匹配 graphql query oprtationName
var operationNameRegStr = `^((query)|(mutation))\s*(\w+)?\s*(\([\w\d\s$!:,]*\))?\s*{\s*(\w+:)?\s*(\w+)\s*(\((\n|.)*\))?\s*({[\s\w\d]*})?\s*}`

/*
	gqlGetSession 获取 session 中的数据到 gqlRootObject 中。
	conn - controller 对象
	obj - gqlRootObject graphql 根对象
	opName - string graphq 操作名
*/
func gqlGetSession(conn *beego.Controller, obj gqlRootObject, rootFieldName string) {
	switch rootFieldName {
	case "IntrospectionQuery", "sendSmsCode":
		// graphiql schema query
		break

	case "register", "resetPassword", "getSmsCode", "userUpdatePhone":
		obj["phone"] = conn.GetSession("phone")
		obj["smsCode"] = conn.GetSession("smsCode")

	case "loginByPassword":
		// 登录操作不需要后面的验证
		// 需要记录 IP UA
		obj["remote_ip"] = conn.Ctx.Input.IP()
		obj["user_agent"] = conn.Ctx.Input.UserAgent()
		fallthrough

	case "logout":
		obj["session_id"] = conn.Ctx.Input.CruSession.SessionID()
	}

	obj["currentUserUUID"] = conn.GetSession("currentUserUUID")
}

// gqlSetSession 根据 graphql.Params.RootObject 中 set session 对应的
// string 数组内容，从 RootObject 中取值并放入 Session
func gqlSetSession(conn *beego.Controller, obj gqlRootObject) {
	if fields := obj["setSession"]; fields != nil {
		for _, field := range fields.([]string) {
			conn.SetSession(field, obj[field])
		}
	}
	if remember := obj["remember"]; remember != nil {
		if !remember.(bool) {
			// 如果用户登录为不记住登录，则将cookie过期时间设置为 回话
			currentSessionID := conn.Ctx.Input.CruSession.SessionID()
			conn.Ctx.Output.Cookie(beego.AppConfig.String("SessionName"), currentSessionID, 0)
		}
	}
}

// 处理 graphiql get 方法的 Logic
func graphiqlGet(conn *beego.Controller) {
	conn.Data["GraphiqlVersion"] = "0.11.11"
	conn.Data["QueryString"] = conn.Ctx.Input.Query("query")
	conn.Data["VariablesString"] = conn.Ctx.Input.Query("variables")
	conn.Data["OperationName"] = conn.Ctx.Input.Query("operationName")
	conn.TplName = "graphql.html"
}

// gql请求参数结构体
type queryParams struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
	OperationName string                 `json:"operationName"`
	RootFieldName string
	Type          string
}

// HandleGraphql graphql request fetch queryParams
func HandleGraphql(ctx *context.Context) {
	params := fetchParams(ctx)
	ctx.Input.SetData("need_auth", true)

	if params.RootFieldName == "" && params.OperationName != "IntrospectionQuery" {
		ctx.Input.SetData("gql_error", utils.LogicError{
			Message: "query root field name missing",
		})
	}

	if params.OperationName == "IntrospectionQuery" {
		ctx.Input.SetData("need_auth", false)
	}

	switch params.RootFieldName {
	case
		"sendSmsCode",
		"register",
		"resetPassword",
		"getSmsCode",
		"loginByPassword":
		ctx.Input.SetData("need_auth", false)
	}

	ctx.Input.SetData("query_params", params)
}

// 解析gql请求参数JSON到结构体中
func fetchParams(ctx *context.Context) queryParams {
	var params queryParams
	json.NewDecoder(ctx.Request.Body).Decode(&params)
	reg := regexp.MustCompile(operationNameRegStr)
	matches := reg.FindStringSubmatch(strings.TrimSpace(params.Query))

	if len(matches) > 1 {
		params.Type = matches[1]
	}

	if params.OperationName == "" && len(matches) > 4 {
		params.OperationName = matches[4]
	}

	if len(matches) > 7 {
		params.RootFieldName = matches[7]
	}

	logs.Info(params.Query)

	return params
}

// 封装GQL error result
func genGQLError(e error) *graphql.Result {
	// 返回错误信息
	return &graphql.Result{
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormattedError{
				Message: e.Error(),
			},
		},
	}
}
