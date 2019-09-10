package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema"
)

// GQLController is graphql controller
type GQLController struct {
	beego.Controller
}

// Get http method
func (conn *GQLController) Get() {
	graphiqlGet(&conn.Controller)
}

// Post http method
func (conn *GQLController) Post() {
	rootObject := gqlRootObject{}

	var result *graphql.Result
	needAuth := conn.Ctx.Input.GetData("need_auth").(bool)

	if err := conn.GetSession("auth_error"); err != nil && needAuth {
		result = genGQLError(err.(error))
	} else if err = conn.Ctx.Input.GetData("gql_error"); err != nil {
		result = genGQLError(err.(error))
	} else {
		params := conn.Ctx.Input.GetData("query_params").(queryParams)
		gqlGetSession(&conn.Controller, rootObject, params.RootFieldName)

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
	conn.Data["json"] = result
	conn.ServeJSON()
}
