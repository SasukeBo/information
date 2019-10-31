package resolver

import (
	"fmt"
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
	"strconv"
	"strings"
	"time"
)

// GetProductDevices 获取产品的生产设备
func GetProductDevices(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	product := models.Product{ID: id}
	devices, err := product.GetDevices(orm.NewOrm())
	if err != nil {
		return nil, err
	}

	return devices, nil
}

// ProductOverView 产品总览页数据接口
func ProductOverView(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()

	response := struct {
		DeviceTotalCount    int
		DeviceProdCount     int
		InstanceCount       int
		QualifiedCount      int
		TodayInstanceCount  int
		TodayQualifiedCount int
	}{}

	id := params.Args["id"].(int)
	product := models.Product{ID: id}
	if err := o.Read(&product); err != nil {
		return nil, models.Error{Message: "product not found.", OriErr: err}
	}

	sql1 := `
	SELECT COUNT(device.id) AS device_total_count
	FROM device JOIN device_product_ship dps ON device.id = dps.device_id
	WHERE dps.product_id = ?
	`
	o.Raw(sql1, id).QueryRow(&response)
	sql2 := `
	SELECT COUNT(device.id) AS device_prod_count
	FROM device JOIN device_product_ship dps ON device.id = dps.device_id
	WHERE dps.product_id = ? AND device.status = ?
	`
	o.Raw(sql2, id, models.DeviceStatus.Prod).QueryRow(&response)
	sql3 := `
	SELECT COUNT(*) AS instance_count
	FROM product_ins JOIN device_product_ship dps ON product_ins.device_product_ship_id = dps.id
	WHERE dps.product_id = ?
	`
	o.Raw(sql3, id).QueryRow(&response)
	sql4 := `
	SELECT COUNT(*) AS qualified_count
	FROM product_ins JOIN device_product_ship dps ON product_ins.device_product_ship_id = dps.id
	WHERE dps.product_id = ? AND product_ins.qualified = true
	`
	o.Raw(sql4, id).QueryRow(&response)
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	sql5 := `
	SELECT COUNT(*) AS today_instance_count
	FROM product_ins JOIN device_product_ship dps ON product_ins.device_product_ship_id = dps.id
	WHERE dps.product_id = ? AND product_ins.created_at > ?
	`
	o.Raw(sql5, id, midnight).QueryRow(&response)

	sql6 := `
	SELECT COUNT(*) AS today_qualified_count
	FROM product_ins JOIN device_product_ship dps ON product_ins.device_product_ship_id = dps.id
	WHERE dps.product_id = ? AND product_ins.qualified = true AND product_ins.created_at > ?
	`
	o.Raw(sql6, id, midnight).QueryRow(&response)

	return response, nil
}

// ProductHistogram _
func ProductHistogram(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	var (
		lowerTime time.Time
		upperTime time.Time
	)
	groups := 80
	id := params.Args["id"].(int)
	detectItemID := params.Args["detectItemID"].(int)
	deviceID := params.Args["deviceID"]

	now := time.Now()

	if v := params.Args["lowerTime"]; v == nil {
		lowerTime = now.AddDate(0, -1, 0)
	} else {
		lowerTime = v.(time.Time)
	}

	if v := params.Args["upperTime"]; v == nil {
		upperTime = now
	} else {
		upperTime = v.(time.Time)
	}

	// 获取 最小最大值 区间长度
	gql := `
	SELECT MAX(value) AS max, MIN(value) AS min
	FROM detect_item_value AS div
	JOIN product_ins AS di ON div.product_ins_id = di.id
	WHERE div.detect_item_id = ? AND di.created_at > ? AND di.created_at < ?;
	`
	var result struct {
		Max float64
		Min float64
	}
	if err := o.Raw(gql, detectItemID, lowerTime, upperTime).QueryRow(&result); err != nil {
		return nil, models.Error{Message: "calculation error.", OriErr: err}
	}
	length := (result.Max - result.Min) / float64(groups)

	// 生成sql
	sql2 := `
	SELECT %s
	FROM detect_item_value AS div
	JOIN product_ins AS pi ON pi.id = div.product_ins_id
	JOIN device_product_ship AS dps ON dps.id = pi.device_product_ship_id
	WHERE dps.product_id = ? AND div.detect_item_id = ? %s
	`
	selectTPL := `SUM(CASE WHEN %f <= div.value AND div.value < %f THEN 1 ELSE 0 END)`
	selects := []string{}
	xAxisData := []string{}
	for i := 0; i < groups; i++ {
		lower := result.Min + (float64(i) * length)
		upper := result.Min + (float64(i+1) * length)
		selects = append(selects, fmt.Sprintf(selectTPL, lower, upper))
		xAxisData = append(xAxisData, fmt.Sprintf(`%0.3f-%5.3f`, lower, upper))
	}
	selectCond := strings.Join(selects, ",")
	var deviceCond string
	args := []interface{}{id, detectItemID}
	if deviceID == nil {
		deviceCond = ""
	} else {
		args = append(args, deviceID)
		deviceCond = "AND dps.device_id = ?"
	}
	query := fmt.Sprintf(sql2, selectCond, deviceCond)
	var result2 []orm.ParamsList
	if _, err := o.Raw(query, args...).ValuesList(&result2); err != nil {
		return nil, models.Error{Message: "fetch histogram results failed.", OriErr: err}
	}

	seriesData := []int{}
	for _, v := range result2[0] {
		s := v.(string)
		i, _ := strconv.Atoi(s)
		seriesData = append(seriesData, i)
	}

	return struct {
		XAxisData  []string
		SeriesData []int
	}{xAxisData, seriesData}, nil
}

// DeviceLoadProduct _
func DeviceLoadProduct(params graphql.ResolveParams) (interface{}, error) {
	device := params.Source.(models.Device)

	o := orm.NewOrm()

	var err error
	var ship models.DeviceProductShip
	err = o.QueryTable("DeviceProductShip").OrderBy("-id").Filter("device_id", device.ID).Limit(1).One(&ship)
	if err != nil {
		logs.Error(err)
		return nil, nil
	}

	_, err = o.LoadRelated(&ship, "product")
	if err != nil {
		logs.Error(err)
		return nil, nil
	}

	return ship.Product, nil
}

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
	product.Token = utils.GenRandomToken(4)

	if total := params.Args["total"]; total != nil {
		product.Total = total.(int)
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

	if orderNum := params.Args["orderNum"]; orderNum != nil {
		product.OrderNum = orderNum.(string)
	}

	if customer := params.Args["customer"]; customer != nil {
		product.Customer = customer.(string)
	}

	if customerContact := params.Args["customerContact"]; customerContact != nil {
		product.CustomerContact = customerContact.(string)
	}

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
		detectItem := models.DetectItem{Sign: sign, Product: &product}

		if upperLimit := value["upperLimit"]; upperLimit != nil {
			value, ok := upperLimit.(float64)
			if !ok {
				o.Rollback()
				return nil, models.Error{Message: "invalid upper_limit."}
			}
			detectItem.UpperLimit = value
		}

		if lowerLimit := value["lowerLimit"]; lowerLimit != nil {
			value, ok := lowerLimit.(float64)
			if !ok {
				o.Rollback()
				return nil, models.Error{Message: "invalid lower_limit."}
			}
			detectItem.LowerLimit = value
		}

		if _, err := o.Insert(&detectItem); err != nil {
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
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)

	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		o.Rollback()
		return nil, models.Error{Message: "begin transaction failed.", OriErr: err}
	}

	qs := o.QueryTable("product").OrderBy("created_at")
	if namePattern := params.Args["namePattern"]; namePattern != nil {
		qs = qs.Filter("name__icontains", namePattern)
	}

	if self := params.Args["self"]; self != nil {
		if self.(bool) {
			qs = qs.Filter("register_id", user.ID)
		}
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
