package controllers

import (
	"encoding/json"

	"github.com/SasukeBo/information/schema"
	"github.com/astaxie/beego"
	"github.com/graphql-go/graphql"
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
	var params struct {
		Query         string                 `json:"query"`
		Variables     map[string]interface{} `json:"variables"`
		OperationName string                 `json:"operationName"`
	}

	json.NewDecoder(conn.Ctx.Request.Body).Decode(&params)

	rootObject := gqlRootObject{}
	setObject(conn, rootObject, params.OperationName)

	gqlParams := graphql.Params{
		Schema:         schema.Schema,
		RequestString:  params.Query,
		OperationName:  params.OperationName,
		VariableValues: params.Variables,
		RootObject:     rootObject,
	}

	result := graphql.Do(gqlParams)
	setSession(conn, gqlParams.RootObject)

	conn.Data["json"] = result
	conn.ServeJSON()
}

// setObject 根据 gql 操作名称向 graphql.Params.RootObject 中放入值
// 至少会放入 currentUser 信息
func setObject(conn *GQLController, obj gqlRootObject, name string) {
	switch name {
	case "register":
		obj["phone"] = conn.GetSession("phone")
		obj["smsCode"] = conn.GetSession("smsCode")
	}
	obj["currentUserUUID"] = conn.GetSession("currentUserUUID")
}

// setSession 根据 graphql.Params.RootObject 中 setSession 对应的
// string 数组内容，从 RootObject 中取值并放入 Session
func setSession(conn *GQLController, obj gqlRootObject) {
	if fields := obj["setSession"]; fields != nil {
		for _, field := range fields.([]string) {
			conn.SetSession(field, obj[field].(string))
		}
	}
}
