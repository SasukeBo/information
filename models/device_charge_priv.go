package models

import (
// "time"

// "github.com/SasukeBo/information/utils"
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
