package models

import (
	"time"

	"github.com/SasukeBo/information/models/errors"
)

// DeviceParam 设备参数模型
type DeviceParam struct {
	Name      string    // 参数名称
	Sign      string    // 参数签名（标识），要求英文及数字组合的字符串
	Type      string    // 参数值类型，string？int？bool？
	Device    *Device   `orm:"rel(fk);on_delete()"`
	ID        int       `orm:"auto;pk;column(id)"`
	Author    *User     `orm:"rel(fk);null;on_delete(set_null)"` // 创建人，删除时置空
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// Get get device_param by id
func (dp *DeviceParam) Get() error {
	if err := Repo.Read(dp); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "DeviceParam",
			Message: "Get() error",
			OriErr:  err,
		}
	}

	return nil
}

// Delete delete device_param by id
func (dp *DeviceParam) Delete() error {
	if err := dp.Get(); err != nil {
		return err
	}
	if _, err := Repo.Delete(dp); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "DeviceParam",
			Message: "Delete() error",
			OriErr:  err,
		}
	}

	return nil
}

// Insert insert device_param
func (dp *DeviceParam) Insert() error {
	if _, err := Repo.Insert(dp); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "DeviceParam",
			Message: "Insert() error",
			OriErr:  err,
		}
	}

	return nil
}

// Update update device_param
func (dp *DeviceParam) Update(cols ...string) error {
	if _, err := Repo.Update(dp, cols...); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "DeviceParam",
			Message: "Update() error",
			OriErr:  err,
		}
	}

	return nil
}

// LoadAuthor _
func (dp *DeviceParam) LoadAuthor() (*User, error) {
	if _, err := Repo.LoadRelated(dp, "Author"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Field:   "DeviceParam",
			Message: "LoadAuthor() error",
			OriErr:  err,
		}
	}

	return dp.Author, nil
}

// LoadDevice _
func (dp *DeviceParam) LoadDevice() (*Device, error) {
	if _, err := Repo.LoadRelated(dp, "Device"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Field:   "DeviceParam",
			Message: "LoadDevice() error",
			OriErr:  err,
		}
	}

	return dp.Device, nil
}
