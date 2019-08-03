package device

import (
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
	"github.com/astaxie/beego/orm"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// Create 创建设备
func Create(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})

	dType := params.Args["type"].(string)
	dName := params.Args["name"].(string)
	token := uuid.New().String()
	description := params.Args["description"].(string)
	uuid := uuid.New().String()

	userUUID := rootValue["currentUserUUID"].(string)
	user := models.User{UUID: userUUID}
	if err := models.Repo.Read(&user, "uuid"); err == orm.ErrNoRows {
		return nil, utils.LogicError{
			Message: "user not found.",
		}
	}

	device := models.Device{
		Type:        dType,
		Name:        dName,
		Token:       token,
		UUID:        uuid,
		User:        &user,
		Description: description,
	}

	if _, err := models.Repo.Insert(&device); err != nil {
		return nil, err
	}
	// TODO: 为创建者分配所有权限

	return device, nil
}

// Bind 绑定设备Mac地址，需要权限验证
func Bind(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	token := params.Args["token"].(string)
	mac := params.Args["mac"].(string)

	currentUserUUID := rootValue["currentUserUUID"]
	if currentUserUUID == nil {
		return nil, utils.LogicError{
			Message: "user not authenticated.",
		}
	}

	user := models.User{UUID: currentUserUUID.(string)}

	if err := models.Repo.Read(&user, "uuid"); err != nil {
		return nil, utils.LogicError{
			Message: "user not found.",
		}
	}

	// TODO: 验证绑定设备的权限
	device := models.Device{Token: token}

	if err := models.Repo.Read(&device, "token"); err != nil {
		return nil, utils.LogicError{
			Message: "device not found.",
		}
	}

	// TODO: 设备状态
	device.Mac = mac
	device.Status = 1 // 已绑定

	if _, err := models.Repo.Update(&device, "mac"); err != nil {
		return nil, err
	}

	return device, nil
}

// Charge 指定设备负责人
func Charge(params graphql.ResolveParams) (interface{}, error) {
	// rootValue := params.Info.RootValue.(map[string]interface{})

	uuid := params.Args["uuid"].(string)
	userUUID := params.Args["userUuid"].(string)

	user := models.User{UUID: userUUID}
	if err := models.Repo.Read(&user, "uuid"); err != nil {
		return nil, utils.LogicError{
			Message: "user not found.",
		}
	}

	device := models.Device{UUID: uuid}
	if err := models.Repo.Read(&device, "uuid"); err != nil {
		return nil, utils.LogicError{
			Message: "device not found.",
		}
	}

	// TODO: 权限验证
	deviceCharge := models.DeviceCharge{
		User:   &user,
		Device: &device,
	}

	if _, err := models.Repo.Insert(&deviceCharge); err != nil {
		return nil, err
	}

	return deviceCharge, nil
}

// UNCharge 取消指定设备负责人
func UNCharge(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	deviceCharge := models.DeviceCharge{ID: id}
	if _, err := models.Repo.Delete(&deviceCharge); err != nil {
		return nil, err
	}

	return "ok", nil
}

// RECharge 重新指定设备负责人
func RECharge(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	userUUID := params.Args["userUuid"].(string)

	deviceCharge := models.DeviceCharge{ID: id}
	if err := models.Repo.Read(&deviceCharge); err != nil {
		return nil, err
	}

	user := models.User{UUID: userUUID}
	if err := models.Repo.Read(&user); err != nil {
		return nil, err
	}

	deviceCharge.User = &user
	if _, err := models.Repo.Update(&deviceCharge, "user_id"); err != nil {
		return nil, err
	}

	return deviceCharge, nil
}
