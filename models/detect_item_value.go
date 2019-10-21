package models

import (
	"github.com/astaxie/beego/orm"
)

// DetectItemValue 测量值
type DetectItemValue struct {
	ID         int         `orm:"auto;pk;column(id)"`
	DetectItem *DetectItem `orm:"rel(fk);on_delete()"` // 所属检测项
	ProductIns *ProductIns `orm:"rel(fk);on_delete()"` // 所属产品实例
	Value      float64     // 检测值
}

// LoadDetectItem _
func (div *DetectItemValue) LoadDetectItem() (*DetectItem, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(div, "DetectItem"); err != nil {
		return nil, nil
	}

	return div.DetectItem, nil
}
