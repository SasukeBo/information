package device

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// ParamCreate 设备参数创建
func ParamCreate(params graphql.ResolveParams) (interface{}, error) {
	currentUserUUID := params.Info.RootValue.(map[string]interface{})["currentUserUUID"].(string)
	user := models.User{UUID: currentUserUUID}
	if err := user.GetBy("uuid"); err != nil {
		return nil, err
	}

	deviceID := params.Args["deviceID"].(int)
	device := models.Device{ID: deviceID}
	if err := device.GetBy("id"); err != nil {
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

	pType := params.Args["type"].(string)
	if err := utils.ValidateStringEmpty(pType, "type"); err != nil {
		return nil, err
	}

	deviceParam := models.DeviceParam{Name: name, Sign: sign, Type: pType, Author: &user, Device: &device}
	if err := deviceParam.Insert(); err != nil {
		return nil, err
	}

	return deviceParam, nil
}

// ParamUpdate 设备参数修改
func ParamUpdate(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	deviceParam := models.DeviceParam{ID: id}
	if err := deviceParam.Get(); err != nil {
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
		if err := utils.ValidateStringEmpty(pType.(string), "type"); err != nil {
			return nil, err
		}
		deviceParam.Type = pType.(string)
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

	return deviceParam, nil
}

// ParamList 根据条件获取设备参数列表
func ParamList(params graphql.ResolveParams) (interface{}, error) {
	qs := models.Repo.QueryTable("device_param")

	namePattern := params.Args["namePattern"]
	signPattern := params.Args["signPattern"]
	pType := params.Args["type"]
	userUUID := params.Args["userUUID"]

	if namePattern != nil {
		qs = qs.Filter("name__icontains", namePattern.(string))
	}

	if signPattern != nil {
		qs = qs.Filter("name__icontains", signPattern.(string))
	}

	if pType != nil {
		qs = qs.Filter("type", pType.(string))
	}

	if userUUID != nil {
		qs = qs.Filter("author__uuid", userUUID.(string))
	}

	var deviceParams []*models.DeviceParam

	if _, err := qs.All(&deviceParams); err != nil {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "DeviceParam",
			Message: "List() error",
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
			Field:   "DeviceParam",
			Message: "ParamRelatedLoad() error",
		}
	}
}
