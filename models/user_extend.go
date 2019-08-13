package models

import (
	"fmt"

	"github.com/SasukeBo/information/models/errors"
)

// UserExtend 用户信息模型
type UserExtend struct {
	ID    int    `orm:"auto;pk;column(id)"`
	User  *User  `orm:"reverse(one);on_delete()"` // 用户删除时删除
	Name  string `orm:"null"`                     // 真实姓名
	Email string `orm:"unique;null"`
}

// Get _
func (ue *UserExtend) Get() error {
	if err := Repo.Read(ue); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "UserExtend",
			Message: "Get() error",
			OriErr:  err,
		}
	}

	return nil
}

// GetBy _
func (ue *UserExtend) GetBy(col string) error {
	if err := Repo.Read(ue, col); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "UserExtend",
			Message: fmt.Sprintf("GetBy(%s) error", col),
			OriErr:  err,
		}
	}

	return nil
}

// Insert _
func (ue *UserExtend) Insert() error {
	if _, err := Repo.Insert(ue); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "UserExtend",
			Message: "Insert() error",
			OriErr:  err,
		}
	}

	return nil
}

// Update _
func (ue *UserExtend) Update(cols ...string) error {
	if _, err := Repo.Update(ue, cols...); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "UserExtend",
			Message: "Update() error",
			OriErr:  err,
		}
	}

	return nil
}

// LoadUser _
func (ue *UserExtend) LoadUser() (*User, error) {
	if _, err := Repo.LoadRelated(ue, "User"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Field:   "UserExtend",
			Message: "LoadUser() error",
			OriErr:  err,
		}
	}

	return ue.User, nil
}
