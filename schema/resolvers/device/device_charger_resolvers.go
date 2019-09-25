package device

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
)

// ChargerCreate 指定设备负责人
func ChargerCreate(params graphql.ResolveParams) (interface{}, error) {
	currentUser := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	device := models.Device{UUID: params.Args["deviceUUID"].(string)}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	if err := device.ValidateAccess(&currentUser); err != nil {
		return nil, err
	}

	name := params.Args["name"].(string)
	charger := models.DeviceCharger{
		Name:   name,
		Device: &device,
	}

	if phone := params.Args["phone"]; phone != nil {
		charger.Phone = phone.(string)
	}

	if department := params.Args["department"]; department != nil {
		charger.Department = department.(string)
	}

	if jobNumber := params.Args["jobNumber"]; jobNumber != nil {
		charger.JobNumber = jobNumber.(string)
	}

	if err := charger.Insert(); err != nil {
		return nil, err
	}

	return charger, nil
}

// ChargerDelete 删除设备负责人
func ChargerDelete(params graphql.ResolveParams) (interface{}, error) {
	currentUser := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	id := params.Args["id"].(int)

	charger := models.DeviceCharger{ID: id}
	if err := charger.GetBy("id"); err != nil {
		return nil, err
	}

	device, err := charger.LoadDevice()
	if err != nil {
		return nil, err
	}

	if err := device.ValidateAccess(&currentUser); err != nil {
		return nil, err
	}

	if err := charger.Delete(); err != nil {
		return nil, err
	}

	return charger.ID, nil
}

// ChargerUpdate 更新设备负责人
func ChargerUpdate(params graphql.ResolveParams) (interface{}, error) {
	currentUser := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	charger := models.DeviceCharger{ID: params.Args["id"].(int)}
	if err := charger.GetBy("id"); err != nil {
		return nil, err
	}

	device, err := charger.LoadDevice()
	if err != nil {
		return nil, err
	}

	if err := device.ValidateAccess(&currentUser); err != nil {
		return nil, err
	}

	charger.Name = params.Args["name"].(string)
	if phone := params.Args["phone"]; phone != nil {
		charger.Phone = phone.(string)
	}

	if department := params.Args["department"]; department != nil {
		charger.Department = department.(string)
	}

	if jobNumber := params.Args["jobNumber"]; jobNumber != nil {
		charger.JobNumber = jobNumber.(string)
	}

	if err := charger.Update("name", "phone", "department", "job_number"); err != nil {
		return nil, err
	}

	return charger, nil
}

// ChargerGet ID查询设备负责关系
func ChargerGet(params graphql.ResolveParams) (interface{}, error) {
	charger := models.DeviceCharger{ID: params.Args["id"].(int)}

	if err := charger.GetBy("id"); err != nil {
		return nil, err
	}

	return charger, nil
}

// ChargerList 条件查询设备负责关系列表
func ChargerList(params graphql.ResolveParams) (interface{}, error) {
	qs := models.Repo.QueryTable("device_charger")
	deviceUUID := params.Args["deviceUUID"].(string)
	device := models.Device{UUID: deviceUUID}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	qs.Filter("device_id", device.ID)
	var chargers []*models.DeviceCharger
	if _, err := qs.All(&chargers); err != nil {
		return nil, errors.LogicError{
			Type:    "Models",
			Message: "get device_chargers error",
			OriErr:  err,
		}
	}

	return chargers, nil
}

// ChargeRelatedLoad 设备负责关系 load related
func ChargeRelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
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
