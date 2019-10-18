package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

var accErr = Error{Message: "only device register can make this operation!"}

// Device 设备模型
type Device struct {
	ID             int              `orm:"auto;pk;column(id)"` // PKey 主键
	Type           string           // 类型
	Name           string           // 设备名称
	Address        string           `orm:"null"`                             // 设备地址
	Number         string           `orm:"null"`                             // 设备编号
	RemoteIP       string           `orm:"null;column(remote_ip)"`           // 接入IP
	Token          string           `orm:"unique;index"`                     // 设备Token，用于数据加密
	Status         int              `orm:"default(0)"`                       // 离线状态
	User           *User            `orm:"rel(fk);null;on_delete(set_null)"` // 注册人
	DeviceChargers []*DeviceCharger `orm:"reverse(many)"`
	CreatedAt      time.Time        `orm:"auto_now_add;type(datetime)"`
	StatusChangeAt time.Time        `orm:"auto_now;type(datetime)"`
	UpdatedAt      time.Time        `orm:"auto_now;type(datetime)"`
}

// LoadUser _
func (d *Device) LoadUser() (*User, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(d, "User"); err != nil {
		return nil, Error{Message: "load related user failed.", OriErr: err}
	}

	return d.User, nil
}

// LoadDeviceCharge _
func (d *Device) LoadDeviceCharge() ([]*DeviceCharger, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(d, "DeviceChargers"); err != nil {
		return nil, Error{Message: "load related device_chargers failed.", OriErr: err}
	}

	return d.DeviceChargers, nil
}

// ValidateAccess _
func (d *Device) ValidateAccess(u *User) error {
	if d.User.ID != u.ID {
		return accErr
	}

	return nil
}
