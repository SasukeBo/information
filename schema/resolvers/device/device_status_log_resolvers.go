package device

import (
	"github.com/graphql-go/graphql"
	"strings"
	"time"

	"github.com/SasukeBo/information/models"
	// "github.com/astaxie/beego/logs"
)

// StatusLogList 设备参数创建
func StatusLogList(params graphql.ResolveParams) (interface{}, error) {
	device := models.Device{UUID: params.Args["deviceUUID"].(string)}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	// 验证权限
	if accessErr := device.ValidateAccess(params); accessErr != nil {
		return nil, accessErr
	}

	qs := models.Repo.QueryTable("device_status_log").Filter("device_id", device.ID).OrderBy("created_at")

	if status := params.Args["status"]; status != nil {
		qs = qs.Filter("status", status)
	}

	if beforeTime := params.Args["beforeTime"]; beforeTime != nil {
		qs = qs.Filter("created_at__lt", beforeTime)
	}

	if afterTime := params.Args["afterTime"]; afterTime != nil {
		qs = qs.Filter("created_at__gt", afterTime)
	}

	if limit := params.Args["limit"]; limit != nil {
		qs = qs.Limit(limit)
	}

	var statusLogs []*models.DeviceStatusLog
	if _, err := qs.All(&statusLogs); err != nil {
		return nil, err
	}

	currentDuration := time.Now().Sub(device.StatusChangeAt).Truncate(time.Second)
	currentStatus := &models.DeviceStatusLog{
		Status:   device.Status,
		Duration: int(currentDuration / 1e9),
	}

	statusLogs = append(statusLogs, currentStatus)

	return statusLogs, nil
}

// StatusRefresh _
func StatusRefresh(params graphql.ResolveParams) (interface{}, error) {
	device := models.Device{ID: params.Args["deviceID"].(int)}
	if err := device.GetBy("id"); err != nil {
		return nil, err
	}

	return device, nil
}

// StatusDuration _
func StatusDuration(params graphql.ResolveParams) (interface{}, error) {
	deviceID := params.Args["deviceID"].(int)
	device := models.Device{ID: deviceID}
	if err := device.GetBy("id"); err != nil {
		return nil, err
	}

	status := params.Args["status"].(int)

	qs := models.Repo.QueryTable("device_status_log").Filter("device_id", device.ID).Filter("status", status)

	var statusLogs []*models.DeviceStatusLog
	if _, err := qs.All(&statusLogs); err != nil {
		return nil, err
	}

	var durations time.Duration
	for _, sl := range statusLogs {
		durations += time.Duration(sl.Duration * 1e9)
	}

	if device.Status == status {
		currentDuration := time.Now().Sub(device.StatusChangeAt)
		durations += currentDuration
	}

	durationStrTemp := durations.Truncate(time.Second).String()

	result := strings.Replace(durationStrTemp, "h", "小时", 1)
	result = strings.Replace(result, "m", "分", 1)
	result = strings.Replace(result, "s", "秒", 1)
	return result, nil
}
