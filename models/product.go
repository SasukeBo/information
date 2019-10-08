package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Product 产品模型
type Product struct {
	ID          int           `orm:"auto;pk;column(id)"`
	Name        string        `orm:"unique"`        // 产品名称
	DetectItems []*DetectItem `orm:"reverse(many)"` // 产品检测项
	CreatedAt   time.Time     `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time     `orm:"auto_now;type(datetime)"`
}

// LoadDetectItem _
func (p *Product) LoadDetectItem() ([]*DetectItem, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(p, "DetectItems"); err != nil {
		return nil, Error{Message: "load related detect_items failed.", OriErr: err}
	}

	return p.DetectItems, nil
}
