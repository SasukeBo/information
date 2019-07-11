package controllers

import (
	"github.com/astaxie/beego"
)

// MainController doc false
type MainController struct {
	beego.Controller
}

// Get http method
func (c *MainController) Get() {
	c.TplName = "index.html"
}
