package models

import (
	"fmt"
	"time"

	"github.com/SasukeBo/information/models/errors"
)

// DeviceCharger 设备负责人
type DeviceCharger struct {
	ID         int `orm:"auto;pk;column(id)"`
	Name       string
	Phone      string
	Department string
	JobNumber  string
	Device     *Device   `orm:"rel(fk);on_delete()"` // 设备，删除时删除
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `orm:"auto_now;type(datetime)"`
}

// TableUnique 自定义唯一键
func (dc *DeviceCharger) TableUnique() [][]string {
	return [][]string{
		[]string{"device_id", "name", "phone"},
	}
}

// GetBy get device_charger by col
func (dc *DeviceCharger) GetBy(col string) error {
	if err := Repo.Read(dc, col); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   col,
			Message: fmt.Sprintf("get device_charger by %v error", col),
			OriErr:  err,
		}
	}

	return nil
}

// Insert _
func (dc *DeviceCharger) Insert() error {
	if _, err := Repo.Insert(dc); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "insert device_charger error",
			OriErr:  err,
		}
	}

	return nil
}

// Delete _
func (dc *DeviceCharger) Delete() error {
	if _, err := Repo.Delete(dc); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "delete device_charger error",
			OriErr:  err,
		}
	}

	return nil
}

// Update _
func (dc *DeviceCharger) Update(cols ...string) error {
	if _, err := Repo.Update(dc, cols...); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "update device_charger error",
			OriErr:  err,
		}
	}

	return nil
}

// LoadDevice _
func (dc *DeviceCharger) LoadDevice() (*Device, error) {
	if _, err := Repo.LoadRelated(dc, "Device"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "device_charger load device error",
			OriErr:  err,
		}
	}

	return dc.Device, nil
}
