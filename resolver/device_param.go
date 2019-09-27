package resolver

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/utils"
)

// CreateDeviceParam 创建设备参数
func CreateDeviceParam(params graphql.ResolveParams) (interface{}, error) {
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	device := models.Device{UUID: params.Args["deviceUUID"].(string)}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}
	// 创建参数的权限验证
	if err := device.ValidateAccess(&user); err != nil {
		return nil, err
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

	deviceParam := models.DeviceParam{Name: name, Sign: sign, Type: pType, Device: &device}
	if err := deviceParam.Insert(); err != nil {
		return nil, err
	}

	return deviceParam, nil
}

// UpdateDeviceParam 修改设备参数
func UpdateDeviceParam(params graphql.ResolveParams) (interface{}, error) {
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	deviceParam := models.DeviceParam{ID: params.Args["id"].(int)}
	if err := deviceParam.Get(); err != nil {
		return nil, err
	}

	device, err := deviceParam.LoadDevice()
	if err != nil {
		return nil, err
	}

	if err := device.ValidateAccess(&user); err != nil {
		return nil, err
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

// DeleteDeviceParam 删除设备参数
func DeleteDeviceParam(params graphql.ResolveParams) (interface{}, error) {
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	id := params.Args["id"].(int)
	deviceParam := models.DeviceParam{ID: id}

	if err := deviceParam.Get(); err != nil {
		return nil, err
	}

	device, err := deviceParam.LoadDevice()
	if err != nil {
		return nil, err
	}

	if err := device.ValidateAccess(&user); err != nil {
		return nil, err
	}

	if err := deviceParam.Delete(); err != nil {
		return nil, err
	}

	return id, nil
}

// GetDeviceParam 获取设备参数
func GetDeviceParam(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	deviceParam := models.DeviceParam{ID: id}
	if err := deviceParam.Get(); err != nil {
		return nil, err
	}

	return deviceParam, nil
}

// ListDeviceParam 根据条件获取设备参数列表
func ListDeviceParam(params graphql.ResolveParams) (interface{}, error) {
	device := models.Device{UUID: params.Args["deviceUUID"].(string)}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
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

// LoadDeviceParam 根据条件获取设备参数列表
func LoadDeviceParam(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.DeviceParamValue:
		return v.LoadDeviceParam()
	case *models.DeviceParamValue:
		return v.LoadDeviceParam()
	case models.Device:
		return v.LoadDeviceParams()
	case *models.Device:
		return v.LoadDeviceParams()
	default:
		return nil, errors.LogicError{
			Type:    "Resolver",
			Message: "load related source type unmatched error.",
		}
	}
}
