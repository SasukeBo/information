package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// ProductIns 产品实例
type ProductIns struct {
	ID                int                `orm:"auto;pk;column(id)"`
	Qualified         bool               `orm:"default(false)"`
	DeviceProductShip *DeviceProductShip `orm:"rel(fk);on_delete()"` // 产品设备关系删除时，会无法定位到该产品实例，将会被删除
	DetectItemValues  []*DetectItemValue `orm:"reverse(many)"`       // 产品实例的检测项值，对应参数表格的一行
	CreatedAt         time.Time          `orm:"type(datetime)"`
}

// LoadDetectItemValues _
func (pi *ProductIns) LoadDetectItemValues() ([]*DetectItemValue, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(pi, "DetectItemValues"); err != nil {
		return nil, nil
	}

	return pi.DetectItemValues, nil
}

// Get get product_ins by id
func (pi *ProductIns) Get(o orm.Ormer) error {
	if err := o.Read(pi); err != nil {
		return Error{Message: "Get product_ins by id failed.", OriErr: err}
	}
	return nil
}
