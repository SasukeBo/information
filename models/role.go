package models

import (
	"time"

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

// Get _
func (r *Role) Get() error {
	if err := Repo.Read(r); err != nil {
		return utils.ORMError{
			Message: "role get error",
			OrmErr:  err,
		}
	}

	return nil
}
