package resolver

import (
	// "fmt"
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
)

// GetDevice 获取设备
func GetDevice(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	id := params.Args["id"].(int)

	device := models.Device{ID: id}
	if err := o.Read(&device, "id"); err != nil {
		return nil, models.Error{Message: "get device failed.", OriErr: err}
	}

	return device, nil
}

// GetDeviceByToken 获取设备
func GetDeviceByToken(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	token := params.Args["token"].(string)

	device := models.Device{Token: token}
	if err := o.Read(&device, "token"); err != nil {
		return nil, models.Error{Message: "get device failed.", OriErr: err}
	}

	return device, nil
}

// ListDevice 获取设备列表
func ListDevice(params graphql.ResolveParams) (interface{}, error) {
	cond := orm.NewCondition()
	o := orm.NewOrm()

	if pattern := params.Args["search"]; pattern != nil {
		subCond := orm.NewCondition()
		cond = cond.AndCond(subCond.Or("type__icontains", pattern).Or("name__icontains", pattern).Or("address__icontains", pattern).Or("number__icontains", pattern))
	}

	if status := params.Args["status"]; status != nil {
		cond = cond.And("status", status)
	}

	if isRegister := params.Args["self"]; isRegister != nil {
		if isRegister.(bool) {
			user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
			cond = cond.And("user_id", user.ID)
		}
	}
	qs := o.QueryTable("device").SetCond(cond).OrderBy("created_at")

	cnt, err := qs.Count()
	if err != nil {
		return nil, err
	}

	if limit := params.Args["limit"]; limit != nil {
		qs = qs.Limit(limit)
	}

	if offset := params.Args["offset"]; offset != nil {
		qs = qs.Offset(offset)
	}

	var devices []*models.Device
	if _, err := qs.All(&devices); err != nil {
		return nil, models.Error{Message: "list device failed.", OriErr: err}
	}

	return struct {
		Total   int64
		Devices []*models.Device
	}{cnt, devices}, nil
}

// CreateDevice 创建设备
func CreateDevice(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		return nil, models.Error{Message: "begin transaction failed.", OriErr: err}
	}

	// 验证用户是否有创建设备的权限
	if err := utils.ValidateAccess(&params, "device_c", models.PrivType.Default); err != nil {
		o.Rollback()
		return nil, err
	}
	rootValue := params.Info.RootValue.(map[string]interface{})
	user := rootValue["currentUser"].(models.User)

	deviceType := params.Args["type"].(string)
	if err := utils.ValidateStringEmpty(deviceType, "type"); err != nil {
		o.Rollback()
		return nil, err
	}

	deviceName := params.Args["name"].(string)
	if err := utils.ValidateStringEmpty(deviceName, "name"); err != nil {
		o.Rollback()
		return nil, err
	}

	productID := params.Args["productID"].(int)
	privateForms := params.Args["privateForms"].([]interface{})
	count := 0
	for _, item := range privateForms {
		privateForm, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		device := &models.Device{
			Type:    deviceType,
			Name:    deviceName,
			Address: privateForm["address"].(string),
			Number:  privateForm["number"].(string),
			Token:   utils.GenRandomToken(8),
			User:    &user,
		}

		if _, err := o.Insert(device); err != nil {
			continue
		}

		ship := &models.DeviceProductShip{
			Device:  device,
			Product: &models.Product{ID: productID},
		}
		o.Insert(ship)

		count++
	}

	if count == 0 {
		o.Rollback()
		return nil, models.Error{Message: "device create failed."}
	}

	o.Commit()
	return count, nil
}

// UpdateDevice 更新设备
func UpdateDevice(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	device := models.Device{ID: params.Args["id"].(int)}
	if err := o.Read(&device, "id"); err != nil {
		return nil, models.Error{Message: "get device failed.", OriErr: err}
	}

	if err := device.ValidateAccess(&user); err != nil {
		return nil, err
	}

	if value := params.Args["address"]; value != nil {
		device.Address = value.(string)
	}

	if value := params.Args["name"]; value != nil {
		if err := utils.ValidateStringEmpty(value.(string), "name"); err != nil {
			return nil, err
		}
		device.Name = value.(string)
	}

	if value := params.Args["number"]; value != nil {
		device.Number = value.(string)
	}

	if value := params.Args["type"]; value != nil {
		if err := utils.ValidateStringEmpty(value.(string), "type"); err != nil {
			return nil, err
		}
		device.Type = value.(string)
	}

	if _, err := o.Update(&device); err != nil {
		return nil, models.Error{Message: "update device failed.", OriErr: err}
	}

	return device, nil
}

// DeleteDevice 更新设备
func DeleteDevice(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	device := models.Device{ID: params.Args["id"].(int)}
	if err := o.Read(&device, "id"); err != nil {
		return nil, models.Error{Message: "get device failed.", OriErr: err}
	}

	if err := device.ValidateAccess(&user); err != nil {
		return nil, err
	}

	if _, err := o.Delete(&device); err != nil {
		return nil, models.Error{Message: "delete device failed.", OriErr: err}
	}

	return device.ID, nil
}

// LoadDevice _
func LoadDevice(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.DeviceCharger:
		return v.LoadDevice()
	case *models.DeviceCharger:
		return v.LoadDevice()
	case models.DeviceStatusLog:
		return v.LoadDevice()
	case *models.DeviceStatusLog:
		return v.LoadDevice()
	default:
		return nil, models.Error{Message: "load related device_charger failed."}
	}
}

// CountDeviceStatus _
func CountDeviceStatus(params graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}
