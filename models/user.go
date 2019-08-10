package models

import (
	"fmt"
	"time"

	"github.com/SasukeBo/information/utils"
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
		return utils.ORMError{
			Message: fmt.Sprintf("user get by %s error", col),
			OrmErr:  err,
		}
	}

	return nil
}

// Update _
func (u *User) Update(cols ...string) error {
	if _, err := Repo.Update(u, cols...); err != nil {
		return utils.ORMError{
			Message: "user update error",
			OrmErr:  err,
		}
	}

	return nil
}

// LoadRole related load role
func (u *User) LoadRole() (*Role, error) {
	if _, err := Repo.LoadRelated(u, "Role"); err != nil {
		return nil, utils.ORMError{
			Message: "user load related role error",
			OrmErr:  err,
		}
	}

	return u.Role, nil
}

// LoadUserExtend related load user_extend
func (u *User) LoadUserExtend() (*UserExtend, error) {
	if _, err := Repo.LoadRelated(u, "UserExtend"); err != nil {
		return nil, utils.ORMError{
			Message: "user load related user_extend error",
			OrmErr:  err,
		}
	}

	return u.UserExtend, nil
}
