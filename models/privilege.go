package models

import (
	// "time"
	"github.com/SasukeBo/information/utils"
)

// PrivType 权限类型
var PrivType = struct {
	Default int
}{0}

// Privilege 权限模型
type Privilege struct {
	ID       int    `orm:"auto;pk;column(id)"`
	PrivName string `orm:"unique"`     // 权限名称
	PrivType int    `orm:"default(0)"` // 权限类型
	Status   int    `orm:"default(0)"` // 基础状态
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
