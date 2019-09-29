package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	// postgres driver
	_ "github.com/lib/pq"
)

// Repo is Ormer
var Repo orm.Ormer

// LogicError Error with logic type
type LogicError struct {
	Type    string
	Field   string
	Message string
	OriErr  error
}

func (e LogicError) Error() string {
	var oriMsg = ""
	if e.OriErr != nil {
		oriMsg = e.OriErr.Error()
	}

	return fmt.Sprintf(
		`{"type": "%s Error", "field": "%s", "message": "%s", "originMessage": "%s"}`,
		e.Type,
		e.Field,
		e.Message,
		oriMsg,
	)
}

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
	orm.RegisterModel(
		new(Role),
		new(User),
		new(RolePriv),
		new(Privilege),
		new(UserLogin),
		new(Device),
		new(DeviceCharger),
		new(DeviceStatusLog),
		new(DeviceProductShip),
		new(Product),
		new(ProductIns),
		new(DetectItem),
		new(DetectItemValue),
	)

	// 自动建表
	if err := orm.RunSyncdb("default", false, true); err != nil {
		logs.Error("Create table failed!", err)
	}

	Repo = orm.NewOrm()
}

// NewCond return a orm.Condition
func NewCond() *orm.Condition {
	return orm.NewCondition()
}

func handleError(err error) {
	if err != nil {
		logs.Error(err)
	}
}
