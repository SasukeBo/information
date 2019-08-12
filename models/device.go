package models

import (
	"fmt"
	"time"

	"github.com/SasukeBo/information/errors"
)

// Device 设备模型
type Device struct {
	Type        string    // 类型
	Name        string    // 设备名称
	Mac         string    // 设备Mac地址
	Token       string    `orm:"unique;index"`                     // 设备Token，用于数据加密
	Status      int       `orm:"default(0)"`                       // 基础状态
	ID          int       `orm:"auto;pk;column(id)"`               // PKey 主键
	UUID        string    `orm:"column(uuid);unique;index"`        // 通用唯一标识符
	User        *User     `orm:"rel(fk);null;on_delete(set_null)"` // 注册人
	Description string    `orm:"null"`                             // 描述
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`
}

// GetBy get device by col
func (d *Device) GetBy(col string) error {
	if err := Repo.Read(d, col); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "Device",
			Message: fmt.Sprintf("GetBy(%s) error", col),
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
			Field:   "Device",
			Message: "Insert() error",
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
			Field:   "Device",
			Message: "Update() error",
			OriErr:  err,
		}
	}

	return nil
}

// DeleteByUUID _
func (d *Device) DeleteByUUID() error {
	if err := d.GetBy("uuid"); err != nil {
		return err
	}

	if _, err := Repo.Delete(d); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "Device",
			Message: "DeleteByUUID() error",
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
			Field:   "Device",
			Message: "LoadUser() error",
			OriErr:  err,
		}
	}

	return d.User, nil
}
