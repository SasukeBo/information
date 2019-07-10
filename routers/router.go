package routers

import (
	"github.com/SasukeBo/information/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
