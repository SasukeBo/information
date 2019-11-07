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
	case *models.DetectItemValue:
		return v.LoadDetectItem()
	default:
		return nil, models.Error{Message: "load related detect_item failed."}
	}
}

// CreateDetectItem _
func CreateDetectItem(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	productID := params.Args["productID"].(int)
	product := models.Product{ID: productID}
	if err := o.Read(&product, "id"); err != nil {
		return nil, models.Error{Message: "get product failed.", OriErr: err}
	}

	sign := params.Args["sign"].(string)
	detectItem := models.DetectItem{Sign: sign, Product: &product}

	if upperLimit := params.Args["upperLimit"]; upperLimit != nil {
		detectItem.UpperLimit = upperLimit.(float64)
	}

	if lowerLimit := params.Args["lowerLimit"]; lowerLimit != nil {
		detectItem.LowerLimit = lowerLimit.(float64)
	}

	if _, err := o.Insert(&detectItem); err != nil {
		return nil, models.Error{Message: "insert detect_item failed.", OriErr: err}
	}

	return detectItem, nil
}

// UpdateDetectItem _
func UpdateDetectItem(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	id := params.Args["id"].(int)
	detectItem := &models.DetectItem{ID: id}
	if err := o.Read(detectItem, "id"); err != nil {
		return nil, models.Error{Message: "get detect_item failed.", OriErr: err}
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

	if _, err := o.Update(detectItem, "sign", "upper_limit", "lower_limit"); err != nil {
		return nil, models.Error{Message: "update detect_item failed.", OriErr: err}
	}

	return detectItem, nil
}

// DeleteDetectItem _
func DeleteDetectItem(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	id := params.Args["id"].(int)
	detectItem := models.DetectItem{ID: id}

	if _, err := o.Delete(&detectItem); err != nil {
		return nil, models.Error{Message: "delete detect_item failed.", OriErr: err}
	}

	return id, nil
}

// GetDetectItem _
func GetDetectItem(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	id := params.Args["id"].(int)
	detectItem := models.DetectItem{ID: id}

	if err := o.Read(&detectItem, "id"); err != nil {
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
		return nil, models.Error{Message: "count detect_item failed", OriErr: err}
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
		return nil, models.Error{Message: "list detect_item failed.", OriErr: err}
	}

	return struct {
		Count       int64
		DetectItems []*models.DetectItem
	}{cnt, detectItems}, nil
}

// ProductDetectItemsCount 项目检测项数统计
func ProductDetectItemsCount(params graphql.ResolveParams) (interface{}, error) {
	product, ok := params.Source.(*models.Product)
	if !ok {
		return nil, nil
	}

	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("SELECT COUNT(*) FROM detect_item AS di WHERE di.product_id = ?", product.ID)
	var res []orm.Params
	_, err := r.Values(&res)
	if err != nil {
		return nil, nil
	}

	return res[0]["count"], nil
}
