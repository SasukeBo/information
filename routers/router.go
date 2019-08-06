package routers

import (
	"golang.org/x/net/websocket"

	"github.com/SasukeBo/information/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/*", &controllers.MainController{})
	beego.Router("/graphql", &controllers.GQLController{})
	beego.Handler("/websocket", websocket.Handler(controllers.Connect))

	beego.InsertFilter("/*", beego.BeforeExec, controllers.AuthFilter)
	beego.InsertFilter("/*", beego.AfterExec, controllers.CleanAuthErrorFilter, false)
}
