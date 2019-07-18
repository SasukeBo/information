package routers

import (
	"github.com/SasukeBo/information/controllers"
	"github.com/SasukeBo/information/schema"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/*", &controllers.MainController{})
	beego.Handler("/graphql", schema.GraphqlHander, true)
}
