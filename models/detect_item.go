package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// DetectItem 检测项
type DetectItem struct {
	ID         int       `orm:"auto;pk;column(id)"`
	Sign       string    // 检测项标识
	Product    *Product  `orm:"rel(fk);on_delete()"`
	UpperLimit float64   `orm:"null"`
	LowerLimit float64   `orm:"null"`
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `orm:"auto_now;type(datetime)"`
}

// TableUnique _
func (di *DetectItem) TableUnique() [][]string {
	return [][]string{
		[]string{"sign", "product_id"},
	}
}

// LoadProduct _
func (di *DetectItem) LoadProduct() (*Product, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(di, "Product"); err != nil {
		return nil, Error{Message: "load related product failed.", OriErr: err}
	}

	return di.Product, nil
}
