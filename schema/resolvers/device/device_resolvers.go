package device

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"

	// "github.com/astaxie/beego/logs"

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

	if err := device.ValidateAccess(params); err != nil {
		return nil, err
	}

	return device, nil
}

// List 获取负责或创建的设备
func List(params graphql.ResolveParams) (interface{}, error) {
	qs := models.Repo.QueryTable("device")
	cond1 := models.NewCond()

	if dType := params.Args["type"]; dType != nil {
		cond1 = cond1.And("type", dType)
	}

	if namePattern := params.Args["namePattern"]; namePattern != nil {
		cond1 = cond1.And("name__icontains", namePattern)
	}

	if status := params.Args["status"]; status != nil {
		cond1 = cond1.And("status", status)
	}

	// 限定用户查询设备列表值域为 负责的设备 + 注册的设备 -- begin
	var charges []*models.DeviceCharge
	currentUser := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)

	if _, err := models.Repo.QueryTable("device_charge").Filter("user_id", currentUser.ID).All(&charges); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Field:   "userUUID",
			Message: "get device_charge list error",
			OriErr:  err,
		}
	}

	var ids []int
	for _, charge := range charges { // 获取用户charge的设备id列表
		ids = append(ids, charge.Device.ID)
	}

	if userUUID := params.Args["userUUID"]; userUUID != nil {
		// 如果提供了注册人uuid，则从当前用户负责的设备中挑出注册人为uuid的设备
		if len(ids) == 0 {
			// 如果用户负责设备个数为0，返回空列表
			return []interface{}{}, nil
		}

		// 获取该用户负责的且由userUUID的用户注册的设备
		cond := models.NewCond().AndCond(cond1).AndCond(models.NewCond().And("id__in", ids).And("user__uuid", userUUID))
		qs = qs.SetCond(cond)
	} else {
		// 获取用户负责的或注册的设备
		cond := models.NewCond().AndCond(cond1).AndCond(models.NewCond().And("id__in", ids).Or("user_id", currentUser.ID))
		qs = qs.SetCond(cond).Distinct()
	}
	// 限定用户查询设备列表值域为 负责的设备 + 注册的设备 -- end

	var devices []*models.Device

	if _, err := qs.All(&devices); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "get device list error.",
			OriErr:  err,
		}
	}

	return devices, nil
}

// Create 创建设备
func Create(params graphql.ResolveParams) (interface{}, error) {
	// 验证用户是否有创建设备的权限
	if err := utils.ValidateAccess(&params, "device_c", models.PrivType.Default); err != nil {
		return nil, err
	}

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
	device := models.Device{UUID: params.Args["uuid"].(string)}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	if err := device.ValidateAccess(params, "device_u"); err != nil {
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
	device := models.Device{UUID: params.Args["uuid"].(string)}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	if err := device.ValidateAccess(params, "device_d"); err != nil {
		return nil, err
	}

	if err := device.Delete(); err != nil {
		return nil, err
	}

	return "ok", nil
}

// Bind 绑定设备Mac地址，需要权限验证
func Bind(params graphql.ResolveParams) (interface{}, error) {
	mac := params.Args["mac"].(string)
	if err := utils.ValidateStringEmpty(mac, "mac"); err != nil {
		return nil, err
	}

	device := models.Device{Token: params.Args["token"].(string)}
	if err := device.GetBy("token"); err != nil {
		return nil, err
	}

	if err := device.ValidateAccess(params, "device_u"); err != nil {
		return nil, err
	}

	device.Status = models.BaseStatus.Publish
	device.Mac = mac

	if err := device.Update("mac", "status"); err != nil {
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
			Message: "load related source type unmatched error.",
		}
	}
}
