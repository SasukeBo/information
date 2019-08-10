package models

import (
	"fmt"

	"github.com/SasukeBo/information/utils"
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
		return utils.ORMError{
			Message: "user_extend get error",
			OrmErr:  err,
		}
	}

	return nil
}

// GetBy _
func (ue *UserExtend) GetBy(col string) error {
	if err := Repo.Read(ue, col); err != nil {
		return utils.ORMError{
			Message: fmt.Sprintf("user_extend get by %s error", col),
			OrmErr:  err,
		}
	}

	return nil
}

// Update _
func (ue *UserExtend) Update(cols ...string) error {
	if _, err := Repo.Update(ue, cols...); err != nil {
		return utils.ORMError{
			Message: "user_extend update error",
			OrmErr:  err,
		}
	}

	if err := ue.Get(); err != nil {
		return err
	}

	return nil
}

// LoadUser _
func (ue *UserExtend) LoadUser() (*User, error) {
	if _, err := Repo.LoadRelated(ue, "User"); err != nil {
		return nil, utils.ORMError{
			Message: "user_extend load related user error",
			OrmErr:  err,
		}
	}

	return ue.User, nil
}
