package device

import (
  "time"

  "github.com/SasukeBo/information/models"
  "github.com/graphql-go/graphql"
)

// StatusLogList 设备参数创建
func StatusLogList(params graphql.ResolveParams) (interface{}, error) {
  deviceID := params.Args["deviceID"].(int)
  status := params.Args["status"]
  beforeTime := params.Args["beforeTime"]
  afterTime := params.Args["afterTime"]

  qs := models.Repo.QueryTable("device_status_log").Filter("device_id", deviceID)

  if status != nil {
    qs = qs.Filter("status", status.(int))
  }

  if beforeTime != nil {
    qs = qs.Filter("change_at__lt", beforeTime.(time.Time))
  }

  if afterTime != nil {
    qs = qs.Filter("change_at__gt", afterTime.(time.Time))
  }

  var statusLogs []*models.DeviceStatusLog
  if _, err := qs.All(&statusLogs); err != nil {
    return nil, err
  }

  return statusLogs, nil
}
