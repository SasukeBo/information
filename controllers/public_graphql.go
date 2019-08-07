package controllers

import (
	"github.com/SasukeBo/information/schema"
	"github.com/astaxie/beego"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

// GQLController is graphql controller
type GQLController struct {
	beego.Controller
}

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
	graphiqlGet(&conn.Controller)
}

// Post http method
func (conn *GQLController) Post() {
	needAuth := true
	rootObject := gqlRootObject{}

	var result *graphql.Result
	params := fetchParams(&conn.Controller)

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
		gqlGetSession(&conn.Controller, rootObject, params.OperationName)

		gqlParams := graphql.Params{
			Schema:         schema.PublicSchema,
			RequestString:  params.Query,
			OperationName:  params.OperationName,
			VariableValues: params.Variables,
			RootObject:     rootObject,
		}

		result = graphql.Do(gqlParams)

		gqlSetSession(&conn.Controller, gqlParams.RootObject)
	}
	// conn.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:9080")
	conn.Data["json"] = result
	conn.ServeJSON()
}
