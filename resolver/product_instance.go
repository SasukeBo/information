package resolver

import (
	"fmt"
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
)

// ProductInsLoadDetectItemValues _
func ProductInsLoadDetectItemValues(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case *models.ProductIns:
		return v.LoadDetectItemValues()
	case models.ProductIns:
		return v.LoadDetectItemValues()
	}
	return nil, nil
}

// LoadProductIns _
func LoadProductIns(params graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}

// CurrentProductInsCount 统计当前产品产量
func CurrentProductInsCount(params graphql.ResolveParams) (interface{}, error) {
	product, ok := params.Source.(*models.Product)
	if !ok {
		return nil, nil
	}

	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("SELECT COUNT(*) FROM product_ins AS pi WHERE pi.device_product_ship_id IN (SELECT id FROM device_product_ship AS dps WHERE dps.product_id = ?)", product.ID)
	var res []orm.Params
	_, err := r.Values(&res)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	return res[0]["count"], nil
}

// GetProductIns _
func GetProductIns(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	var id = params.Args["id"].(int)
	instance := models.ProductIns{ID: id}
	if err := instance.Get(o); err != nil {
		return nil, err
	}

	return instance, nil
}
