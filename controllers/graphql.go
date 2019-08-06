package controllers

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/SasukeBo/information/schema"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

// GQLController is graphql controller
type GQLController struct {
	beego.Controller
}

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
var operationNameRegStr = `^query (\w+) {`

/*
// Options http method
func (conn *GQLController) Options() {
	// TODO: 关闭跨域请求
	conn.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:9080")
	conn.Ctx.Output.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
	conn.Ctx.Output.Header("Access-Control-Allow-Headers", "content-type")
	conn.Ctx.Output.SetStatus(204)
}
*/

// Get http method
func (conn *GQLController) Get() {
	queryString := conn.Ctx.Input.Query("query")
	variablesString := conn.Ctx.Input.Query("variables")
	operationName := conn.Ctx.Input.Query("operationName")
	graphiqlVersion := "0.11.11"
	conn.TplName = "graphql.html"
	conn.Data["GraphiqlVersion"] = graphiqlVersion
	conn.Data["QueryString"] = queryString
	conn.Data["VariablesString"] = variablesString
	conn.Data["OperationName"] = operationName
}

// Post http method
func (conn *GQLController) Post() {
	needAuth := true
	rootObject := gqlRootObject{}

	var result *graphql.Result
	var params struct {
		Query         string                 `json:"query"`
		Variables     map[string]interface{} `json:"variables"`
		OperationName string                 `json:"operationName"`
	}

	json.NewDecoder(conn.Ctx.Request.Body).Decode(&params)
	logs.Info(params.Query)

	if params.OperationName == "" {
		// 解决 query 中包含了 operationName 但是请求体 JSON 缺失 operationName 选项。
		reg := regexp.MustCompile(operationNameRegStr)
		if matches := reg.FindStringSubmatch(strings.TrimSpace(params.Query)); len(matches) > 1 {
			params.OperationName = matches[1]
		}
	}

	switch params.OperationName {
	case "IntrospectionQuery", "sendSmsCode":
		fallthrough
	case "register", "resetPassword", "getSmsCode":
		fallthrough
	case "loginByPassword":
		needAuth = false
	}

	if err := conn.GetSession("auth_error"); err != nil && needAuth {
		// 返回错误信息
		result = &graphql.Result{
			Errors: []gqlerrors.FormattedError{
				gqlerrors.FormattedError{
					Message: err.(error).Error(),
				},
			},
		}

	} else {
		gqlGetSession(conn, rootObject, params.OperationName)

		gqlParams := graphql.Params{
			Schema:         schema.Schema,
			RequestString:  params.Query,
			OperationName:  params.OperationName,
			VariableValues: params.Variables,
			RootObject:     rootObject,
		}

		result = graphql.Do(gqlParams)

		gqlSetSession(conn, gqlParams.RootObject)
	}
	// conn.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:9080")
	conn.Data["json"] = result
	conn.ServeJSON()
}
