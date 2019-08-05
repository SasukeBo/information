package device

import (
	"github.com/SasukeBo/information/models"
	"github.com/graphql-go/graphql"
)

// ParamCreate 设备参数创建
func ParamCreate(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	name := params.Args["name"].(string)
	sign := params.Args["sign"].(string)
	pType := params.Args["type"].(string)
	currentUserUUID := rootValue["currentUserUUID"].(string)

	user := models.User{UUID: currentUserUUID}
	if err := models.Repo.Read(&user, "uuid"); err != nil {
		return nil, err
	}

	deviceParam := models.DeviceParam{Name: name, Sign: sign, Type: pType, Author: &user}
	if _, err := models.Repo.Insert(&deviceParam); err != nil {
		return nil, err
	}

	return deviceParam, nil
}

// ParamUpdate 设备参数修改
func ParamUpdate(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	deviceParam := models.DeviceParam{ID: id}
	if err := models.Repo.Read(&deviceParam); err != nil {
		return nil, err
	}

	name := params.Args["name"]
	sign := params.Args["sign"]
	pType := params.Args["type"]

	if name != nil {
		deviceParam.Name = name.(string)
	}

	if sign != nil {
		deviceParam.Sign = sign.(string)
	}

	if pType != nil {
		deviceParam.Type = pType.(string)
	}

	if _, err := models.Repo.Update(&deviceParam); err != nil {
		return nil, err
	}

	return deviceParam, nil
}

// ParamDelete 设备参数删除
func ParamDelete(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	deviceParam := models.DeviceParam{ID: id}
	if _, err := models.Repo.Delete(&deviceParam); err != nil {
		return nil, err
	}

	return "ok", nil
}

// ParamGet ID获取设备参数
func ParamGet(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	deviceParam := models.DeviceParam{ID: id}
	if err := models.Repo.Read(&deviceParam); err != nil {
		return nil, err
	}

	return deviceParam, nil
}

// ParamList 根据条件获取设备参数列表
func ParamList(params graphql.ResolveParams) (interface{}, error) {
	qs := models.Repo.QueryTable("device_param")

	namePattern := params.Args["namePattern"]
	signPattern := params.Args["signPattern"]
	pType := params.Args["type"]
	userUUID := params.Args["userUUID"]

	if namePattern != nil {
		qs = qs.Filter("name__icontains", namePattern.(string))
	}

	if signPattern != nil {
		qs = qs.Filter("name__icontains", signPattern.(string))
	}

	if pType != nil {
		qs = qs.Filter("type", pType.(string))
	}

	if userUUID != nil {
		qs = qs.Filter("author__uuid", userUUID.(string))
	}

	var deviceParams []*models.DeviceParam

	if _, err := qs.All(&deviceParams); err != nil {
		return nil, err
	}

	return deviceParams, nil
}
