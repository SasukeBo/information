package models

import (
	// "time"
	"github.com/SasukeBo/information/utils"
)

// PrivType 权限类型
var PrivType = struct {
	Default int
	Device  int
}{0, 1}

// Privilege 权限模型
type Privilege struct {
	Name     string // 权限名称
	ID       int    `orm:"auto;pk;column(id)"`
	Sign     string `orm:"unique"`     // 权限签名
	PrivType int    `orm:"default(0)"` // 权限类型
}

// Get get privilege by id
func (p *Privilege) Get() error {
	if err := Repo.Read(p); err != nil {
		return utils.ORMError{
			Message: "privilege get error",
			OrmErr:  err,
		}
	}

	return nil
}
