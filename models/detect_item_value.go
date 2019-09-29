package models

// DetectItemValue 测量值
type DetectItemValue struct {
	ID         int         `orm:"auto;pk;column(id)"`
	DetectItem *DetectItem `orm:"rel(fk);on_delete()"` // 所属检测项
	ProductIns *ProductIns `orm:"rel(fk);on_delete()"` // 所属产品实例
	Value      float64     // 检测值
}

// Insert _
func (div *DetectItemValue) Insert() error {
	if _, err := Repo.Insert(div); err != nil {
		return LogicError{
			Type:    "Model",
			Message: "Insert detect_item_value failed.",
			OriErr:  err,
		}
	}

	return nil
}
