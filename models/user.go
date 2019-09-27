package models

import (
	"fmt"
	"time"

	"github.com/SasukeBo/information/models/errors"
)

// User 用户模型
type User struct {
	AvatarURL string    `orm:"column(avatar_url);null"` // 头像链接
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Email     string    `orm:"null"`
	ID        int       `orm:"auto;pk;column(id)"`
	Name      string    // 真实姓名
	Password  string
	Phone     string    `orm:"unique"`
	Role      *Role     `orm:"rel(fk);null;on_delete(set_null)"` // 用户角色，删除时置空
	Status    int       `orm:"default(0)"`                       // 基础状态
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	UUID      string    `orm:"column(uuid);unique;index"`
}

// GetBy get user by col
func (u *User) GetBy(col string) error {
	if err := Repo.Read(u, col); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   col,
			Message: fmt.Sprintf("get user by %s error", col),
			OriErr:  err,
		}
	}

	return nil
}

// Insert _
func (u *User) Insert() error {
	if _, err := Repo.Insert(u); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "insert user error",
			OriErr:  err,
		}
	}

	return nil
}

// Update _
func (u *User) Update(cols ...string) error {
	if _, err := Repo.Update(u, cols...); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "update user error",
			OriErr:  err,
		}
	}

	return nil
}

// LoadRole related load role
func (u *User) LoadRole() (*Role, error) {
	if _, err := Repo.LoadRelated(u, "Role"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "user load role error",
			OriErr:  err,
		}
	}

	return u.Role, nil
}
