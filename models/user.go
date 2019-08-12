package models

import (
	"fmt"
	"time"

	"github.com/SasukeBo/information/errors"
)

// User 用户模型
type User struct {
	Password   string
	ID         int         `orm:"auto;pk;column(id)"`
	UUID       string      `orm:"column(uuid);unique;index"`
	Phone      string      `orm:"unique"`
	AvatarURL  string      `orm:"column(avatar_url);null"`          // 头像链接
	Role       *Role       `orm:"rel(fk);null;on_delete(set_null)"` // 用户角色，删除时置空
	UserExtend *UserExtend `orm:"rel(one)"`                         // 用户信息拓展
	Status     int         `orm:"default(0)"`                       // 基础状态
	CreatedAt  time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time   `orm:"auto_now;type(datetime)"`
}

// GetBy get user by col
func (u *User) GetBy(col string) error {
	if err := Repo.Read(u, col); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "User",
			Message: fmt.Sprintf("GetBy(%s) error", col),
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
			Field:   "User",
			Message: "Insert() error",
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
			Field:   "User",
			Message: "Update() error",
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
			Field:   "User",
			Message: "LoadRole() error",
			OriErr:  err,
		}
	}

	return u.Role, nil
}

// LoadUserExtend related load user_extend
func (u *User) LoadUserExtend() (*UserExtend, error) {
	if _, err := Repo.LoadRelated(u, "UserExtend"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Field:   "User",
			Message: "LoadUserExtend() error",
			OriErr:  err,
		}
	}

	return u.UserExtend, nil
}
