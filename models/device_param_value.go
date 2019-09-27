package models

import (
	"fmt"
	"github.com/SasukeBo/information/models/errors"
	"time"
)

// DeviceParamValue 设备参数值模型
type DeviceParamValue struct {
	Value       string       // 参数值
	ID          int          `orm:"auto;pk;column(id)"`
	DeviceParam *DeviceParam `orm:"rel(fk);on_delete()"`
	CreatedAt   time.Time    `orm:"type(datetime)"`
}

// GetBy _
func (dpv *DeviceParamValue) GetBy(col string) error {
	if err := Repo.Read(dpv, col); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   col,
			OriErr:  err,
			Message: fmt.Sprintf("device_param_value get by %s error", col),
		}
	}

	return nil
}

// Insert _
func (dpv *DeviceParamValue) Insert() error {
	if _, err := Repo.Insert(dpv); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "device_param_value insert error",
			OriErr:  err,
		}
	}

	return nil
}

// LoadDeviceParam related load device_param
func (dpv *DeviceParamValue) LoadDeviceParam() (*DeviceParam, error) {
	if _, err := Repo.LoadRelated(dpv, "DeviceParam"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "device_param_value load device_param error",
			OriErr:  err,
		}
	}

	return dpv.DeviceParam, nil
}
