package device

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/utils"
)

// Get 获取设备
func Get(params graphql.ResolveParams) (interface{}, error) {
	uuid := params.Args["uuid"].(string)

	device := models.Device{UUID: uuid}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	return device, nil
}

// List _
func List(params graphql.ResolveParams) (interface{}, error) {
	dType := params.Args["type"]
	namePattern := params.Args["namePattern"]
	status := params.Args["status"]
	userUUID := params.Args["userUUID"]

	qs := models.Repo.QueryTable("device")

	if dType != nil {
		qs = qs.Filter("type", dType.(string))
	}

	if namePattern != nil {
		qs = qs.Filter("name__icontains", namePattern.(string))
	}

	if status != nil {
		qs = qs.Filter("status", status.(int))
	}

	if userUUID != nil {
		qs = qs.Filter("user__uuid", userUUID.(string))
	}

	var devices []*models.Device

	if _, err := qs.All(&devices); err != nil {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "Device",
			Message: "List() error",
			OriErr:  err,
		}
	}

	return devices, nil
}

// Create 创建设备
func Create(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})

	dType := params.Args["type"].(string)
	if err := utils.ValidateStringEmpty(dType, "type"); err != nil {
		return nil, err
	}

	dName := params.Args["name"].(string)
	if err := utils.ValidateStringEmpty(dName, "name"); err != nil {
		return nil, err
	}

	token := utils.GenRandomToken(8)
	description := params.Args["description"].(string)
	uuid := uuid.New().String()
	user := rootValue["currentUser"].(models.User)

	device := models.Device{
		Type:        dType,
		Name:        dName,
		Token:       token,
		UUID:        uuid,
		User:        &user,
		Description: description,
	}

	if err := device.Insert(); err != nil {
		return nil, err
	}

	return device, nil
}

// Update 更新设备
func Update(params graphql.ResolveParams) (interface{}, error) {
	uuid := params.Args["uuid"].(string)

	device := models.Device{UUID: uuid}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	if dType := params.Args["type"]; dType != nil {
		if err := utils.ValidateStringEmpty(dType.(string), "type"); err != nil {
			return nil, err
		}
		device.Type = dType.(string)
	}

	if name := params.Args["name"]; name != nil {
		if err := utils.ValidateStringEmpty(name.(string), "name"); err != nil {
			return nil, err
		}
		device.Name = name.(string)
	}

	if status := params.Args["status"]; status != nil {
		device.Status = status.(int)
	}

	if description := params.Args["description"]; description != nil {
		device.Description = description.(string)
	}

	if err := device.Update("type", "name", "status", "description"); err != nil {
		return nil, err
	}

	return device, nil
}

// Delete 更新设备
func Delete(params graphql.ResolveParams) (interface{}, error) {
	uuid := params.Args["uuid"].(string)

	device := models.Device{UUID: uuid}
	if err := device.DeleteByUUID(); err != nil {
		return nil, err
	}

	return "ok", nil
}

// Bind 绑定设备Mac地址，需要权限验证
func Bind(params graphql.ResolveParams) (interface{}, error) {
	// rootValue := params.Info.RootValue.(map[string]interface{})
	token := params.Args["token"].(string)
	// user := rootValue["currentUser"].(models.User)

	mac := params.Args["mac"].(string)
	if err := utils.ValidateStringEmpty(mac, "mac"); err != nil {
		return nil, err
	}

	// TODO: 验证绑定设备的权限
	device := models.Device{Token: token}
	if err := device.GetBy("token"); err != nil {
		return nil, err
	}

	// TODO: 设备状态
	device.Mac = mac

	if err := device.Update("mac"); err != nil {
		return nil, err
	}

	return device, nil
}

// RelatedLoad _
func RelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.DeviceCharge:
		return v.LoadDevice()
	case *models.DeviceCharge:
		return v.LoadDevice()
	case models.DeviceStatusLog:
		return v.LoadDevice()
	case *models.DeviceStatusLog:
		return v.LoadDevice()
	case models.DeviceParam:
		return v.LoadDevice()
	case *models.DeviceParam:
		return v.LoadDevice()
	default:
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "Device",
			Message: "RelatedLoad() error",
		}
	}
}
