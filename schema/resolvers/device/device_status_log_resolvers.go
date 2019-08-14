package device

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
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

	qs := models.Repo.QueryTable("device_status_log").Filter("device_id", device.ID)

	if status := params.Args["status"]; status != nil {
		qs = qs.Filter("status", status)
	}

	if beforeTime := params.Args["beforeTime"]; beforeTime != nil {
		qs = qs.Filter("change_at__lt", beforeTime)
	}

	if afterTime := params.Args["afterTime"]; afterTime != nil {
		qs = qs.Filter("change_at__gt", afterTime)
	}

	var statusLogs []*models.DeviceStatusLog
	if _, err := qs.All(&statusLogs); err != nil {
		return nil, err
	}

	return statusLogs, nil
}
