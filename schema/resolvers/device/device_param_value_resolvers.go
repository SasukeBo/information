package device

import (
	"fmt"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
	"github.com/graphql-go/graphql"
)

// ParamValueList 设备参数值列表查询
func ParamValueList(params graphql.ResolveParams) (interface{}, error) {
	deviceParam := models.DeviceParam{ID: params.Args["deviceParamID"].(int)}
	if err := deviceParam.Get(); err != nil {
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
	id := params.Args["id"].(int)
	value := models.DeviceParamValue{ID: id}
	if err := value.GetBy("id"); err != nil {
		return nil, err
	}

	return value, nil
}

// ParamValueCount _
func ParamValueCount(params graphql.ResolveParams) (interface{}, error) {
	deviceUUID := params.Args["deviceUUID"].(string)
	device := models.Device{UUID: deviceUUID}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	deviceParams, err := device.LoadDeviceParams()
	if err != nil {
		return nil, err
	}
	if len(deviceParams) == 0 {
		return nil, nil
	}

	durationCond := models.NewCond()
	if beforeTime := params.Args["beforeTime"]; beforeTime != nil {
		durationCond = durationCond.And("created_at__lt", beforeTime)
	}

	if afterTime := params.Args["afterTime"]; afterTime != nil {
		durationCond = durationCond.And("created_at__gt", afterTime)
	}

	cnts := make([]int64, 0)
	for _, deviceParam := range deviceParams {
		cond := models.NewCond().And("device_param_id", deviceParam.ID)
		if !durationCond.IsEmpty() {
			cond = cond.AndCond(durationCond)
		}
		qs := models.Repo.QueryTable("device_param_value").SetCond(cond)
		cnt, err := qs.Count()

		if err != nil {
			cnts = append(cnts, 0)
		} else {
			cnts = append(cnts, cnt)
		}
	}

	maxCnt := int64(0)

	for _, cnt := range cnts {
		if cnt > maxCnt {
			maxCnt = cnt
		}
	}

	return maxCnt, nil
}

// ParamValueHistogram _
func ParamValueHistogram(params graphql.ResolveParams) (interface{}, error) {
	var valueMin float32
	var valueMax float32
	var err error
	category := make([]string, 0)
	serie := make([]int, 0)
	paramID := params.Args["paramID"].(int)

	err = models.Repo.Raw("select min(value) as value_min from device_param_value where device_param_id = ?", paramID).QueryRow(&valueMin)
	if err != nil {
		return nil, err
	}

	err = models.Repo.Raw("select max(value) as value_max from device_param_value where device_param_id = ?", paramID).QueryRow(&valueMax)
	if err != nil {
		return nil, err
	}

	duration := (valueMax - valueMin) / 20
	queryCount := models.Repo.Raw("select count(*) from device_param_value where cast(value as float) > ? and cast(value as float) < ?")
	for i := 0; i < 20; i++ {
		var count int
		upLimit := valueMin + duration*float32(i+1)
		lowerLimit := valueMin + duration*float32(i)

		queryCount.SetArgs(lowerLimit, upLimit).QueryRow(&count)
		category = append(category, fmt.Sprintf("%.2f-%.2f", lowerLimit, upLimit))
		serie = append(serie, count)
	}

	result := struct {
		Category []string
		Serie    []int
	}{category, serie}

	return result, nil
}

// ValueRelatedLoad _
func ValueRelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.DeviceParam:
		return v.LoadDeviceParamValues(params)
	case *models.DeviceParam:
		return v.LoadDeviceParamValues(params)
	default:
		return nil, errors.LogicError{
			Type:    "Resolver",
			Message: "load related source type unmatched error.",
		}
	}
}
