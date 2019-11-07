package models

import "github.com/astaxie/beego/orm"

// RolePriv 角色权限关联关系模型
type RolePriv struct {
	ID        int        `orm:"auto;pk;column(id)"`
	Role      *Role      `orm:"rel(fk);on_delete()"` // 关联角色，删除时删除
	Privilege *Privilege `orm:"rel(fk);on_delete()"` // 关联权利，删除时删除
}

// TableUnique 自定义唯一键
func (rp *RolePriv) TableUnique() [][]string {
	return [][]string{
		[]string{"role_id", "privilege_id"},
	}
}

// LoadPrivilege related load privilege
func (rp *RolePriv) LoadPrivilege() (*Privilege, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(rp, "Privilege"); err != nil {
		return nil, Error{Message: "load related privilege failed.", OriErr: err}
	}

	return rp.Privilege, nil
}

// LoadRole related load role
func (rp *RolePriv) LoadRole() (*Role, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(rp, "Role"); err != nil {
		return nil, Error{Message: "load related role failed", OriErr: err}
	}

	return rp.Role, nil
}
