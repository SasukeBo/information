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
var BaseStatus = struct {
	Default int
	Publish int
	Block   int
	Deleted int
}{0, 1, 2, 3}

// PrivType 权限类型
var PrivType = struct {
	Default int
}{0}

/* 									models begin
---------------------------------------------------- */

// User 用户模型
type User struct {
	Password   string
	ID         int         `orm:"auto;pk;column(id)"`
	UUID       string      `orm:"column(uuid);unique;index"`
	Phone      string      `orm:"unique"`
	AvatarURL  string      `orm:"column(avatar_url);null"`          // 头像链接
	Role       *Role       `orm:"rel(fk);null;on_delete(set_null)"` // 用户角色，删除时置空
	UserExtend *UserExtend `orm:"rel(one)"`                         // 用户信息拓展
	Status     int         `orm:"default(0)"`                       // 基础状态
	CreatedAt  time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time   `orm:"auto_now;type(datetime)"`
}

// UserLogin 用户登录模型
// 用户请求到达服务器时，
type UserLogin struct {
	EncryptedPasswd string    // 加密后的密码
	UserAgent       string    // 用户代理
	ID              int       `orm:"auto;pk;column(id)"`
	User            *User     `orm:"rel(fk);on_delete()"`       // 用户删除时删除
	RemoteIP        string    `orm:"column(remote_ip)"`         // 登录IP
	SessionID       string    `orm:"column(session_id);unique"` // session id
	Remembered      bool      `orm:"default(true)"`             // 记住登录
	Logout          bool      `orm:"default(false)"`            // 是否登出
	CreatedAt       time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt       time.Time `orm:"auto_now;type(datetime)"`
}

// UserExtend 用户信息模型
type UserExtend struct {
	ID    int     `orm:"auto;pk;column(id)"`
	User  *User   `orm:"reverse(one);on_delete()"` // 用户删除时删除
	Name  string  `orm:"null"`                     // 真实姓名
	Email *string `orm:"unique;null"`
}

// Privilege 角色模型
type Privilege struct {
	ID        int       `orm:"auto;pk;column(id)"`
	PrivName  string    `orm:"unique"`     // 权限名称
	PrivType  int       `orm:"default(0)"` // 权限类型
	Status    int       `orm:"default(0)"` // 基础状态
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

// RolePriv 角色权限关联关系模型
type RolePriv struct {
	ID        int        `orm:"auto;pk;column(id)"`
	Role      *Role      `orm:"rel(fk);on_delete()"` // 关联角色，删除时删除
	Privilege *Privilege `orm:"rel(fk);on_delete()"` // 关联权利，删除时删除
}

// Role 角色模型
type Role struct {
	ID        int         `orm:"auto;pk;column(id)"`
	RoleName  string      `orm:"unique"`        // 角色名称
	Status    int         `orm:"default(0)"`    // 基础状态
	RolePriv  []*RolePriv `orm:"reverse(many)"` // 角色权限关联关系
	CreatedAt time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time   `orm:"auto_now;type(datetime)"`
}

// Device 设备模型
type Device struct {
	Type        string    // 类型
	Name        string    // 设备名称
	Mac         string    // 设备Mac地址
	Token       string    // 设备Token，用于数据加密
	Status      int       `orm:"default(0)"`                // 基础状态
	ID          int       `orm:"auto;pk;column(id)"`        // PKey 主键
	UUID        string    `orm:"column(uuid);unique;index"` // 通用唯一标识符
	User        *User     `orm:"rel(fk)"`                   // 注册人
	Description string    `orm:"null"`                      // 描述
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`
}

// DeviceCharge 设备负责人关系模型
type DeviceCharge struct {
	ID        int       `orm:"auto;pk;column(id)"`
	User      *User     `orm:"rel(fk);on_delete()"` // 设备负责人，用户删除时删除
	Device    *Device   `orm:"rel(fk);on_delete()"` // 设备，删除时删除
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

// DeviceChargeAbility 设备负责人权限模型
type DeviceChargeAbility struct {
	ID           int           `orm:"auto;pk;column(id)"`
	DeviceCharge *DeviceCharge `orm:"rel(fk);on_delete()"` // 设备负责关系删除时删除
	Privilege    *Privilege    `orm:"rel(fk);on_delete()"` // 权限删除时删除
}

// DeviceParam 设备参数模型
type DeviceParam struct {
	Name      string    // 参数名称
	Sign      string    // 参数签名（标识），要求英文及数字组合的字符串
	Type      string    // 参数值类型，string？int？bool？
	ID        int       `orm:"auto;pk;column(id)"`
	Author    *User     `orm:"rel(fk);null;on_delete(set_null)"` // 创建人，删除时置空
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// DeviceParamValue 设备参数值模型
type DeviceParamValue struct {
	Value       string       // 参数值
	ID          int          `orm:"auto;pk;column(id)"`
	DeviceParam *DeviceParam `orm:"rel(fk);on_delete()"`
	CreatedAt   time.Time    `orm:"auto_now_add;type(datetime)"`
}

// DeviceStatus 设备运行状态变更模型
type DeviceStatus struct {
	Status   int       // 设备运行状态
	ID       int       `orm:"auto;pk;column(id)"`
	Device   *Device   `orm:"rel(fk);on_delete()"`
	ChangeAt time.Time `orm:"auto_now_add;type(datetime)"`
}

/* 									models end
---------------------------------------------------- */

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
		new(UserLogin),
		new(Device),
		new(DeviceCharge),
		new(DeviceChargeAbility),
		new(DeviceParam),
		new(DeviceParamValue),
		new(DeviceStatus),
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
