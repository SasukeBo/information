package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	// postgres driver
	_ "github.com/lib/pq"
)

var repo orm.Ormer

// BaseStatus 基础状态类型
type BaseStatus int

const (
	// Default 初始状态
	Default BaseStatus = iota
	// Publish 发布状态
	Publish
	// Block 屏蔽状态
	Block
	// Deleted 删除状态
	Deleted
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

	err := orm.RegisterDataBase("default", "postgres", dbconfig)
	handleError(err)
	err = orm.RegisterDriver("postgres", orm.DRPostgres)
	handleError(err)
	// 注册 model
	orm.RegisterModel(new(Role), new(UserProfile), new(User))

	// 自动建表
	if err := orm.RunSyncdb("default", false, true); err != nil {
		logs.Error("Create table failed!", err)
	}

	repo = orm.NewOrm()
}

func handleError(err error) {
	if err != nil {
		logs.Error(err)
	}
}
