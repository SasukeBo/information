package models

import (
	"time"
)

// ProductIns 产品实例
type ProductIns struct {
	ID               int                `orm:"auto;pk;column(id)"`
	Product          *Product           `orm:"rel(fk);on_delete()"` // 产品类
	DetectItemValues []*DetectItemValue `orm:"reverse(many)"`       // 产品实例的检测项值，对应参数表格的一行
	CreatedAt        time.Time          `orm:"auto_now_add;type(datetime)"`
}
