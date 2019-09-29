package models

import (
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

// Insert _
func (p *Product) Insert() error {
	if _, err := Repo.Insert(p); err != nil {
		return LogicError{
			Type:    "Model",
			Message: "create product failed!",
			OriErr:  err,
		}
	}

	return nil
}

// Delete _
func (p *Product) Delete() error {
	if _, err := Repo.Delete(p); err != nil {
		return LogicError{
			Type:    "Model",
			Message: "delete product failed!",
			OriErr:  err,
		}
	}

	return nil
}

// GetBy _
func (p *Product) GetBy(col string) error {
	if err := Repo.Read(p, col); err != nil {
		return LogicError{
			Type:    "Model",
			Message: "get product failed",
			OriErr:  err,
		}
	}

	return nil
}

// LoadDetectItem _
func (p *Product) LoadDetectItem() ([]*DetectItem, error) {
	if _, err := Repo.LoadRelated(p, "DetectItems"); err != nil {
		return nil, LogicError{
			Type:    "Model",
			Message: "Load related detectItems failed.",
			OriErr:  err,
		}
	}

	return p.DetectItems, nil
}
