package controllers

import (
	"github.com/SasukeBo/information/schema"
	"github.com/astaxie/beego"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
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
	params := fetchParams(&conn.Controller)

	if err := conn.GetSession("auth_error"); err != nil {
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
			Schema:         schema.AdminSchema,
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
