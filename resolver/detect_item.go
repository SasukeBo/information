package resolver

import (
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
)

// LoadDetectItem _
func LoadDetectItem(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.Product:
		return v.LoadDetectItem()
	case *models.Product:
		return v.LoadDetectItem()
	default:
		return nil, models.LogicError{
			Type:    "Resolver",
			Message: "load detectItem failed.",
		}
	}
}

// CreateDetectItem _
func CreateDetectItem(params graphql.ResolveParams) (interface{}, error) {
	productID := params.Args["productID"].(int)
	product := models.Product{ID: productID}
	if err := product.GetBy("id"); err != nil {
		return nil, err
	}

	sign := params.Args["sign"].(string)
	detectItem := models.DetectItem{Sign: sign, Product: &product}

	if upperLimit := params.Args["upperLimit"]; upperLimit != nil {
		detectItem.UpperLimit = upperLimit.(float64)
	}

	if lowerLimit := params.Args["lowerLimit"]; lowerLimit != nil {
		detectItem.LowerLimit = lowerLimit.(float64)
	}

	if err := detectItem.Insert(); err != nil {
		return nil, err
	}

	return detectItem, nil
}

// UpdateDetectItem _
func UpdateDetectItem(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	detectItem := &models.DetectItem{ID: id}
	if err := detectItem.GetBy("id"); err != nil {
		return nil, err
	}

	if value := params.Args["sign"]; value != nil {
		sign := value.(string)
		if len(sign) > 0 {
			detectItem.Sign = sign
		}
	}

	if upperLimit := params.Args["upperLimit"]; upperLimit != nil {
		detectItem.UpperLimit = upperLimit.(float64)
	}

	if lowerLimit := params.Args["lowerLimit"]; lowerLimit != nil {
		detectItem.LowerLimit = lowerLimit.(float64)
	}

	if err := detectItem.Update("sign", "upper_limit", "lower_limit"); err != nil {
		return nil, err
	}

	return detectItem, nil
}

// DeleteDetectItem _
func DeleteDetectItem(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	detectItem := models.DetectItem{ID: id}

	if err := detectItem.Delete(); err != nil {
		return nil, err
	}

	return id, nil
}

// GetDetectItem _
func GetDetectItem(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	detectItem := models.DetectItem{ID: id}

	if err := detectItem.GetBy("id"); err != nil {
		return nil, err
	}

	return detectItem, nil
}

// ListDetectItem _
func ListDetectItem(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	o.Begin()
	productID := params.Args["productID"]
	qs := o.QueryTable("detect_item").Filter("product_id", productID).OrderBy("-created_at")

	cnt, err := qs.Count()
	if err != nil {
		o.Rollback()
		return nil, models.LogicError{
			Type:    "Model",
			Message: "Count detect_item failed.",
			OriErr:  err,
		}
	}

	if limit := params.Args["limit"]; limit != nil {
		qs = qs.Limit(limit)
	}

	if offset := params.Args["offset"]; offset != nil {
		qs = qs.Offset(offset)
	}

	var detectItems []*models.DetectItem
	if _, err := qs.All(&detectItems); err != nil {
		o.Rollback()
		return nil, models.LogicError{
			Type:    "Model",
			Message: "Get list of detect_items failed.",
			OriErr:  err,
		}
	}

	return struct {
		Count       int64
		DetectItems []*models.DetectItem
	}{cnt, detectItems}, nil
}
