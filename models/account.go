package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// Account is account model
type Account struct {
	ID   int `orm:"auto;pk"`
	Name string
}

// Insert a account to database
func (a *Account) Insert() {
	o := orm.NewOrm()
	if n, err := o.Insert(a); err != nil {
		logs.Error("User insert failed:", err)
	} else {
		logs.Info("User insert success: ", n)
	}
}
