package models

import (
	"github.com/SasukeBo/information/models/errors"
	"time"
)

// DeviceStatus 设备状态
var DeviceStatus = struct {
	Prod    int // 运行
	Stop    int // 停机
	OffLine int // 离线
	OnLine  int // 在线
}{0, 1, 2, 3}

// DeviceStatusLog 设备运行状态变更模型
type DeviceStatusLog struct {
	Status   int       // 设备运行状态
	ID       int       `orm:"auto;pk;column(id)"`
	Device   *Device   `orm:"rel(fk);on_delete()"`
	ChangeAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// Insert _
func (dsl *DeviceStatusLog) Insert() error {
	if _, err := Repo.Insert(dsl); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "device_status_log insert error",
			OriErr:  err,
		}
	}

	return nil
}

// LoadDevice _
func (dsl *DeviceStatusLog) LoadDevice() (*Device, error) {
	if _, err := Repo.LoadRelated(dsl, "Device"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "device_status_log load device error",
			OriErr:  err,
		}
	}

	return dsl.Device, nil
}
