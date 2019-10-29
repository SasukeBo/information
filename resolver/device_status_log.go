package resolver

import (
	"fmt"
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
	"time"
	// "strings"
)

// CountDeviceStopType 统计设备停机类型数量
func CountDeviceStopType(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()

	deviceID := params.Args["deviceID"]
	beginTime := utils.TruncateDay(params.Args["beginTime"].(time.Time))
	endTime := utils.TruncateDay(params.Args["endTime"].(time.Time))

	sql := `
	SELECT COUNT(rt.name),rt.name, DATE_TRUNC('day', dsl.begin_at) AS date
	FROM device_status_log AS dsl
	JOIN device_stop_reason AS dsr ON dsl.code = dsr.code
	JOIN reason_type AS rt ON dsr.type_id = rt.id
	WHERE dsl.status = 2 AND dsl.device_id = ? AND dsl.begin_at > ? AND dsl.begin_at < ?
	GROUP BY rt.name, date
	`

	var rows []struct {
		Count int
		Name  string
		Date  time.Time
	}

	if _, err := o.Raw(sql, deviceID, beginTime, endTime).QueryRows(&rows); err != nil {
		return nil, models.Error{Message: "count device stop type failed.", OriErr: err}
	}

	// 先获取日期列表
	dayCount := utils.DaySub(endTime, beginTime)
	dates := generateDays(beginTime, dayCount)

	datas := make(map[string][]int)
	// 计算日期与起始日期差值，放入数据
	for _, row := range rows {
		name := row.Name
		counts := datas[name]
		if counts == nil {
			counts = make([]int, dayCount)
		}
		fmt.Println(row.Date)
		index := utils.DaySub(row.Date, beginTime)
		counts[index] = row.Count
		datas[name] = counts
	}

	types := make([]string, 0)
	type typeCount struct {
		Name    string
		Numbers []int
	}
	typeCounts := make([]typeCount, 0)

	for k, v := range datas {
		types = append(types, k)
		typeCounts = append(typeCounts, typeCount{k, v})
	}

	return struct {
		Counts []typeCount
		Days   []string
		Types  []string
	}{typeCounts, dates, types}, nil
}

// ListDeviceStopLogs 设备停机日志列表查询
func ListDeviceStopLogs(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	deviceID := params.Args["deviceID"]
	qs := o.QueryTable("DeviceStatusLog").Filter("device_id", deviceID).Filter("status", models.DeviceStatus.Stop).OrderBy("-begin_at")

	if v := params.Args["beginTime"]; v != nil {
		qs = qs.Filter("begin_at__gt", v)
	}

	if v := params.Args["endTime"]; v != nil {
		qs = qs.Filter("begin_at__lt", v)
	}

	total, err := qs.Count()
	if err != nil {
		return nil, models.Error{Message: "get device_stop_log total count failed.", OriErr: err}
	}

	if v := params.Args["limit"]; v != nil {
		qs = qs.Limit(v)
	}

	if v := params.Args["offset"]; v != nil {
		qs = qs.Offset(v)
	}

	var logs []*models.DeviceStatusLog
	if _, err := qs.All(&logs); err != nil {
		return nil, models.Error{Message: "get device_stop_log failed.", OriErr: err}
	}

	return struct {
		Total int64
		Logs  []*models.DeviceStatusLog
	}{total, logs}, nil
}

// RefreshDeviceStatus _
func RefreshDeviceStatus(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	device := models.Device{ID: params.Args["deviceID"].(int)}
	if err := o.Read(&device, "id"); err != nil {
		return nil, models.Error{Message: "get device failed.", OriErr: err}
	}

	return device, nil
}
