package resolver

import (
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
	"time"
)

// LoadProduct _
func LoadProduct(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.DetectItem:
		return v.LoadProduct()
	case *models.DetectItem:
		return v.LoadProduct()
	default:
		return nil, models.Error{Message: "load related product failed."}
	}
}

// CreateProduct 创建产品信息
func CreateProduct(params graphql.ResolveParams) (interface{}, error) {
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		o.Rollback()
		return nil, models.Error{Message: "begin transaction failed.", OriErr: err}
	}

	name := params.Args["name"].(string)
	product := models.Product{Name: name, Register: &user}

	if _, err := o.Insert(&product); err != nil {
		o.Rollback()
		return nil, models.Error{Message: "insert product failed.", OriErr: err}
	}

	detectItems := params.Args["detectItems"].([]interface{})
	for _, item := range detectItems {
		var ok bool
		value := item.(map[string]interface{})
		sign, ok := value["sign"].(string)
		if !ok {
			o.Rollback()
			return nil, models.Error{Message: "invalid sign."}
		}

		upperLimit, ok := value["upperLimit"].(float64)
		if !ok {
			o.Rollback()
			return nil, models.Error{Message: "invalid upper_limit."}
		}

		lowerLimit, ok := value["lowerLimit"].(float64)
		if !ok {
			o.Rollback()
			return nil, models.Error{Message: "invalid lower_limit."}
		}

		detectItem := &models.DetectItem{
			Sign:       sign,
			UpperLimit: upperLimit,
			LowerLimit: lowerLimit,
			Product:    &product,
		}

		if _, err := o.Insert(detectItem); err != nil {
			o.Rollback()
			return nil, models.Error{Message: "insert detect_item failed.", OriErr: err}
		}
	}

	o.Commit()
	return product, nil
}

// DeleteProduct _
func DeleteProduct(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	id := params.Args["id"].(int)

	product := &models.Product{ID: id}
	if _, err := o.Delete(product); err != nil {
		return nil, models.Error{Message: "delete product failed.", OriErr: err}
	}

	return id, nil
}

// GetProduct _
func GetProduct(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	id := params.Args["id"].(int)

	product := &models.Product{ID: id}
	if err := o.Read(product, "id"); err != nil {
		return nil, models.Error{Message: "get product failed.", OriErr: err}
	}

	return product, nil
}

// ListProduct _
func ListProduct(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		o.Rollback()
		return nil, models.Error{Message: "begin transaction failed.", OriErr: err}
	}

	qs := o.QueryTable("product").OrderBy("-created_at")
	if namePattern := params.Args["namePattern"]; namePattern != nil {
		qs = qs.Filter("name__icontains", namePattern)
	}

	cnt, err := qs.Count()
	if err != nil {
		o.Rollback()
		return nil, models.Error{Message: "count product failed.", OriErr: err}
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
		return nil, models.Error{Message: "list product failed.", OriErr: err}
	}

	o.Commit()

	return struct {
		Count    int64
		Products []*models.Product
	}{cnt, products}, nil
}

// UpdateProduct _
func UpdateProduct(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	id := params.Args["id"].(int)
	product := models.Product{ID: id}
	if err := o.Read(&product); err != nil {
		return nil, models.Error{Message: "read product failed.", OriErr: err}
	}

	if name := params.Args["name"]; name != nil {
		product.Name = name.(string)
	}

	if productorContact := params.Args["productorContact"]; productorContact != nil {
		product.ProductorContact = productorContact.(string)
	}

	if productor := params.Args["productor"]; productor != nil {
		product.Productor = productor.(string)
	}

	if finishTime := params.Args["finishTime"]; finishTime != nil {
		product.FinishTime = finishTime.(time.Time)
	}

	if total := params.Args["total"]; total != nil {
		product.Total = total.(int)
	}

	if orderNum := params.Args["orderNum"]; orderNum != nil {
		product.OrderNum = orderNum.(string)
	}

	if customer := params.Args["customer"]; customer != nil {
		product.Customer = customer.(string)
	}

	if customerContact := params.Args["customerContact"]; customerContact != nil {
		product.CustomerContact = customerContact.(string)
	}

	if _, err := o.Update(&product); err != nil {
		return nil, models.Error{Message: "update product failed.", OriErr: err}
	}

	return product, nil
}
