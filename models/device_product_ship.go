package models

// DeviceProductShip 设备与产品关系
type DeviceProductShip struct {
	ID        int           `orm:"auto;pk;column(id)"`
	Device    *Device       `orm:"null;rel(fk);on_delete(set_null)"`
	Product   *Product      `orm:"rel(fk);on_delete()"` // 如果产品删除了 则删除设备产品关系
	Instances []*ProductIns `orm:"reverse(many)"`
}

// TableUnique _
func (dps *DeviceProductShip) TableUnique() [][]string {
	return [][]string{
		[]string{"device_id", "product_id"},
	}
}
