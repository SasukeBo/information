package models

// DetectItemValue 测量值
type DetectItemValue struct {
	ID         int         `orm:"auto;pk;column(id)"`
	DetectItem *DetectItem `orm:"rel(fk);on_delete()"` // 所属检测项
	ProductIns *ProductIns `orm:"rel(fk);on_delete()"` // 所属产品实例
	Value      float32     // 检测值
}
