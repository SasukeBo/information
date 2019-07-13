package models

import (
	"github.com/astaxie/beego/logs"
	"time"
)

// Enum the User status

// User 用户模型
type User struct {
	Password    string
	Uuid        string       `orm:"pk"`
	Account     string       `orm:"unique"`
	UserProfile *UserProfile `orm:"rel(one)"`
	Role        *Role        `orm:"rel(fk)"`
	Status      BaseStatus   `orm:"default(0)"`
	CreatedAt   time.Time    `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time    `orm:"auto_now;type(datetime)"`
}

// Insert 创建用户
func (a *User) Insert() {
	if n, err := repo.Insert(a); err != nil {
		logs.Error("User insert failed:", err)
	} else {
		logs.Info("User insert success: ", n)
	}
}
