package models

import (
	"fmt"
	"time"

	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models/errors"
)

// Device 设备模型
type Device struct {
	Type          string          // 类型
	Name          string          // 设备名称
	Mac           string          // 设备Mac地址
	Token         string          `orm:"unique;index"`                     // 设备Token，用于数据加密
	Status        int             `orm:"default(2)"`                       // 基础状态
	ID            int             `orm:"auto;pk;column(id)"`               // PKey 主键
	UUID          string          `orm:"column(uuid);unique;index"`        // 通用唯一标识符
	User          *User           `orm:"rel(fk);null;on_delete(set_null)"` // 注册人
	DeviceCharges []*DeviceCharge `orm:"reverse(many)"`
	Description   string          `orm:"null"` // 描述
	CreatedAt     time.Time       `orm:"auto_now_add;type(datetime)"`
	UpdatedAt     time.Time       `orm:"auto_now;type(datetime)"`
}

// GetBy get device by col
func (d *Device) GetBy(col string) error {
	if err := Repo.Read(d, col); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   col,
			Message: fmt.Sprintf("get device by %s error", col),
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
			Message: "insert device error",
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
			Message: "update device error",
			OriErr:  err,
		}
	}

	return nil
}

// Delete _
func (d *Device) Delete() error {
	if _, err := Repo.Delete(d); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   "uuid",
			Message: "delete device by uuid error",
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
			Message: "device load user error",
			OriErr:  err,
		}
	}

	return d.User, nil
}

// LoadDeviceCharge _
func (d *Device) LoadDeviceCharge() ([]*DeviceCharge, error) {
	if _, err := Repo.LoadRelated(d, "DeviceCharges"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "device load device_charges error",
			OriErr:  err,
		}
	}

	return d.DeviceCharges, nil
}

// LoadDeviceParams _
func (d *Device) LoadDeviceParams() ([]*DeviceParam, error) {
	qs := Repo.QueryTable("device_param").Filter("device_id", d.ID)

	var dps []*DeviceParam
	if _, err := qs.All(&dps); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "device load device_params error",
			OriErr:  err,
		}
	}

	return dps, nil
}

// ValidateAccess validate access to device
func (d *Device) ValidateAccess(params graphql.ResolveParams, args ...string) error {
	currentUser := params.Info.RootValue.(map[string]interface{})["currentUser"].(User)

	if d.User.ID == currentUser.ID {
		return nil
	}

	var dc DeviceCharge
	qs := Repo.QueryTable("device_charge").Filter("device_id", d.ID).Filter("user_id", currentUser.ID)

	// 没有指派
	if err := qs.One(&dc); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "user not charge on this device",
			OriErr:  err,
		}
	}

	if len(args) == 0 {
		return nil
	}

	// charge 权限验证
	if err := dc.Validate(args[0]); err != nil {
		return err
	}

	// 权限通过
	return nil
}
