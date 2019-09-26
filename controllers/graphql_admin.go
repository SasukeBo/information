package controllers

import (
	"github.com/astaxie/beego"
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schemaadmin"
)

// AdminGQLController is graphql controller
type AdminGQLController struct {
	beego.Controller
}

// Get http method
func (conn *AdminGQLController) Get() {
	graphiqlGet(&conn.Controller)
}

// Post http method
func (conn *AdminGQLController) Post() {
	rootObject := gqlRootObject{}

	var result *graphql.Result

	if err := conn.GetSession("auth_error"); err != nil {
		result = genGQLError(err.(error))
	} else if err = conn.Ctx.Input.GetData("gql_error"); err != nil {
		result = genGQLError(err.(error))
	} else {
		params := conn.Ctx.Input.GetData("query_params").(queryParams)
		gqlGetSession(&conn.Controller, rootObject, params.RootFieldName)

		gqlParams := graphql.Params{
			Schema:         schemaadmin.Root,
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
