package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	// postgres driver
	_ "github.com/lib/pq"
)

func init() {
	env := beego.AppConfig.String
	if env("runmode") == "dev" {
		orm.Debug = true
	}

	dbconfig := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		env("dbusername"), env("dbpassword"), env("dbname"), env("dbhost"),
		env("dbport"), env("dbsslmode"),
	)

	orm.RegisterDataBase("default", "postgres", dbconfig)
	orm.RegisterDriver("postgres", orm.DRPostgres)
	// 注册 model
	orm.RegisterModel(new(Account))

	// 自动建表
	if err := orm.RunSyncdb("default", false, true); err != nil {
		logs.Error("Create table failed!", err)
	}
}
