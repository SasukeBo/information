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
	cond := models.NewCond()

	ownership, ok := params.Args["ownership"].(string)
	if !ok {
		return nil, errors.LogicError{
			Type:    "Resolvers",
			Field:   "ownership",
			Message: "invalid value of ownership",
		}
	}

	currentUser := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)

	switch ownership {
	case "register":
		cond = cond.And("user_id", currentUser.ID)
	case "charger":
		ids := chargeIDs(&currentUser)
		cond = cond.And("id__in", ids)
	case "both":
		ids := chargeIDs(&currentUser)
		subCond := models.NewCond().And("id__in", ids).Or("user_id", currentUser.ID)
		cond = cond.AndCond(subCond)
	}

	if dType := params.Args["type"]; dType != nil {
		cond = cond.And("type", dType)
	}

	if namePattern := params.Args["namePattern"]; namePattern != nil {
		cond = cond.And("name__icontains", namePattern)
	}

	if status := params.Args["status"]; status != nil {
		cond = cond.And("status", status)
	}

	// 只有当用户本人是负责人时，创建者uuid才是有效筛选条件
	if userUUID := params.Args["userUUID"]; userUUID != nil && ownership == "charger" {
		cond = cond.And("user__uuid", userUUID)
	}

	qs = qs.SetCond(cond).Distinct()
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

func chargeIDs(user *models.User) []int {
	var charges []*models.DeviceCharge

	if _, err := models.Repo.QueryTable("device_charge").Filter("user_id", user.ID).All(&charges); err != nil {
		return []int{}
	}

	var ids []int
	for _, charge := range charges { // 获取用户charge的设备id列表
		ids = append(ids, charge.Device.ID)
	}

	return ids
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
