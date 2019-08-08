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

/* 									models begin
---------------------------------------------------- */

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

// Privilege 权限模型
type Privilege struct {
  ID       int    `orm:"auto;pk;column(id)"`
  PrivName string `orm:"unique"`     // 权限名称
  PrivType int    `orm:"default(0)"` // 权限类型
  Status   int    `orm:"default(0)"` // 基础状态
}

// PrivType 权限类型
var PrivType = struct {
  Default int
}{0}

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

// DeviceStatus 设备状态
var DeviceStatus = struct {
  Prod    int // 生产
  Stop    int // 停机
  OffLine int // 离线
}{0, 1, 2}

// DeviceParamValue 设备参数值模型
type DeviceParamValue struct {
  Value       string       // 参数值
  ID          int          `orm:"auto;pk;column(id)"`
  DeviceParam *DeviceParam `orm:"rel(fk);on_delete()"`
  CreatedAt   time.Time    `orm:"auto_now_add;type(datetime)"`
}

// DeviceStatusLog 设备运行状态变更模型
type DeviceStatusLog struct {
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
    new(DeviceStatusLog),
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
