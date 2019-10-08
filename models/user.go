package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// User 用户模型
type User struct {
	ID        int    `orm:"auto;pk;column(id)"`
	Name      string // 真实姓名
	Phone     string `orm:"unique"`
	Password  string
	AvatarURL string    `orm:"column(avatar_url);null"` // 头像链接
	Email     string    `orm:"null"`
	Role      *Role     `orm:"rel(fk);null;on_delete(set_null)"` // 用户角色，删除时置空
	Status    int       `orm:"default(0)"`                       // 基础状态
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

// LoadRole related load role
func (u *User) LoadRole() (*Role, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(u, "Role"); err != nil {
		return nil, Error{Message: "load related role failed.", OriErr: err}
	}

	return u.Role, nil
}
