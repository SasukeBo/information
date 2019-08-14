package models

import (
	"fmt"
	"github.com/graphql-go/graphql"

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

// GetBy _
func (dca *DeviceChargeAbility) GetBy(col string) error {
	if err := Repo.Read(dca, col); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   col,
			Message: fmt.Sprintf("get device_charge_ability by %s error", col),
			OriErr:  err,
		}
	}

	return nil
}

// Delete _
func (dca *DeviceChargeAbility) Delete() error {
	if _, err := Repo.Delete(dca); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "device_charge_ability delete error",
			OriErr:  err,
		}
	}

	return nil
}

// LoadDeviceCharge related load device_charge
func (dca *DeviceChargeAbility) LoadDeviceCharge() (*DeviceCharge, error) {
	if _, err := Repo.LoadRelated(dca, "DeviceCharge"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "device_charge_ability load device_charge error",
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
			Message: "device_charge_ability load privilege error",
			OriErr:  err,
		}
	}

	return dca.Privilege, nil
}

// ValidateAccess _
func (dca *DeviceChargeAbility) ValidateAccess(params graphql.ResolveParams, sign ...string) error {
	var deviceCharge *DeviceCharge
	var err error
	if deviceCharge, err = dca.LoadDeviceCharge(); err != nil {
		return err
	}

	if err := deviceCharge.ValidateAccess(params, sign...); err != nil {
		return err
	}

	return nil
}
