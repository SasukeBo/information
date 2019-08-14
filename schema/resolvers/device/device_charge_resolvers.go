package device

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
)

// ChargeCreate 指定设备负责人
func ChargeCreate(params graphql.ResolveParams) (interface{}, error) {
	device := models.Device{UUID: params.Args["uuid"].(string)}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	if err := device.ValidateAccess(params, "device_charge_c"); err != nil {
		return nil, err
	}

	user := models.User{UUID: params.Args["userUUID"].(string)}
	if err := user.GetBy("uuid"); err != nil {
		return nil, err
	}

	deviceCharge := models.DeviceCharge{
		User:   &user,
		Device: &device,
	}

	if err := deviceCharge.Insert(); err != nil {
		return nil, err
	}

	return deviceCharge, nil
}

// ChargeDelete 取消指定设备负责人
func ChargeDelete(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	deviceCharge := models.DeviceCharge{ID: id}
	if err := deviceCharge.Get(); err != nil {
		return nil, err
	}

	// 校验用户是否可以删除设备负责人
	if accessErr := deviceCharge.ValidateAccess(params, "device_charge_d"); accessErr != nil {
		return nil, accessErr
	}

	if err := deviceCharge.Delete(); err != nil {
		return nil, err
	}

	return "ok", nil
}

// ChargeUpdate 重新指定设备负责人
func ChargeUpdate(params graphql.ResolveParams) (interface{}, error) {
	deviceCharge := models.DeviceCharge{ID: params.Args["id"].(int)}
	if err := deviceCharge.Get(); err != nil {
		return nil, err
	}

	// 校验用户是否可以修改设备负责人
	if accessErr := deviceCharge.ValidateAccess(params, "device_charge_u"); accessErr != nil {
		return nil, accessErr
	}

	user := models.User{UUID: params.Args["userUUID"].(string)}
	if err := user.GetBy("uuid"); err != nil {
		return nil, err
	}

	deviceCharge.User = &user
	if err := deviceCharge.Update("user_id"); err != nil {
		return nil, err
	}

	return deviceCharge, nil
}

// ChargeGet ID查询设备负责关系
func ChargeGet(params graphql.ResolveParams) (interface{}, error) {
	deviceCharge := models.DeviceCharge{ID: params.Args["id"].(int)}
	if err := deviceCharge.Get(); err != nil {
		return nil, err
	}
	// 校验用户是否为设备的负责人
	if accessErr := deviceCharge.ValidateAccess(params); accessErr != nil {
		return nil, accessErr
	}

	return deviceCharge, nil
}

// ChargeList 条件查询设备负责关系列表
func ChargeList(params graphql.ResolveParams) (interface{}, error) {
	qs := models.Repo.QueryTable("device_charge")
	if deviceUUID := params.Args["deviceUUID"]; deviceUUID != nil {
		// 指定了设备 uuid
		// 仅当用户 注册了该设备 或 负责该设备才获取
		device := models.Device{UUID: deviceUUID.(string)}
		if err := device.GetBy("uuid"); err != nil {
			return nil, err
		}

		if err := device.ValidateAccess(params); err != nil {
			return nil, err
		}

		qs = qs.Filter("device__uuid", deviceUUID)
	} else { // 没有指定设备uuid，获取用户负责的设备 或 用户注册的设备
		currentUser := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
		cond := models.NewCond().And("device__user_id", currentUser.ID)
		var deviceCharges []*models.DeviceCharge
		if _, err := models.Repo.QueryTable("device_charge").Filter("user_id", currentUser.ID).All(&deviceCharges); err != nil {
			return nil, errors.LogicError{
				Type:    "Model",
				Message: "get current user's device_charge list error.",
				OriErr:  err,
			}
		}

		var ids []int
		for _, charge := range deviceCharges {
			ids = append(ids, charge.Device.ID)
		}

		if len(ids) > 0 {
			cond = cond.Or("device_id__in", ids)
		}

		qs.SetCond(cond)
	}

	var charges []*models.DeviceCharge

	if _, err := qs.All(&charges); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "get device_charge list error.",
			OriErr:  err,
		}
	}

	return charges, nil
}

// ChargeRelatedLoad 设备负责关系 load related
func ChargeRelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.DeviceChargeAbility:
		return v.LoadDeviceCharge()
	case *models.DeviceChargeAbility:
		return v.LoadDeviceCharge()
	case models.Device:
		return v.LoadDeviceCharge()
	case *models.Device:
		return v.LoadDeviceCharge()
	default:
		return nil, errors.LogicError{
			Type:    "Resolver",
			Message: "load related source type unmatched error.",
		}
	}
}
