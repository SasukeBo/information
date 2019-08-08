package device

import (
  "time"

  "github.com/SasukeBo/information/models"
  "github.com/graphql-go/graphql"
)

// ParamValueList 设备参数创建
func ParamValueList(params graphql.ResolveParams) (interface{}, error) {
  deviceParamID := params.Args["deviceParamID"].(int)
  beforeTime := params.Args["beforeTime"]
  afterTime := params.Args["afterTime"]

  qs := models.Repo.QueryTable("device_param_value").Filter("device_param_id", deviceParamID)

  if beforeTime != nil {
    qs = qs.Filter("created_at__lt", beforeTime.(time.Time))
  }

  if afterTime != nil {
    qs = qs.Filter("created_at__gt", afterTime.(time.Time))
  }

  var paramValues []*models.DeviceParamValue
  if _, err := qs.All(&paramValues); err != nil {
    return nil, err
  }

  return paramValues, nil
}
