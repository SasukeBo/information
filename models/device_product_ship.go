package models

// DeviceProductShip 设备与产品关系
type DeviceProductShip struct {
	ID      int      `orm:"auto;pk;column(id)"`
	Device  *Device  `orm:"rel(fk)"`
	Product *Product `orm:"rel(fk)"`
}

// TableUnique _
func (dps *DeviceProductShip) TableUnique() [][]string {
	return [][]string{
		[]string{"device_id", "product_id"},
	}
}
