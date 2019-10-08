package resolver

import (
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
)

// CreateDeviceCharger 增加设备负责人
func CreateDeviceCharger(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	currentUser := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	device := models.Device{ID: params.Args["deviceID"].(int)}
	if err := o.Read(&device, "id"); err != nil {
		return nil, models.Error{Message: "get device failed.", OriErr: err}
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

	if _, err := o.Insert(&charger); err != nil {
		return nil, models.Error{Message: "insert device_charger failed.", OriErr: err}
	}

	return charger, nil
}

// DeleteDeviceCharger 删除设备负责人
func DeleteDeviceCharger(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	currentUser := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	id := params.Args["id"].(int)

	charger := models.DeviceCharger{ID: id}
	if err := o.Read(&charger, "id"); err != nil {
		return nil, models.Error{Message: "get device_charger failed.", OriErr: err}
	}

	device, err := charger.LoadDevice()
	if err != nil {
		return nil, err
	}

	if err := device.ValidateAccess(&currentUser); err != nil {
		return nil, err
	}

	if _, err := o.Delete(&charger); err != nil {
		return nil, models.Error{Message: "delete device_charger failed.", OriErr: err}
	}

	return charger.ID, nil
}

// UpdateDeviceCharger 更新设备负责人
func UpdateDeviceCharger(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	currentUser := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	charger := models.DeviceCharger{ID: params.Args["id"].(int)}
	if err := o.Read(&charger, "id"); err != nil {
		return nil, models.Error{Message: "get device_charger failed.", OriErr: err}
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

	if _, err := o.Update(&charger, "name", "phone", "department", "job_number"); err != nil {
		return nil, models.Error{Message: "update device_charger failed.", OriErr: err}
	}

	return charger, nil
}

// GetDeviceCharger ID查询设备负责人
func GetDeviceCharger(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	charger := models.DeviceCharger{ID: params.Args["id"].(int)}

	if err := o.Read(&charger, "id"); err != nil {
		return nil, models.Error{Message: "get device_charger failed.", OriErr: err}
	}

	return charger, nil
}

// ListDeviceCharger 条件查询设备负责关系列表
func ListDeviceCharger(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("device_charger")
	deviceID := params.Args["deviceID"].(int)
	device := models.Device{ID: deviceID}
	if err := o.Read(&device, "uuid"); err != nil {
		return nil, models.Error{Message: "get device failed.", OriErr: err}
	}

	qs.Filter("device_id", device.ID)
	var chargers []*models.DeviceCharger
	if _, err := qs.All(&chargers); err != nil {
		return nil, models.Error{Message: "list device_charger failed.", OriErr: err}
	}

	return chargers, nil
}

// LoadDeviceCharger 设备负责关系 load related
func LoadDeviceCharger(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.Device:
		return v.LoadDeviceCharge()
	case *models.Device:
		return v.LoadDeviceCharge()
	default:
		return nil, models.Error{Message: "load related device_charger failed."}
	}
}
