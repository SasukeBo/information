package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/SasukeBo/information/utils"
)

// Role 角色模型
type Role struct {
	ID        int         `orm:"auto;pk;column(id)"`
	RoleName  string      `orm:"unique"`        // 角色名称
	Status    int         `orm:"default(0)"`    // 基础状态
	RolePriv  []*RolePriv `orm:"reverse(many)"` // 角色权限关联关系
	CreatedAt time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time   `orm:"auto_now;type(datetime)"`
}

// GetBy _
func (r *Role) GetBy(col string) error {
	if err := Repo.Read(r, col); err != nil {
		return utils.ORMError{
			Message: fmt.Sprintf("role get by %s error", col),
			OrmErr:  err,
		}
	}

	return nil
}

// Delete _
func (r *Role) Delete() error {
	if err := r.GetBy("id"); err != nil {
		return err
	}

	if _, err := Repo.Delete(r); err != nil {
		return utils.ORMError{
			Message: "role delete by id error",
			OrmErr:  err,
		}
	}

	return nil
}

// Update _
func (r *Role) Update(cols ...string) error {
	if _, err := Repo.Update(r, cols...); err != nil {
		return utils.ORMError{
			Message: "role update error",
			OrmErr:  err,
		}
	}

	return nil
}

// LoadRolePriv _
func (r *Role) LoadRolePriv() ([]*RolePriv, error) {
	if _, err := Repo.LoadRelated(r, "RolePriv"); err != nil {
		return nil, utils.ORMError{
			Message: "related load role_priv error",
			OrmErr:  err,
		}
	}

	return r.RolePriv, nil
}

// Validate _
func (r *Role) Validate(sign string) error {
	qs := Repo.QueryTable("role_priv").Filter("role_id", r.ID).Filter("privilege__sign", sign)
	var rp RolePriv
	if err := qs.One(&rp, "id"); err == orm.ErrNoRows {
		return utils.LogicError{
			Message: fmt.Sprintf("without %s ability", sign),
		}
	} else if err != nil {
		return utils.LogicError{
			Message: "unknown error",
		}
	}

	return nil
}
