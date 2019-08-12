package models

import (
	"time"

	"github.com/SasukeBo/information/errors"
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
		return errors.LogicError{
			Type:    "Model",
			Field:   "DeviceCharge",
			Message: "Get() error",
			OriErr:  err,
		}
	}

	return nil
}

// Insert _
func (dc *DeviceCharge) Insert() error {
	if _, err := Repo.Insert(dc); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "DeviceCharge",
			Message: "Insert() error",
			OriErr:  err,
		}
	}

	return nil
}

// Delete _
func (dc *DeviceCharge) Delete() error {
	if err := dc.Get(); err != nil {
		return err
	}

	if _, err := Repo.Delete(dc); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "DeviceCharge",
			Message: "Delete() error",
			OriErr:  err,
		}
	}

	return nil
}

// Update _
func (dc *DeviceCharge) Update(cols ...string) error {
	if _, err := Repo.Update(dc, cols...); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "DeviceCharge",
			Message: "Update() error",
			OriErr:  err,
		}
	}

	return nil
}

// LoadUser _
func (dc *DeviceCharge) LoadUser() (*User, error) {
	if _, err := Repo.LoadRelated(dc, "User"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Field:   "DeviceCharge",
			Message: "LoadUser() error",
			OriErr:  err,
		}
	}

	return dc.User, nil
}

// LoadDevice _
func (dc *DeviceCharge) LoadDevice() (*Device, error) {
	if _, err := Repo.LoadRelated(dc, "Device"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Field:   "DeviceCharge",
			Message: "LoadDevice() error",
			OriErr:  err,
		}
	}

	return dc.Device, nil
}
