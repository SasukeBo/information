package models

import (
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

// Get get user by id
func (u *User) Get() error {
	if err := Repo.Read(u); err != nil {
		return utils.ORMError{
			Message: "user get error",
			OrmErr:  err,
		}
	}

	return nil
}

// GetByUUID get user by uuid
func (u *User) GetByUUID() error {
	if err := Repo.Read(u, "uuid"); err != nil {
		return utils.ORMError{
			Message: "user get by uuid error",
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
