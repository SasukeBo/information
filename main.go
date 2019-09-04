package main

import (
	_ "github.com/SasukeBo/information/models"
	_ "github.com/SasukeBo/information/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
