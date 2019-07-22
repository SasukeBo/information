package models

import (
	"time"
)

// Privilege 角色模型
type Privilege struct {
	ID        int         `orm:"auto;pk;column(id)"`
	PrivName  string      `orm:"unique"`
	PrivType  PrivType    `orm:"default(0)"`
	Status    BaseStatus  `orm:"default(0)"`
	RolePriv  []*RolePriv `orm:"reverse(many)"`
	CreatedAt time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time   `orm:"auto_now;type(datetime)"`
}
