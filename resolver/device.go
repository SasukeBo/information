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

// ComputeDeviceOEE 计算设备当班OEE 稼动率 良率 等指标
func ComputeDeviceOEE(params graphql.ResolveParams) (interface{}, error) {
	var device *models.Device
	switch v := params.Source.(type) {
	case *models.Device:
		device = v
	case models.Device:
		device = &v
	}

	response, err := device.GetCurrentClassOEE()
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*                   begin
---------------------------------------------- */

// CountStatusDailyDuration 统计设备每日每状态持续时间
func CountStatusDailyDuration(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	id := params.Args["id"].(int)
	daysCount := params.Args["daysCount"].(int)
	begin, end := timeIntervalBeforeNow(daysCount)

	duration := struct {
		Days    []string
		OffLine []float64
		Prod    []float64
		Stop    []float64
	}{
		generateDays(begin, daysCount),
		make([]float64, daysCount),
		make([]float64, daysCount),
		make([]float64, daysCount),
	}

	query := `
	SELECT dsl.begin_at, dsl.finish_at, status
	FROM device_status_log AS dsl
	WHERE dsl.device_id = ? AND
	(dsl.finish_at > ? OR dsl.begin_at < ?)
	`

	type row struct {
		BeginAt  time.Time
		FinishAt time.Time
		Status   int
	}
	var rows []row
	_, err := o.Raw(query, id, begin, end).QueryRows(&rows)
	if err != nil {
		return nil, models.Error{Message: "获取失败", OriErr: err}
	}

	for _, row := range rows {
		switch row.Status {
		case models.DeviceStatus.Prod:
			hours := utils.CalculateDurations(row.BeginAt, row.FinishAt)
			offset := utils.DaySub(row.BeginAt, begin)
			for i, hour := range hours {
				index := i + offset
				if index >= daysCount {
					break
				} else if index < 0 {
					continue
				}
				duration.Prod[index] += hour
			}

		case models.DeviceStatus.Stop:
			hours := utils.CalculateDurations(row.BeginAt, row.FinishAt)
			offset := utils.DaySub(row.BeginAt, begin)
			for i, hour := range hours {
				index := i + offset
				if index >= daysCount {
					break
				} else if index < 0 {
					continue
				}
				duration.Stop[index] += hour
			}

		case models.DeviceStatus.OffLine:
			hours := utils.CalculateDurations(row.BeginAt, row.FinishAt)
			offset := utils.DaySub(row.BeginAt, begin)
			for i, hour := range hours {
				index := i + offset
				if index >= daysCount {
					break
				} else if index < 0 {
					continue
				}
				duration.OffLine[index] += hour
			}
		}
	}

	return duration, nil
}

// private funcs

// 获取从当前时间开始的前 daysCount 天时间区间
func timeIntervalBeforeNow(daysCount int) (time.Time, time.Time) {
	endTime := time.Now()
	beginTime := endTime.AddDate(0, 0, -daysCount)
	y1, m1, d1 := endTime.Date()
	end := time.Date(y1, m1, d1, 0, 0, 0, 0, time.UTC)
	y2, m2, d2 := beginTime.Date()
	begin := time.Date(y2, m2, d2, 0, 0, 0, 0, time.UTC)

	return begin, end
}

func formatDate(t time.Time) string {
	return fmt.Sprintf("%d/%d/%d", t.Year(), t.Month(), t.Day())
}

func generateDays(begin time.Time, daysCount int) []string {
	days := []string{}
	for i := 0; i < daysCount; i++ {
		days = append(days, formatDate(begin.AddDate(0, 0, i)))
	}
	return days
}

/*                    end
---------------------------------------------- */

// 转换duration为int 单位（秒），注意该方法仅适用于duration不大于一天的情况
func durationToInt(duration string) int {
	times := splitTime(duration)
	s := 0
	if v := times["seconds"]; v != "" {
		seconds, err := strconv.Atoi(v)
		if err == nil {
			s += seconds
		}
	}

	if v := times["minutes"]; v != "" {
		minutes, err := strconv.Atoi(v)
		if err == nil {
			s += minutes * 60
		}
	}

	if v := times["hours"]; v != "" {
		hours, err := strconv.Atoi(v)
		if err == nil {
			s += hours * 60 * 60
		}
	}

	return s
}

// GetDetectItemChartInitData 获取设备检测项图表初始化数据
func GetDetectItemChartInitData(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	var (
		deviceID  = params.Args["deviceID"]
		productID = params.Args["productID"]
		limit     = params.Args["limit"]
	)

	var detectItems []*models.DetectItem
	if _, err := o.QueryTable("DetectItem").Filter("product_id", productID).All(&detectItems); err != nil {
		return nil, models.Error{Message: "Get product detect_items failed.", OriErr: err}
	}

	detectItemValuesSQL := `
	SELECT div.value FROM detect_item_value div
	JOIN product_ins pi ON div.product_ins_id = pi.id
	JOIN device_product_ship dps ON pi.device_product_ship_id = dps.id
	WHERE dps.device_id = ? AND dps.product_id = ? AND div.detect_item_id = ? ORDER BY pi.created_at DESC LIMIT ?;
	`
	for _, di := range detectItems {
		var values []*models.DetectItemValue
		if _, err := o.Raw(detectItemValuesSQL, deviceID, productID, di.ID, limit).QueryRows(&values); err != nil {
			return nil, models.Error{Message: "get detect_item_values failed.", OriErr: err}
		}
		di.Values = values
	}

	productCreateTimestampSQL := `
	SELECT created_at FROM product_ins pi
	JOIN device_product_ship dps ON pi.device_product_ship_id = dps.id
	WHERE dps.device_id = ? AND dps.product_id = ? ORDER BY created_at DESC LIMIT ?;
	`
	var times []time.Time
	if _, err := o.Raw(productCreateTimestampSQL, deviceID, productID, limit).QueryRows(&times); err != nil {
		return nil, models.Error{Message: "get product_ins create timestamp failed.", OriErr: err}
	}

	return struct {
		DetectItems []*models.DetectItem
		Timestamps  []time.Time
	}{detectItems, times}, nil
}

// MonthlyAnalyzeDeviceFormatTime 格式化时间长度
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
	SELECT SUM(dsl.finish_at - dsl.begin_at) AS duration
	FROM device_status_log AS dsl
	WHERE status = ? AND finish_at > TIMESTAMP'0001-01-01 00:00:00+00' And device_id = ?
	`
	var prodDuration []orm.Params
	_, err = o.Raw(sql, models.DeviceStatus.Prod, device.ID).Values(&prodDuration)
	if err != nil {
		return nil, models.Error{Message: "analyze error.", OriErr: err}
	}
	p, ok := prodDuration[0]["duration"].(string)
	if ok {
		response.RunningTime = p
	}

	var stopDuration []orm.Params
	_, err = o.Raw(sql, models.DeviceStatus.Stop, device.ID).Values(&stopDuration)
	if err != nil {
		return nil, models.Error{Message: "analyze error.", OriErr: err}
	}
	s, ok := stopDuration[0]["duration"].(string)
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

	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	if !user.HasAccess("device_c", models.PrivType.Default) {
		return nil, models.Error{Message: "user can't create device without device_c ability."}
	}

	deviceType := params.Args["type"].(string)
	if deviceType == "" {
		return nil, models.Error{Message: "device type can't be blank"}
	}

	deviceName := params.Args["name"].(string)
	if deviceName == "" {
		return nil, models.Error{Message: "device name can't be blank"}
	}

	prodSpeed := params.Args["prodSpeed"].(float64)
	productID := params.Args["productID"].(int)
	privateForms := params.Args["privateForms"].([]interface{})
	count := 0

	if err := o.Begin(); err != nil {
		return nil, models.Error{Message: "begin transaction failed.", OriErr: err}
	}

	for _, item := range privateForms {
		privateForm, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		device := &models.Device{
			Type:      deviceType,
			Name:      deviceName,
			ProdSpeed: prodSpeed,
			Address:   privateForm["address"].(string),
			Number:    privateForm["number"].(string),
			Token:     utils.GenRandomToken(8),
			User:      &user,
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

	if value, ok := params.Args["name"].(string); ok && value != "" {
		device.Name = value
	}

	if value, ok := params.Args["type"].(string); ok && value != "" {
		device.Type = value
	}

	if value := params.Args["prodSpeed"]; value != nil {
		device.ProdSpeed = value.(float64)
	}

	if value := params.Args["address"]; value != nil {
		device.Address = value.(string)
	}

	if value := params.Args["number"]; value != nil {
		device.Number = value.(string)
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
	case models.DeviceStatusLog:
		return v.LoadDevice()
	case *models.DeviceStatusLog:
		return v.LoadDevice()
	case models.StopReason:
		return v.LoadDevice()
	default:
		return nil, models.Error{Message: "load related device failed."}
	}
}

// CountDeviceStatus _
func CountDeviceStatus(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	var sql string
	filter := params.Args["filter"].(string)

	switch filter {
	case "all":
		sql = `SELECT COUNT(status), status FROM device GROUP BY status`
	case "register":
		user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
		sql = fmt.Sprintf("SELECT COUNT(status), status FROM device WHERE user_id = %d GROUP BY status", user.ID)
	default:
		return nil, nil
	}

	var results []*struct {
		Count  int
		Status int
	}

	if _, err := o.Raw(sql).QueryRows(&results); err != nil {
		return nil, models.Error{Message: "count device by status failed.", OriErr: err}
	}

	var response struct {
		Prod    int
		Stop    int
		Offline int
	}

	for _, r := range results {
		switch r.Status {
		case models.DeviceStatus.Prod:
			response.Prod = r.Count
		case models.DeviceStatus.Stop:
			response.Stop = r.Count
		case models.DeviceStatus.OffLine:
			response.Offline = r.Count
		}
	}

	return response, nil
}

// private functions

func splitTime(dbTimeDuration string) map[string]string {
	dbDurationPattern := `^(\d*)( days? )?(\d+):(\d{2}):(\d{2})(\.\d*)?$`
	reg := regexp.MustCompile(dbDurationPattern)
	matches := reg.FindStringSubmatch(dbTimeDuration)
	var (
		days    string
		hours   string
		minutes string
		seconds string
	)

	if len(matches) > 1 {
		days = matches[1]
	}

	if len(matches) > 3 {
		hours = matches[3]
	}

	if len(matches) > 4 {
		minutes = matches[4]
	}

	if len(matches) > 5 {
		seconds = matches[5]
	}

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

	var (
		seconds int
		minutes int
		hours   int
		days    int
	)

	if v, e := strconv.Atoi(times["seconds"]); e == nil {
		seconds = v
	}

	if v, e := strconv.Atoi(times["minutes"]); e == nil {
		minutes = v
	}

	if v, e := strconv.Atoi(times["days"]); e == nil {
		days = v
	}

	if v, e := strconv.Atoi(times["hours"]); e == nil {
		if v >= 24 {
			hours = v % 24
		}

		days += v / 24
	}

	format = strings.Replace(format, "%S", strconv.FormatInt(int64(seconds), 10), 1)
	format = strings.Replace(format, "%M", strconv.FormatInt(int64(minutes), 10), 1)
	format = strings.Replace(format, "%H", strconv.FormatInt(int64(hours), 10), 1)
	format = strings.Replace(format, "%D", strconv.FormatInt(int64(days), 10), 1)

	return format
}
