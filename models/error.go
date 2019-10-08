package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var env = beego.AppConfig.String

// Error _
type Error struct {
	Message string
	OriErr  error
}

// Error _
func (e Error) Error() string {
	if env("runmode") == "dev" {
		logs.Error("%s %v", e.Message, e.OriErr)
	}

	return e.Message
}
