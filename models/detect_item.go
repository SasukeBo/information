package models

import (
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

// GetBy _
func (di *DetectItem) GetBy(col string) error {
	if err := Repo.Read(di, col); err != nil {
		return LogicError{
			Type:    "Model",
			Field:   col,
			Message: "get detect_item failed.",
			OriErr:  err,
		}
	}

	return nil
}

// Insert _
func (di *DetectItem) Insert() error {
	if _, err := Repo.Insert(di); err != nil {
		return LogicError{
			Type:    "Model",
			Message: "Insert detect_item failed.",
			OriErr:  err,
		}
	}

	return nil
}

// LoadProduct _
func (di *DetectItem) LoadProduct() (*Product, error) {
	if _, err := Repo.LoadRelated(di, "Product"); err != nil {
		return nil, LogicError{
			Type:    "Model",
			Message: "Load related product failed.",
			OriErr:  err,
		}
	}

	return di.Product, nil
}

// Update _
func (di *DetectItem) Update(cols ...string) error {
	if _, err := Repo.Update(di, cols...); err != nil {
		return LogicError{
			Type:    "Model",
			Message: "Update detect_item failed.",
			OriErr:  err,
		}
	}

	return nil
}

// Delete _
func (di *DetectItem) Delete() error {
	if _, err := Repo.Delete(di); err != nil {
		return LogicError{
			Type:    "Model",
			Message: "Delete detect_item failed.",
			OriErr:  err,
		}
	}

	return nil
}
