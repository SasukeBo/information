package models

import (
	"github.com/SasukeBo/information/models/errors"
	"time"
)

// DeviceParamValue 设备参数值模型
type DeviceParamValue struct {
	Value       string       // 参数值
	ID          int          `orm:"auto;pk;column(id)"`
	DeviceParam *DeviceParam `orm:"rel(fk);on_delete()"`
	CreatedAt   time.Time    `orm:"auto_now_add;type(datetime)"`
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
