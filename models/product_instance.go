package models

import (
	"time"
)

// ProductIns 产品实例
type ProductIns struct {
	ID                int                `orm:"auto;pk;column(id)"`
	DeviceProductShip *DeviceProductShip `orm:"rel(fk);on_delete()"` // 产品设备关系删除时，会无法定位到该产品实例，将会被删除
	DetectItemValues  []*DetectItemValue `orm:"reverse(many)"`       // 产品实例的检测项值，对应参数表格的一行
	CreatedAt         time.Time          `orm:"type(datetime)"`
}
