package models

import (
	"fmt"
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
	message := fmt.Sprintf(`{"message": %s, "originError": %v}`, e.Message, e.OriErr)
	if env("runmode") == "dev" {
		logs.Error(message)
	}

	return message
}
