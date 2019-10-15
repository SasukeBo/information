package resolver

import (
	"fmt"
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
)

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
