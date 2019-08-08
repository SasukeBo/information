package device

import (
	"github.com/graphql-go/graphql"

	// "github.com/astaxie/beego/logs"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// ChargeCreate 指定设备负责人
func ChargeCreate(params graphql.ResolveParams) (interface{}, error) {
	// rootValue := params.Info.RootValue.(map[string]interface{})

	uuid := params.Args["uuid"].(string)
	userUUID := params.Args["userUUID"].(string)

	user := models.User{UUID: userUUID}
	if err := models.Repo.Read(&user, "uuid"); err != nil {
		return nil, utils.ORMError{
			Message: "user read error",
			OrmErr:  err,
		}
	}

	device := models.Device{UUID: uuid}
	if err := models.Repo.Read(&device, "uuid"); err != nil {
		return nil, utils.ORMError{
			Message: "device read error",
			OrmErr:  err,
		}
	}

	// TODO: 权限验证
	deviceCharge := models.DeviceCharge{
		User:   &user,
		Device: &device,
	}

	if _, err := models.Repo.Insert(&deviceCharge); err != nil {
		return nil, utils.ORMError{
			Message: "device_charge insert error",
			OrmErr:  err,
		}
	}

	return deviceCharge, nil
}

// ChargeDelete 取消指定设备负责人
func ChargeDelete(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	deviceCharge := models.DeviceCharge{ID: id}
	if err := models.Repo.Read(&deviceCharge); err != nil {
		return nil, utils.ORMError{
			Message: "device_charge read error",
			OrmErr:  err,
		}
	}

	if _, err := models.Repo.Delete(&deviceCharge); err != nil {
		return nil, utils.ORMError{
			Message: "device_charge delete error",
			OrmErr:  err,
		}
	}

	return "ok", nil
}

// ChargeUpdate 重新指定设备负责人
func ChargeUpdate(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	userUUID := params.Args["userUUID"].(string)

	deviceCharge := models.DeviceCharge{ID: id}
	if err := models.Repo.Read(&deviceCharge); err != nil {
		return nil, utils.ORMError{
			Message: "device_charge read error",
			OrmErr:  err,
		}
	}

	user := models.User{UUID: userUUID}
	if err := models.Repo.Read(&user, "uuid"); err != nil {
		return nil, utils.ORMError{
			Message: "user read error",
			OrmErr:  err,
		}
	}

	deviceCharge.User = &user
	if _, err := models.Repo.Update(&deviceCharge, "user_id"); err != nil {
		return nil, utils.ORMError{
			Message: "device_charge update error",
			OrmErr:  err,
		}
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
		return nil, utils.ORMError{
			Message: "device_charge list error",
			OrmErr:  err,
		}
	}

	return charges, nil
}

// ChargeRelatedLoad 设备负责关系 load related
func ChargeRelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	var id int

	switch v := params.Source.(type) {
	case models.DeviceChargeAbility:
		id = v.DeviceCharge.ID
	default:
		return nil, utils.LogicError{
			Message: "reloated device_charge load error",
		}
	}

	deviceCharge := models.DeviceCharge{ID: id}
	if err := deviceCharge.Get(); err != nil {
		return nil, err
	}

	return deviceCharge, nil
}
