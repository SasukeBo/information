package models

import (
	"fmt"

	"github.com/SasukeBo/information/models/errors"
)

// PrivType 权限类型
var PrivType = struct {
	Default int
	Admin   int
	Device  int
}{0, 1, 2}

// Privilege 权限模型
type Privilege struct {
	Name     string // 权限名称
	ID       int    `orm:"auto;pk;column(id)"`
	Sign     string // 权限签名
	PrivType int    `orm:"default(0)"` // 权限类型
}

// TableUnique 自定义唯一键
func (p *Privilege) TableUnique() [][]string {
	return [][]string{
		[]string{"sign", "priv_type"},
	}
}

// Get get privilege by id
func (p *Privilege) Get() error {
	if err := Repo.Read(p); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "get privilege error",
			OriErr:  err,
		}
	}

	return nil
}

// GetBy get privilege by col
func (p *Privilege) GetBy(col string) error {
	if err := Repo.Read(p, col); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   col,
			Message: fmt.Sprintf("get privilege by %s error", col),
			OriErr:  err,
		}
	}

	return nil
}

// Delete _
func (p *Privilege) Delete() error {
	if err := p.Get(); err != nil {
		return err
	}

	if _, err := Repo.Delete(p); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "delete privilege error",
			OriErr:  err,
		}
	}

	return nil
}
