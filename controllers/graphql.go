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

// Get http method
func (conn *GQLController) Get() {
	queryString := conn.Ctx.Input.Query("query")
	variablesString := conn.Ctx.Input.Query("variables")
	graphiqlVersion := "0.11.11"
	conn.TplName = "graphql.html"
	conn.Data["GraphiqlVersion"] = graphiqlVersion
	conn.Data["QueryString"] = queryString
	conn.Data["VariablesString"] = variablesString
}

// Post http method
func (conn *GQLController) Post() {
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	json.NewDecoder(conn.Ctx.Request.Body).Decode(&params)

	currentUser := conn.GetSession("current_user")

	rootObject := map[string]interface{}{
		"currentUser": currentUser,
	}

	gqlParams := graphql.Params{
		Schema:         schema.Schema,
		RequestString:  params.Query,
		VariableValues: params.Variables,
		OperationName:  params.OperationName,
		RootObject:     rootObject,
	}

	result := graphql.Do(gqlParams)
	conn.SetSession("current_user", gqlParams.RootObject["currentUser"])

	conn.Data["json"] = result
	conn.ServeJSON()
}
