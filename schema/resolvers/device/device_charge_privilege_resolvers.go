package device

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
)

/*
FIXME: 暂时不需要这两个接口
// ChargePrivGet 获取设备负责人权限
func ChargePrivGet(params graphql.ResolveParams) (interface{}, error) {
	deviceChargeAbility := models.DeviceChargeAbility{ID: params.Args["id"].(int)}
	if err := deviceChargeAbility.GetBy("id"); err != nil {
		return nil, err
	}

	var deviceCharge *models.DeviceCharge
	var err error
	if deviceCharge

	return deviceChargeAbility, nil
}

// ChargePrivList 根据设备负责人关系ID获取权限list
func ChargePrivList(params graphql.ResolveParams) (interface{}, error) {
	deviceChargeID := params.Args["deviceChargeID"].(int)
	qs := models.Repo.QueryTable("device_charge_ability")

	var chargePrivs []*models.DeviceChargeAbility
	if _, err := qs.Filter("device_charge_id", deviceChargeID).All(&chargePrivs); err != nil {
		return nil, err
	}

	return chargePrivs, nil
}
*/

// ChargePrivCreate 为设备负责人添加权限
func ChargePrivCreate(params graphql.ResolveParams) (interface{}, error) {
	deviceCharge := models.DeviceCharge{ID: params.Args["deviceChargeID"].(int)}
	if err := deviceCharge.Get(); err != nil {
		return nil, err
	}

	// 验证增加权限
	if accessErr := deviceCharge.ValidateAccess(params, "device_charge_ability_c"); accessErr != nil {
		return nil, accessErr
	}

	privilege := models.Privilege{ID: params.Args["privilegeID"].(int)}
	if err := privilege.Get(); err != nil {
		return nil, err
	}

	deviceChargeAbility := models.DeviceChargeAbility{
		DeviceCharge: &deviceCharge,
		Privilege:    &privilege,
	}

	if _, err := models.Repo.Insert(&deviceChargeAbility); err != nil {
		return nil, err
	}

	return deviceChargeAbility, nil
}

// ChargePrivDelete 删除设备负责人的权限
func ChargePrivDelete(params graphql.ResolveParams) (interface{}, error) {
	deviceChargeAbility := models.DeviceChargeAbility{ID: params.Args["id"].(int)}
	if err := deviceChargeAbility.GetBy("id"); err != nil {
		return nil, err
	}

	// 验证删除权限
	if accessErr := deviceChargeAbility.ValidateAccess(params, "device_charge_ability_d"); accessErr != nil {
		return nil, accessErr
	}

	if err := deviceChargeAbility.Delete(); err != nil {
		return nil, err
	}

	return "ok", nil
}

// ChargeAbilityRelatedLoad _
func ChargeAbilityRelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.DeviceCharge:
		return v.LoadDeviceChargeAbility()
	case *models.DeviceCharge:
		return v.LoadDeviceChargeAbility()
	default:
		return nil, errors.LogicError{
			Type:    "Resolver",
			Message: "load related source type unmatched error.",
		}
	}
}
