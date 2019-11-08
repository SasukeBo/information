package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	// postgres driver
	_ "github.com/lib/pq"
)

// BaseStatus 基础状态类型
var BaseStatus = struct {
	Default int
	Publish int
	Block   int
	Deleted int
}{0, 1, 2, 3}

func init() {
	env := beego.AppConfig.String
	if env("runmode") == "dev" {
		orm.Debug = false
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
	orm.RegisterModel(
		new(Role),
		new(User),
		new(RolePriv),
		new(Privilege),
		new(UserLogin),
		new(Device),
		new(SystemConf),
		new(DeviceStatusLog),
		new(StopReason),
		new(DeviceProductShip),
		new(Product),
		new(ProductIns),
		new(DetectItem),
		new(DetectItemValue),
		// new(DeviceCharger),
	)

	// 自动建表
	if err := orm.RunSyncdb("default", false, true); err != nil {
		logs.Error("Create table failed!", err)
	}
}

func handleError(err error) {
	if err != nil {
		logs.Error(err)
	}
}
