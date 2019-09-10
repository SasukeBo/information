package models

import (
	"github.com/SasukeBo/information/models/errors"
	"time"
)

// DeviceStatus 设备状态
var DeviceStatus = struct {
	OffLine int // 离线
	Prod    int // 运行
	Stop    int // 在线但是未运行，就是停机
}{0, 1, 2}

// DeviceStatusLog 设备运行状态变更模型
type DeviceStatusLog struct {
	Status   int       // 设备运行状态
	ID       int       `orm:"auto;pk;column(id)"`
	Device   *Device   `orm:"rel(fk);on_delete()"`
	ChangeAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// Insert _
func (dsl *DeviceStatusLog) Insert() error {
	remoteIP := dsl.Device.RemoteIP
	if _, err := Repo.Insert(dsl); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "device_status_log insert error",
			OriErr:  err,
		}
	}

	device, err := dsl.LoadDevice()
	if err != nil {
		return err
	}

	device.Status = dsl.Status
	device.RemoteIP = remoteIP
	if err := device.Update("status", "remote_ip"); err != nil {
		return err
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
