package device

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/utils"
)

// ParamCreate 设备参数创建
func ParamCreate(params graphql.ResolveParams) (interface{}, error) {
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)

	device := models.Device{ID: params.Args["deviceID"].(int)}
	if err := device.GetBy("id"); err != nil {
		return nil, err
	}
	// 创建参数的权限验证
	if accessErr := device.ValidateAccess(params, "device_param_c"); accessErr != nil {
		return nil, accessErr
	}

	name := params.Args["name"].(string)
	if err := utils.ValidateStringEmpty(name, "name"); err != nil {
		return nil, err
	}

	sign := params.Args["sign"].(string)
	if err := utils.ValidateStringEmpty(sign, "sign"); err != nil {
		return nil, err
	}

	pType := params.Args["type"].(int)

	deviceParam := models.DeviceParam{Name: name, Sign: sign, Type: pType, Author: &user, Device: &device}
	if err := deviceParam.Insert(); err != nil {
		return nil, err
	}

	return deviceParam, nil
}

// ParamUpdate 设备参数修改
func ParamUpdate(params graphql.ResolveParams) (interface{}, error) {
	deviceParam := models.DeviceParam{ID: params.Args["id"].(int)}
	if err := deviceParam.Get(); err != nil {
		return nil, err
	}

	// 修改设备参数的权限
	if accessErr := deviceParam.ValidateAccess(params, "device_param_u"); accessErr != nil {
		return nil, accessErr
	}

	if name := params.Args["name"]; name != nil {
		if err := utils.ValidateStringEmpty(name.(string), "name"); err != nil {
			return nil, err
		}
		deviceParam.Name = name.(string)
	}

	if sign := params.Args["sign"]; sign != nil {
		if err := utils.ValidateStringEmpty(sign.(string), "sign"); err != nil {
			return nil, err
		}
		deviceParam.Sign = sign.(string)
	}

	if pType := params.Args["type"]; pType != nil {
		deviceParam.Type = pType.(int)
	}

	if err := deviceParam.Update("name", "sign", "type"); err != nil {
		return nil, err
	}

	return deviceParam, nil
}

// ParamDelete 设备参数删除
func ParamDelete(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	deviceParam := models.DeviceParam{ID: id}
	if err := deviceParam.Get(); err != nil {
		return nil, err
	}

	// 删除权限验证
	if accessErr := deviceParam.ValidateAccess(params, "device_param_u"); accessErr != nil {
		return nil, accessErr
	}

	if err := deviceParam.Delete(); err != nil {
		return nil, err
	}

	return "ok", nil
}

// ParamGet ID获取设备参数
func ParamGet(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	deviceParam := models.DeviceParam{ID: id}
	if err := deviceParam.Get(); err != nil {
		return nil, err
	}

	if accessErr := deviceParam.ValidateAccess(params); accessErr != nil {
		return nil, accessErr
	}

	return deviceParam, nil
}

// ParamList 根据条件获取设备参数列表
func ParamList(params graphql.ResolveParams) (interface{}, error) {
	device := models.Device{UUID: params.Args["deviceUUID"].(string)}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	// 验证访问权限
	if accessErr := device.ValidateAccess(params); accessErr != nil {
		return nil, accessErr
	}

	qs := models.Repo.QueryTable("device_param").Filter("device_id", device.ID).OrderBy("-created_at")

	if namePattern := params.Args["namePattern"]; namePattern != nil {
		qs = qs.Filter("name__icontains", namePattern)
	}

	if signPattern := params.Args["signPattern"]; signPattern != nil {
		qs = qs.Filter("name__icontains", signPattern)
	}

	if pType := params.Args["type"]; pType != nil {
		qs = qs.Filter("type", pType)
	}

	if userUUID := params.Args["userUUID"]; userUUID != nil {
		qs = qs.Filter("author__uuid", userUUID)
	}

	var deviceParams []*models.DeviceParam

	if _, err := qs.All(&deviceParams); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "get device_param list error",
			OriErr:  err,
		}
	}

	return deviceParams, nil
}

// ParamRelatedLoad 根据条件获取设备参数列表
func ParamRelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.DeviceParamValue:
		return v.LoadDeviceParam()
	case *models.DeviceParamValue:
		return v.LoadDeviceParam()
	default:
		return nil, errors.LogicError{
			Type:    "Resolver",
			Message: "load related source type unmatched error.",
		}
	}
}
