package models

import (
	"github.com/SasukeBo/information/models/errors"
)

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

// Get _
func (rp *RolePriv) Get() error {
	if err := Repo.Read(rp); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "get role_priv error",
			OriErr:  err,
		}
	}

	return nil
}

// Insert _
func (rp *RolePriv) Insert() error {
	if _, err := Repo.Insert(rp); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "insert role_priv error",
			OriErr:  err,
		}
	}

	return nil
}

// Delete _
func (rp *RolePriv) Delete() error {
	if err := rp.Get(); err != nil {
		return err
	}

	if _, err := Repo.Delete(rp); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "delete role_priv error",
			OriErr:  err,
		}
	}

	return nil
}

// LoadPrivilege related load privilege
func (rp *RolePriv) LoadPrivilege() (*Privilege, error) {
	if _, err := Repo.LoadRelated(rp, "Privilege"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "role_priv load privilege error",
			OriErr:  err,
		}
	}

	return rp.Privilege, nil
}

// LoadRole related load role
func (rp *RolePriv) LoadRole() (*Role, error) {
	if _, err := Repo.LoadRelated(rp, "Role"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "role_priv load role error",
			OriErr:  err,
		}
	}

	return rp.Role, nil
}
