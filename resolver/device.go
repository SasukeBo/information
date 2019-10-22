package resolver

import (
	"fmt"
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type statistics struct {
	RunningTime string
	Activation  float64
	YieldRate   float64
	Yield       int
}

// GetRealTimeStatistics 获取设备实时数据
func GetRealTimeStatistics(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()

	deviceID := params.Args["deviceID"].(int)
	productID := params.Args["productID"].(int)
	floatPrecision := params.Args["floatPrecision"].(int)
	afterTime := params.Args["afterTime"].(time.Time)
	limit := params.Args["limit"].(int)

	dps := models.DeviceProductShip{Device: &models.Device{ID: deviceID}, Product: &models.Product{ID: productID}}
	if err := o.Read(&dps, "device_id", "product_id"); err != nil {
		return nil, models.Error{Message: "product not producing on this device", OriErr: err}
	}

	var signs []string
	sql1 := `SELECT sign FROM detect_item where detect_item.product_id = ?`
	if _, err := o.Raw(sql1, productID).QueryRows(&signs); err != nil {
		return nil, models.Error{Message: "get product detect_items failed.", OriErr: err}
	}

	type seriesData struct {
		Sign  string
		Value float64
		Time  time.Time
	}
	seriesDatas := []seriesData{}

	sql2 := `
	select pi.created_at as time, round(div.value::numeric, ?) as value
	from detect_item_value as div
	join product_ins as pi on div.product_ins_id = pi.id
	join device_product_ship as dps on pi.device_product_ship_id = dps.id
	join detect_item as di on di.id = div.detect_item_id
	where dps.id = ? and di.sign = ? and pi.created_at > ? order by time desc limit ?;
	`
	for _, s := range signs {
		var data seriesData
		if err := o.Raw(sql2, floatPrecision, dps.ID, s, afterTime, limit).QueryRow(&data); err != nil {
			return nil, models.Error{Message: "get seriesData failed.", OriErr: err}
		}
		data.Sign = s
		seriesDatas = append(seriesDatas, data)
	}

	return seriesDatas, nil
}

// MonthlyAnalyzeDeviceFormatTime 分析设备月数据
func MonthlyAnalyzeDeviceFormatTime(params graphql.ResolveParams) (interface{}, error) {
	format := "%D days %H hours %M minutes"
	if value := params.Args["format"]; value != nil {
		format = value.(string)
	}

	response := params.Source.(statistics)
	return formatTime(response.RunningTime, format), nil
}

// MonthlyAnalyzeDevice 分析设备月数据
func MonthlyAnalyzeDevice(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	var err error
	response := statistics{}

	id := params.Args["id"].(int)

	device := models.Device{ID: id}
	if err := o.Read(&device, "id"); err != nil {
		return nil, models.Error{Message: "get device failed.", OriErr: err}
	}

	var sql = `
	SELECT SUM(duration) FROM (
		SELECT (dsl.finish_at - dsl.begin_at) AS duration
		FROM device_status_log AS dsl
		WHERE status = ? AND finish_at > TIMESTAMP'0001-01-01 00:00:00+00' And device_id = ?
	) AS duration_list ;
	`
	var prodDuration []orm.Params
	_, err = o.Raw(sql, models.DeviceStatus.Prod, device.ID).Values(&prodDuration)
	if err != nil {
		return nil, models.Error{Message: "analyze error.", OriErr: err}
	}
	p, ok := prodDuration[0]["sum"].(string)
	if ok {
		response.RunningTime = p
	}

	var stopDuration []orm.Params
	_, err = o.Raw(sql, models.DeviceStatus.Stop, device.ID).Values(&stopDuration)
	if err != nil {
		return nil, models.Error{Message: "analyze error.", OriErr: err}
	}
	s, ok := stopDuration[0]["sum"].(string)
	if ok {
		pd, errPD := parseTimeDurationFromDB(p)
		sd, errSD := parseTimeDurationFromDB(s)
		if errPD == nil && errSD == nil {
			response.Activation = float64(pd.Seconds() / (pd.Seconds() + sd.Seconds()))
		}
	}

	var insSQL = `
	SELECT COUNT(*) FROM product_ins AS pi WHERE pi.device_product_ship_id = (
		SELECT id FROM device_product_ship WHERE device_id = ? ORDER BY id DESC LIMIT 1
	) AND pi.qualified = ? AND pi.created_at > (SELECT CURRENT_TIMESTAMP - '30 day'::INTERVAL);
	`

	var qualifiedCount []orm.Params
	var unqualifiedCount []orm.Params
	_, err = o.Raw(insSQL, device.ID, true).Values(&qualifiedCount)
	if err != nil {
		return nil, models.Error{Message: "count qualified products failed."}
	}
	_, err = o.Raw(insSQL, device.ID, false).Values(&unqualifiedCount)
	if err != nil {
		return nil, models.Error{Message: "count unqualified products failed."}
	}

	qcStr := qualifiedCount[0]["count"].(string)
	ucStr := unqualifiedCount[0]["count"].(string)
	qc, _ := strconv.Atoi(qcStr)
	uc, _ := strconv.Atoi(ucStr)

	response.Yield = qc + uc
	if response.Yield != 0 {
		response.YieldRate = float64(qc) / float64(response.Yield)
	}

	return response, nil
}

// GetDevice 获取设备
func GetDevice(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	id := params.Args["id"].(int)

	device := models.Device{ID: id}
	if err := o.Read(&device, "id"); err != nil {
		return nil, models.Error{Message: "get device failed.", OriErr: err}
	}

	return device, nil
}

// GetDeviceByToken 获取设备
func GetDeviceByToken(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	token := params.Args["token"].(string)

	device := models.Device{Token: token}
	if err := o.Read(&device, "token"); err != nil {
		return nil, models.Error{Message: "get device failed.", OriErr: err}
	}

	return device, nil
}

// ListDevice 获取设备列表
func ListDevice(params graphql.ResolveParams) (interface{}, error) {
	cond := orm.NewCondition()
	o := orm.NewOrm()

	if pattern := params.Args["search"]; pattern != nil {
		subCond := orm.NewCondition()
		cond = cond.AndCond(subCond.Or("type__icontains", pattern).Or("name__icontains", pattern).Or("address__icontains", pattern).Or("number__icontains", pattern))
	}

	if status := params.Args["status"]; status != nil {
		cond = cond.And("status", status)
	}

	if isRegister := params.Args["self"]; isRegister != nil {
		if isRegister.(bool) {
			user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
			cond = cond.And("user_id", user.ID)
		}
	}
	qs := o.QueryTable("device").SetCond(cond).OrderBy("created_at")

	cnt, err := qs.Count()
	if err != nil {
		return nil, err
	}

	if limit := params.Args["limit"]; limit != nil {
		qs = qs.Limit(limit)
	}

	if offset := params.Args["offset"]; offset != nil {
		qs = qs.Offset(offset)
	}

	var devices []*models.Device
	if _, err := qs.All(&devices); err != nil {
		return nil, models.Error{Message: "list device failed.", OriErr: err}
	}

	return struct {
		Total   int64
		Devices []*models.Device
	}{cnt, devices}, nil
}

// CreateDevice 创建设备
func CreateDevice(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		return nil, models.Error{Message: "begin transaction failed.", OriErr: err}
	}

	// 验证用户是否有创建设备的权限
	if err := utils.ValidateAccess(&params, "device_c", models.PrivType.Default); err != nil {
		o.Rollback()
		return nil, err
	}
	rootValue := params.Info.RootValue.(map[string]interface{})
	user := rootValue["currentUser"].(models.User)

	deviceType := params.Args["type"].(string)
	if err := utils.ValidateStringEmpty(deviceType, "type"); err != nil {
		o.Rollback()
		return nil, err
	}

	deviceName := params.Args["name"].(string)
	if err := utils.ValidateStringEmpty(deviceName, "name"); err != nil {
		o.Rollback()
		return nil, err
	}

	productID := params.Args["productID"].(int)
	privateForms := params.Args["privateForms"].([]interface{})
	count := 0
	for _, item := range privateForms {
		privateForm, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		device := &models.Device{
			Type:    deviceType,
			Name:    deviceName,
			Address: privateForm["address"].(string),
			Number:  privateForm["number"].(string),
			Token:   utils.GenRandomToken(8),
			User:    &user,
		}

		if _, err := o.Insert(device); err != nil {
			continue
		}

		ship := &models.DeviceProductShip{
			Device:  device,
			Product: &models.Product{ID: productID},
		}
		o.Insert(ship)

		count++
	}

	if count == 0 {
		o.Rollback()
		return nil, models.Error{Message: "device create failed."}
	}

	o.Commit()
	return count, nil
}

// UpdateDevice 更新设备
func UpdateDevice(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	device := models.Device{ID: params.Args["id"].(int)}
	if err := o.Read(&device, "id"); err != nil {
		return nil, models.Error{Message: "get device failed.", OriErr: err}
	}

	if err := device.ValidateAccess(&user); err != nil {
		return nil, err
	}

	if value := params.Args["address"]; value != nil {
		device.Address = value.(string)
	}

	if value := params.Args["name"]; value != nil {
		if err := utils.ValidateStringEmpty(value.(string), "name"); err != nil {
			return nil, err
		}
		device.Name = value.(string)
	}

	if value := params.Args["number"]; value != nil {
		device.Number = value.(string)
	}

	if value := params.Args["type"]; value != nil {
		if err := utils.ValidateStringEmpty(value.(string), "type"); err != nil {
			return nil, err
		}
		device.Type = value.(string)
	}

	if _, err := o.Update(&device); err != nil {
		return nil, models.Error{Message: "update device failed.", OriErr: err}
	}

	return device, nil
}

// DeleteDevice 更新设备
func DeleteDevice(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	device := models.Device{ID: params.Args["id"].(int)}
	if err := o.Read(&device, "id"); err != nil {
		return nil, models.Error{Message: "get device failed.", OriErr: err}
	}

	if err := device.ValidateAccess(&user); err != nil {
		return nil, err
	}

	if _, err := o.Delete(&device); err != nil {
		return nil, models.Error{Message: "delete device failed.", OriErr: err}
	}

	return device.ID, nil
}

// LoadDevice _
func LoadDevice(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.DeviceCharger:
		return v.LoadDevice()
	case *models.DeviceCharger:
		return v.LoadDevice()
	case models.DeviceStatusLog:
		return v.LoadDevice()
	case *models.DeviceStatusLog:
		return v.LoadDevice()
	default:
		return nil, models.Error{Message: "load related device_charger failed."}
	}
}

// CountDeviceStatus _
func CountDeviceStatus(params graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}

// private functions

func splitTime(dbTimeDuration string) map[string]string {
	dbDurationPattern := `^(\d*)( day )?(\d{2}):(\d{2}):(\d{2})(\.\d*)?$`
	reg := regexp.MustCompile(dbDurationPattern)
	matches := reg.FindStringSubmatch(dbTimeDuration)
	days := matches[1]
	hours := matches[3]
	minutes := matches[4]
	seconds := matches[5]

	return map[string]string{
		"days":    days,
		"hours":   hours,
		"minutes": minutes,
		"seconds": seconds,
	}
}

func parseTimeDurationFromDB(dbTimeDuration string) (time.Duration, error) {
	times := splitTime(dbTimeDuration)
	hours, err := strconv.Atoi(times["hours"])
	if err != nil {
		hours = 0
	}

	if day, err := strconv.Atoi(times["days"]); err == nil {
		hours = day*24 + hours
	}

	duration, err := time.ParseDuration(fmt.Sprintf(
		"%vh%vm%vs",
		hours,
		times["minutes"],
		times["seconds"],
	))

	if err != nil {
		return time.Duration(0), err
	}

	return duration, nil
}

func formatTime(dbTimeDuration, format string) string {
	times := splitTime(dbTimeDuration)

	if v, e := strconv.Atoi(times["days"]); e == nil {
		format = strings.Replace(format, "%D", strconv.FormatInt(int64(v), 10), 1)
	} else {
		format = strings.Replace(format, "%D", "0", 1)
	}

	if v, e := strconv.Atoi(times["hours"]); e == nil {
		format = strings.Replace(format, "%H", strconv.FormatInt(int64(v), 10), 1)
	} else {
		format = strings.Replace(format, "%H", "0", 1)
	}

	if v, e := strconv.Atoi(times["minutes"]); e == nil {
		format = strings.Replace(format, "%M", strconv.FormatInt(int64(v), 10), 1)
	} else {
		format = strings.Replace(format, "%M", "0", 1)
	}

	if v, e := strconv.Atoi(times["seconds"]); e == nil {
		format = strings.Replace(format, "%S", strconv.FormatInt(int64(v), 10), 1)
	} else {
		format = strings.Replace(format, "%S", "0", 1)
	}

	return format
}
