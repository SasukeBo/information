package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "设备资讯化"
	c.Data["Name"] = "SasukeBo"
	c.Data["Email"] = "809754210@qq.com"
	c.Data["path"] = "static"

	if v := c.GetSession("freshCount"); v == nil {
		c.SetSession("freshCount", int(2))
		c.Data["freshCount"] = 1
	} else {
		c.SetSession("freshCount", v.(int)+1)
		c.Data["freshCount"] = v.(int)
	}

	c.TplName = "index.html"
}
