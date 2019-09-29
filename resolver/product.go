package resolver

import (
	// "fmt"
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
)

// LoadProduct _
func LoadProduct(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.DetectItem:
		return v.LoadProduct()
	case *models.DetectItem:
		return v.LoadProduct()
	default:
		return nil, models.LogicError{
			Type:    "Resolver",
			Message: "load product failed.",
		}
	}
}

// CreateProduct 创建产品信息
func CreateProduct(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		o.Rollback()
		return nil, models.LogicError{Type: "Model", Message: "begin transaction failed.", OriErr: err}
	}

	name := params.Args["name"].(string)
	product := models.Product{Name: name}

	if _, err := o.Insert(&product); err != nil {
		o.Rollback()
		return nil, models.LogicError{Type: "Model", Message: "Insert product failed.", OriErr: err}
	}

	detectItems := params.Args["detectItems"].([]interface{})
	for _, item := range detectItems {
		var ok bool
		value := item.(map[string]interface{})
		sign, ok := value["sign"].(string)
		if !ok {
			o.Rollback()
			return nil, models.LogicError{Type: "Resolver", Field: "Sign", Message: "value invalid."}
		}

		upperLimit, ok := value["upperLimit"].(float64)
		if !ok {
			o.Rollback()
			return nil, models.LogicError{Type: "Resolver", Field: "UpperLimit", Message: "value invalid."}
		}

		lowerLimit, ok := value["lowerLimit"].(float64)
		if !ok {
			o.Rollback()
			return nil, models.LogicError{Type: "Resolver", Field: "LowerLimit", Message: "value invalid."}
		}

		detectItem := &models.DetectItem{
			Sign:       sign,
			UpperLimit: upperLimit,
			LowerLimit: lowerLimit,
			Product:    &product,
		}
		if _, err := o.Insert(detectItem); err != nil {
			o.Rollback()
			return nil, models.LogicError{Type: "Model", Message: "Insert product failed.", OriErr: err}
		}
	}

	o.Commit()
	return product, nil
}

// DeleteProduct _
func DeleteProduct(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	product := &models.Product{ID: id}
	if err := product.Delete(); err != nil {
		return nil, err
	}

	return id, nil
}

// GetProduct _
func GetProduct(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	product := &models.Product{ID: id}
	if err := product.GetBy("id"); err != nil {
		return nil, err
	}

	return product, nil
}

// ListProduct _
func ListProduct(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		o.Rollback()
		return nil, models.LogicError{Type: "Model", Message: "begin transaction failed.", OriErr: err}
	}

	qs := o.QueryTable("product").OrderBy("-created_at")
	if namePattern := params.Args["namePattern"]; namePattern != nil {
		qs = qs.Filter("name__icontains", namePattern)
	}

	cnt, err := qs.Count()
	if err != nil {
		o.Rollback()
		return nil, models.LogicError{Type: "Model", Message: "count product list failed."}
	}

	if limit := params.Args["limit"]; limit != nil {
		qs = qs.Limit(limit)
	}

	if offset := params.Args["offset"]; offset != nil {
		qs = qs.Offset(offset)
	}

	var products []*models.Product
	if _, err := qs.All(&products); err != nil {
		o.Rollback()
		return nil, models.LogicError{Type: "Model", Message: "get list of product failed."}
	}

	o.Commit()

	return struct {
		Count    int64
		Products []*models.Product
	}{cnt, products}, nil
}
