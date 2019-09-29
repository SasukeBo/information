package models

import (
	"fmt"
	"time"

	"github.com/SasukeBo/information/models/errors"
)

var accErr = errors.LogicError{
	Type:    "Resolvers",
	Message: "only device register can make this operation!",
}

// Device 设备模型
type Device struct {
	Type           string           // 类型
	Name           string           // 设备名称
	Address        string           `orm:"null"` // 设备地址
	Number         string           `orm:"null"` // 设备编号
	RemoteIP       string           `orm:"null;column(remote_ip)"`
	Token          string           `orm:"unique;index"`                     // 设备Token，用于数据加密
	Status         int              `orm:"default(0)"`                       // 离线状态
	ID             int              `orm:"auto;pk;column(id)"`               // PKey 主键
	User           *User            `orm:"rel(fk);null;on_delete(set_null)"` // 注册人
	DeviceChargers []*DeviceCharger `orm:"reverse(many)"`
	Description    string           `orm:"null"` // 描述
	CreatedAt      time.Time        `orm:"auto_now_add;type(datetime)"`
	StatusChangeAt time.Time        `orm:"auto_now;type(datetime)"`
	UpdatedAt      time.Time        `orm:"auto_now;type(datetime)"`
}

// GetBy get device by col
func (d *Device) GetBy(col string) error {
	if err := Repo.Read(d, col); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   col,
			Message: fmt.Sprintf("get device by %s error", col),
			OriErr:  err,
		}
	}

	return nil
}

// Insert _
func (d *Device) Insert() error {
	if _, err := Repo.Insert(d); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "insert device error",
			OriErr:  err,
		}
	}

	return nil
}

// Update device with cols
func (d *Device) Update(cols ...string) error {
	if _, err := Repo.Update(d, cols...); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "update device error",
			OriErr:  err,
		}
	}

	return nil
}

// Delete _
func (d *Device) Delete() error {
	if _, err := Repo.Delete(d); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "uuid",
			Message: "delete device by uuid error",
			OriErr:  err,
		}
	}

	return nil
}

// LoadUser _
func (d *Device) LoadUser() (*User, error) {
	if _, err := Repo.LoadRelated(d, "User"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "device load user error",
			OriErr:  err,
		}
	}

	return d.User, nil
}

// LoadDeviceCharge _
func (d *Device) LoadDeviceCharge() ([]*DeviceCharger, error) {
	if _, err := Repo.LoadRelated(d, "DeviceChargers"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "device load device_chargers error",
			OriErr:  err,
		}
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
