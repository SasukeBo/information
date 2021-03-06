package controllers

import (
	"encoding/json"
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"regexp"
	"strings"
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
var operationNameRegStr = `^((query|mutation)\s*(\w+)?\s*(\([\w\[\]\d\s$!:,]*\))?\s*|^){\s*(\w+:)?\s*(\w+)\s*(\((\n|.)*\))?\s*({(\n|.)*})?\s*}`

/*
  gqlGetSession 获取 session 中的数据到 gqlRootObject 中。
  conn - controller 对象
  obj - gqlRootObject graphql 根对象
  opName - string graphq 操作名
*/
func gqlGetSession(conn *beego.Controller, obj gqlRootObject, rootFieldName string) {
	switch rootFieldName {
	case
		"signUp",
		"resetPassword",
		"getSmsCode",
		"userUpdate":
		obj["phone"] = conn.GetSession("phone")
		obj["smsCode"] = conn.GetSession("smsCode")

	case "signIn":
		// 登录操作不需要后面的验证
		// 需要记录 IP UA
		obj["remote_ip"] = conn.Ctx.Input.IP()
		obj["user_agent"] = conn.Ctx.Input.UserAgent()
		fallthrough

	case "signOut", "getLastLogin", "getThisLogin":
		obj["session_id"] = conn.Ctx.Input.CruSession.SessionID()
	}

	obj["currentUser"] = conn.GetSession("currentUser")
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
		ctx.Input.SetData("gql_error", models.Error{Message: "query root field name missing"})
	}

	if params.OperationName == "IntrospectionQuery" {
		ctx.Input.SetData("need_auth", false)
	}

	switch params.RootFieldName {
	case
		"sendSmsCode",
		"signUp",
		"resetPassword",
		"getSmsCode",
		"signIn":
		ctx.Input.SetData("need_auth", false)
	}

	ctx.Input.SetData("query_params", params)
}

// HandleAdminGraphql validate user role isAdmin before HandleGraphql
func HandleAdminGraphql(ctx *context.Context) {
	if currentUser, ok := ctx.Input.Session("currentUser").(models.User); ok {
		var role *models.Role
		var err error

		if role, err = currentUser.LoadRole(); err != nil {
			ctx.Input.SetData("gql_error", err)
		}

		if !role.IsAdmin {
			ctx.Input.SetData("gql_error", models.Error{Message: "user is not admin."})
		}
	}

	HandleGraphql(ctx)
}

// 解析gql请求参数JSON到结构体中
func fetchParams(ctx *context.Context) queryParams {
	var params queryParams
	json.NewDecoder(ctx.Request.Body).Decode(&params)
	reg := regexp.MustCompile(operationNameRegStr)
	matches := reg.FindStringSubmatch(strings.TrimSpace(params.Query))

	if len(matches) > 2 {
		if matches[2] == "" {
			params.Type = "query"
		} else {
			params.Type = matches[2]
		}
	}

	if params.OperationName == "" && len(matches) > 3 {
		params.OperationName = matches[3]
	}

	if len(matches) > 6 {
		params.RootFieldName = matches[6]
	}

	env := beego.AppConfig.String
	if env("runmode") == "dev" && false {
		logs.Info(params.Query)
	}

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
