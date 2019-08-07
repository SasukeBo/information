package routers

import (
	"golang.org/x/net/websocket"

	"github.com/SasukeBo/information/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/*", &controllers.MainController{})
	beego.InsertFilter("/*", beego.BeforeExec, controllers.AuthFilter)
	beego.InsertFilter("/*", beego.AfterExec, controllers.CleanAuthErrorFilter, false)

	beego.Router("/graphql", &controllers.GQLController{})
	beego.InsertFilter("/graphql", beego.BeforeExec, controllers.HandleGraphql)

	beego.Router("/admin_graphql", &controllers.AdminGQLController{})
	beego.InsertFilter("/admin_graphql", beego.BeforeExec, controllers.HandleGraphql)

	beego.Handler("/websocket", websocket.Handler(controllers.Connect))
}
