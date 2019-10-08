package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

// Role 角色模型
type Role struct {
	ID        int         `orm:"auto;pk;column(id)"`
	RoleName  string      `orm:"unique"`         // 角色名称
	Status    int         `orm:"default(0)"`     // 基础状态
	IsAdmin   bool        `orm:"default(false)"` // 是否为管理员角色
	RolePriv  []*RolePriv `orm:"reverse(many)"`  // 角色权限关联关系
	CreatedAt time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time   `orm:"auto_now;type(datetime)"`
}

// LoadRolePriv _
func (r *Role) LoadRolePriv() ([]*RolePriv, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(r, "RolePriv"); err != nil {
		return nil, Error{Message: "load related role_priv failed.", OriErr: err}
	}

	return r.RolePriv, nil
}

// Validate _
func (r *Role) Validate(sign string, privType int) error {
	o := orm.NewOrm()
	qs := o.QueryTable("role_priv").Filter("role_id", r.ID).Filter("privilege__sign", sign).Filter("privilege__priv_type", privType)

	var rp RolePriv
	if err := qs.One(&rp, "id"); err != nil {
		return Error{Message: fmt.Sprintf("can't access without %s ability", sign), OriErr: err}
	}

	return nil
}
