package device

import (
	"github.com/SasukeBo/information/models"
	"github.com/graphql-go/graphql"
)

// ParamValueList 设备参数值列表查询
func ParamValueList(params graphql.ResolveParams) (interface{}, error) {
	deviceParam := models.DeviceParam{ID: params.Args["deviceParamID"].(int)}
	if err := deviceParam.Get(); err != nil {
		return nil, err
	}

	// 验证访问权限
	if err := deviceParam.ValidateAccess(params); err != nil {
		return nil, err
	}

	qs := models.Repo.QueryTable("device_param_value").Filter("device_param_id", deviceParam.ID).OrderBy("-created_at")

	if limit := params.Args["limit"]; limit != nil {
		qs = qs.Limit(limit)
	}

	if offset := params.Args["offset"]; offset != nil {
		qs = qs.Offset(offset)
	}

	if beforeTime := params.Args["beforeTime"]; beforeTime != nil {
		qs = qs.Filter("created_at__lt", beforeTime)
	}

	if afterTime := params.Args["afterTime"]; afterTime != nil {
		qs = qs.Filter("created_at__gt", afterTime)
	}

	var paramValues []*models.DeviceParamValue
	if _, err := qs.All(&paramValues); err != nil {
		return nil, err
	}

	return paramValues, nil
}

// ParamValueAdd _
func ParamValueAdd(params graphql.ResolveParams) (interface{}, error) {
	// paramID := params.Args["deviceParamID"]

	return nil, nil
}
