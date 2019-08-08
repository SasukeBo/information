package models

import (
	"github.com/SasukeBo/information/utils"
)

// UserExtend 用户信息模型
type UserExtend struct {
	ID    int     `orm:"auto;pk;column(id)"`
	User  *User   `orm:"reverse(one);on_delete()"` // 用户删除时删除
	Name  string  `orm:"null"`                     // 真实姓名
	Email *string `orm:"unique;null"`
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
