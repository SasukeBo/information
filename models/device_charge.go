package models

import (
	"time"

	"github.com/SasukeBo/information/utils"
)

// DeviceCharge 设备负责人关系模型
type DeviceCharge struct {
	ID        int       `orm:"auto;pk;column(id)"`
	User      *User     `orm:"rel(fk);on_delete()"` // 设备负责人，用户删除时删除
	Device    *Device   `orm:"rel(fk);on_delete()"` // 设备，删除时删除
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

// Get get device_charge by id
func (dc *DeviceCharge) Get() error {
	if err := Repo.Read(dc); err != nil {
		return utils.ORMError{
			Message: "device_charge get error",
			OrmErr:  err,
		}
	}

	return nil
}

// LoadUser _
func (dc *DeviceCharge) LoadUser() (*User, error) {
	if _, err := Repo.LoadRelated(dc, "User"); err != nil {
		return nil, utils.ORMError{
			Message: "devcice_charge load related user error",
			OrmErr:  err,
		}
	}

	return dc.User, nil
}

// LoadDevice _
func (dc *DeviceCharge) LoadDevice() (*Device, error) {
	if _, err := Repo.LoadRelated(dc, "Device"); err != nil {
		return nil, utils.ORMError{
			Message: "devcice_charge load related device error",
			OrmErr:  err,
		}
	}

	return dc.Device, nil
}
