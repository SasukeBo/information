package models

import (
	// "time"
	"github.com/SasukeBo/information/utils"
)

// RolePriv 角色权限关联关系模型
type RolePriv struct {
	ID        int        `orm:"auto;pk;column(id)"`
	Role      *Role      `orm:"rel(fk);on_delete()"` // 关联角色，删除时删除
	Privilege *Privilege `orm:"rel(fk);on_delete()"` // 关联权利，删除时删除
}

// Get _
func (rp *RolePriv) Get() error {
	if err := Repo.Read(rp); err != nil {
		return utils.ORMError{
			Message: "role_priv get error",
			OrmErr:  err,
		}
	}

	return nil
}

// LoadPrivilege related load privilege
func (rp *RolePriv) LoadPrivilege() (*Privilege, error) {
	if _, err := Repo.LoadRelated(rp, "Privilege"); err != nil {
		return nil, utils.ORMError{
			Message: "role_priv load related privilege error",
			OrmErr:  err,
		}
	}

	return rp.Privilege, nil
}

// LoadRole related load role
func (rp *RolePriv) LoadRole() (*Role, error) {
	if _, err := Repo.LoadRelated(rp, "Role"); err != nil {
		return nil, utils.ORMError{
			Message: "role_priv load related role error",
			OrmErr:  err,
		}
	}

	return rp.Role, nil
}
