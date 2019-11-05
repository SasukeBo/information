package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Role 角色模型
type Role struct {
	ID        int         `orm:"auto;pk;column(id)"`
	RoleName  string      `orm:"unique"`         // 角色名称
	Status    int         `orm:"default(0)"`     // 基础状态
	IsAdmin   bool        `orm:"default(false)"` // 是否为管理员角色
	RolePrivs []*RolePriv `orm:"reverse(many)"`  // 角色权限关联关系
	CreatedAt time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time   `orm:"auto_now;type(datetime)"`
}

// LoadPrivilege _
func (r *Role) LoadPrivilege() ([]*Privilege, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(r, "RolePrivs"); err != nil {
		return nil, Error{Message: "load related role_priv failed.", OriErr: err}
	}

	rolePrivs := r.RolePrivs
	privs := []*Privilege{}
	for _, p := range rolePrivs {
		if _, err := o.LoadRelated(p, "Privilege"); err != nil {
			return nil, Error{Message: "load related privilege failed.", OriErr: err}
		}
		privs = append(privs, p.Privilege)
	}

	return privs, nil
}

// Validate 校验角色是否具备权限
// sign - 权限sign
// privType - 权限类型 see models.PrivType
func (r *Role) Validate(sign string, privType int) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("role_priv").Filter("role_id", r.ID).Filter("privilege__sign", sign).Filter("privilege__priv_type", privType)

	var rp RolePriv
	if err := qs.One(&rp, "id"); err != nil {
		return false
	}

	return true
}
