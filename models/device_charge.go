package models

import (
	"fmt"
	"time"

	"github.com/SasukeBo/information/models/errors"
)

// DeviceCharge 设备负责人关系模型
type DeviceCharge struct {
	ID        int       `orm:"auto;pk;column(id)"`
	User      *User     `orm:"rel(fk);on_delete()"` // 设备负责人，用户删除时删除
	Device    *Device   `orm:"rel(fk);on_delete()"` // 设备，删除时删除
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

// TableUnique 自定义唯一键
func (dc *DeviceCharge) TableUnique() [][]string {
	return [][]string{
		[]string{"device_id", "user_id"},
	}
}

// Get get device_charge by id
func (dc *DeviceCharge) Get() error {
	if err := Repo.Read(dc); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "get device_charge error",
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
			Message: "insert device_charge error",
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
			Message: "delete device_charge error",
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
			Message: "update device_charge error",
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
			Message: "device_charge load user error",
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
			Message: "device_charge load device error",
			OriErr:  err,
		}
	}

	return dc.Device, nil
}

// Validate _
func (dc *DeviceCharge) Validate(sign string) error {
	qs := Repo.QueryTable("device_charge_ability").Filter("device_charge__id", dc.ID).Filter("privilege__sign", sign)
	var dca DeviceChargeAbility
	if err := qs.One(&dca); err != nil {
		return errors.LogicError{
			Type:    "Validate",
			Field:   sign,
			Message: fmt.Sprintf("can't access without %s ability", sign),
			OriErr:  err,
		}
	}

	return nil
}
