package main

import (
	_ "github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/proto"
	_ "github.com/SasukeBo/information/routers"
	"github.com/astaxie/beego"
)

func main() {
	go proto.Run()
	beego.Run()
}
