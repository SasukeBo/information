package resolver

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"

	// "github.com/astaxie/beego/logs"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/utils"
)

// GetDevice 获取设备
func GetDevice(params graphql.ResolveParams) (interface{}, error) {
	uuid := params.Args["uuid"].(string)

	device := models.Device{UUID: uuid}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	return device, nil
}

// GetDeviceByToken 获取设备
func GetDeviceByToken(params graphql.ResolveParams) (interface{}, error) {
	token := params.Args["token"].(string)

	device := models.Device{Token: token}
	if err := device.GetBy("token"); err != nil {
		return nil, err
	}

	return device, nil
}

// ListDevice 获取设备列表
func ListDevice(params graphql.ResolveParams) (interface{}, error) {
	cond := models.NewCond()

	if pattern := params.Args["pattern"]; pattern != nil {
		subCond := models.NewCond()
		cond = cond.AndCond(subCond.Or("type__icontains", pattern).Or("name__icontains", pattern))
	}

	if status := params.Args["status"]; status != nil {
		cond = cond.And("status", status)
	}

	if isRegister := params.Args["isRegister"]; isRegister != nil {
		if isRegister.(bool) {
			user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
			cond = cond.And("user_id", user.ID)
		}
	}
	qs := models.Repo.QueryTable("device").SetCond(cond).OrderBy("-created_at")

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
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "get device list error.",
			OriErr:  err,
		}
	}

	return struct {
		Total   int64
		Devices []*models.Device
	}{cnt, devices}, nil
}

// CreateDevice 创建设备
func CreateDevice(params graphql.ResolveParams) (interface{}, error) {
	// 验证用户是否有创建设备的权限
	if err := utils.ValidateAccess(&params, "device_c", models.PrivType.Default); err != nil {
		return nil, err
	}
	rootValue := params.Info.RootValue.(map[string]interface{})

	device := models.Device{}
	user := rootValue["currentUser"].(models.User)
	device.User = &user

	dType := params.Args["type"].(string)
	if err := utils.ValidateStringEmpty(dType, "type"); err != nil {
		return nil, err
	}
	device.Type = dType

	dName := params.Args["name"].(string)
	if err := utils.ValidateStringEmpty(dName, "name"); err != nil {
		return nil, err
	}
	device.Name = dName
	if description := params.Args["description"]; description != nil {
		device.Description = description.(string)
	}

	if address := params.Args["address"]; address != nil {
		device.Address = address.(string)
	}

	count := params.Args["count"].(int)
	devices := []models.Device{}
	for i := 0; i < count; i++ {
		device.Token = utils.GenRandomToken(8)
		device.UUID = uuid.New().String()
		devices = append(devices, device)
	}

	successNums, err := models.Repo.InsertMulti(count, devices)
	if err != nil {
		return nil, err
	}

	return successNums, nil
}

// UpdateDevice 更新设备
func UpdateDevice(params graphql.ResolveParams) (interface{}, error) {
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	device := models.Device{UUID: params.Args["uuid"].(string)}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	if err := device.ValidateAccess(&user); err != nil {
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

// DeleteDevice 更新设备
func DeleteDevice(params graphql.ResolveParams) (interface{}, error) {
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	device := models.Device{UUID: params.Args["uuid"].(string)}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	if err := device.ValidateAccess(&user); err != nil {
		return nil, err
	}

	if err := device.Delete(); err != nil {
		return nil, err
	}

	return "ok", nil
}

// RelatedLoad _
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

// CountDeviceStatus _
func CountDeviceStatus(params graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}
