package models

import (
// "time"
// "github.com/SasukeBo/information/utils"
)

// RolePriv 角色权限关联关系模型
type RolePriv struct {
	ID        int        `orm:"auto;pk;column(id)"`
	Role      *Role      `orm:"rel(fk);on_delete()"` // 关联角色，删除时删除
	Privilege *Privilege `orm:"rel(fk);on_delete()"` // 关联权利，删除时删除
}
