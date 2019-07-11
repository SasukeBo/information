package controllers

import (
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego"
)

// AccountController doc false
type AccountController struct {
	beego.Controller
}

// Post doc false
func (c *AccountController) Post() {
	a := new(models.Account)
	a.Name = c.GetString("username")
	a.Insert()

	c.Data["message"] = "insert success, name: " + a.Name
	c.TplName = "result.html"
}

// Get doc false
func (c *AccountController) Get() {
	c.TplName = "account.html"
}
