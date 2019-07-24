package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	// postgres driver
	_ "github.com/lib/pq"
)

// Repo is Ormer
var Repo orm.Ormer

// BaseStatus 基础状态类型
type BaseStatus int

// PrivType 权限类型
type PrivType int

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
	orm.RegisterModel(
		new(Role),
		new(UserExtend),
		new(User),
		new(RolePriv),
		new(Privilege),
	)

	// 自动建表
	if err := orm.RunSyncdb("default", false, true); err != nil {
		logs.Error("Create table failed!", err)
	}

	Repo = orm.NewOrm()
}

func handleError(err error) {
	if err != nil {
		logs.Error(err)
	}
}

// User 用户模型
type User struct {
	ID         int    `orm:"auto;pk;column(id)"`
	UUID       string `orm:"column(uuid);unique;index"`
	Phone      string `orm:"unique"`
	Password   string
	AvatarURL  string      `orm:"column(avatar_url);null"`
	Role       *Role       `orm:"rel(fk);null"`
	UserExtend *UserExtend `orm:"rel(one)"`
	Status     BaseStatus  `orm:"default(0)"`
	CreatedAt  time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time   `orm:"auto_now;type(datetime)"`
}

// UserExtend 用户信息模型
type UserExtend struct {
	ID    int    `orm:"auto;pk;column(id)"`
	User  *User  `orm:"reverse(one)"`
	Name  string `orm:"null"`
	Email string `orm:"unique;null"`
}

// Privilege 角色模型
type Privilege struct {
	ID        int         `orm:"auto;pk;column(id)"`
	PrivName  string      `orm:"unique"`
	PrivType  PrivType    `orm:"default(0)"`
	Status    BaseStatus  `orm:"default(0)"`
	RolePriv  []*RolePriv `orm:"reverse(many)"`
	CreatedAt time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time   `orm:"auto_now;type(datetime)"`
}

// RolePriv 角色权限关联关系模型
type RolePriv struct {
	ID        int        `orm:"auto;pk;column(id)"`
	Role      *Role      `orm:"rel(fk)"`
	Privilege *Privilege `orm:"rel(fk)"`
}

// Role 角色模型
type Role struct {
	ID        int         `orm:"auto;pk;column(id)"`
	RoleName  string      `orm:"unique"`
	Status    BaseStatus  `orm:"default(0)"`
	RolePriv  []*RolePriv `orm:"reverse(many)"`
	CreatedAt time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time   `orm:"auto_now;type(datetime)"`
}
