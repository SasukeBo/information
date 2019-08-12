package device

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/errors"
	"github.com/SasukeBo/information/models"
)

// ChargeCreate 指定设备负责人
func ChargeCreate(params graphql.ResolveParams) (interface{}, error) {
	// rootValue := params.Info.RootValue.(map[string]interface{})

	uuid := params.Args["uuid"].(string)
	userUUID := params.Args["userUUID"].(string)

	user := models.User{UUID: userUUID}
	if err := user.GetBy("uuid"); err != nil {
		return nil, err
	}

	device := models.Device{UUID: uuid}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	// TODO: 权限验证
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
	if err := deviceCharge.Delete(); err != nil {
		return nil, err
	}

	return "ok", nil
}

// ChargeUpdate 重新指定设备负责人
func ChargeUpdate(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	userUUID := params.Args["userUUID"].(string)

	deviceCharge := models.DeviceCharge{ID: id}
	if err := deviceCharge.Get(); err != nil {
		return nil, err
	}

	user := models.User{UUID: userUUID}
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
	id := params.Args["id"].(int)

	deviceCharge := models.DeviceCharge{ID: id}

	if err := deviceCharge.Get(); err != nil {
		return nil, err
	}

	return deviceCharge, nil
}

// ChargeList 条件查询设备负责关系列表
func ChargeList(params graphql.ResolveParams) (interface{}, error) {
	userUUID := params.Args["userUUID"]
	deviceUUID := params.Args["deviceUUID"]

	qs := models.Repo.QueryTable("device_charge")

	if userUUID != nil {
		qs = qs.Filter("user__uuid", userUUID.(string))
	}

	if deviceUUID != nil {
		qs = qs.Filter("device__uuid", deviceUUID.(string))
	}

	var charges []*models.DeviceCharge

	if _, err := qs.All(&charges); err != nil {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "DeviceCharge",
			Message: "ChargeList() error",
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
	default:
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "DeviceCharge",
			Message: "ChargeRelatedLoad() error",
		}
	}
}
