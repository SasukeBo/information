package models

import (
	// "time"

	"github.com/SasukeBo/information/models/errors"
)

// DeviceChargeAbility 设备负责人权限模型
type DeviceChargeAbility struct {
	ID           int           `orm:"auto;pk;column(id)"`
	DeviceCharge *DeviceCharge `orm:"rel(fk);on_delete()"` // 设备负责关系删除时删除
	Privilege    *Privilege    `orm:"rel(fk);on_delete()"` // 权限删除时删除
}

// TableUnique 自定义唯一键
func (dca *DeviceChargeAbility) TableUnique() [][]string {
	return [][]string{
		[]string{"device_charge_id", "privilege_id"},
	}
}

// LoadDeviceCharge related load device_charge
func (dca *DeviceChargeAbility) LoadDeviceCharge() (*DeviceCharge, error) {
	if _, err := Repo.LoadRelated(dca, "DeviceCharge"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Field:   "DeviceChargeAbility",
			Message: "LoadDeviceCharge() error",
			OriErr:  err,
		}
	}

	return dca.DeviceCharge, nil
}

// LoadPrivilege related load privilege
func (dca *DeviceChargeAbility) LoadPrivilege() (*Privilege, error) {
	if _, err := Repo.LoadRelated(dca, "Privilege"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Field:   "DeviceChargeAbility",
			Message: "LoadPrivilege() error",
			OriErr:  err,
		}
	}

	return dca.Privilege, nil
}
