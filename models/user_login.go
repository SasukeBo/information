package models

import (
	"time"

	"github.com/SasukeBo/information/utils"
)

// UserLogin 用户登录模型
// 用户请求到达服务器时，
type UserLogin struct {
	EncryptedPasswd string    // 加密后的密码
	UserAgent       string    // 用户代理
	ID              int       `orm:"auto;pk;column(id)"`
	User            *User     `orm:"rel(fk);on_delete()"` // 用户删除时删除
	RemoteIP        string    `orm:"column(remote_ip)"`   // 登录IP
	SessionID       string    `orm:"column(session_id)"`  // session id
	Remembered      bool      `orm:"default(true)"`       // 记住登录
	Logout          bool      `orm:"default(false)"`      // 是否登出
	CreatedAt       time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt       time.Time `orm:"auto_now;type(datetime)"`
}

// TableUnique 自定义唯一键
func (ul *UserLogin) TableUnique() [][]string {
	return [][]string{
		[]string{"user_id", "session_id"},
	}
}

// LoadUser _
func (ul *UserLogin) LoadUser() (*User, error) {
	if _, err := Repo.LoadRelated(ul, "User"); err != nil {
		return nil, utils.ORMError{
			Message: "user_login load related user error",
			OrmErr:  err,
		}
	}

	return ul.User, nil
}
