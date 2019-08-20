package main

import (
	"github.com/SasukeBo/information/models"
	_ "github.com/SasukeBo/information/routers"
	"github.com/astaxie/beego"
)

func main() {
	go models.RunTCP()
	beego.Run()
}
