package device

import (
  // "github.com/astaxie/beego/logs"
  "github.com/astaxie/beego/orm"
  "github.com/google/uuid"
  "github.com/graphql-go/graphql"

  "github.com/SasukeBo/information/models"
  "github.com/SasukeBo/information/utils"
)

// Get 获取设备
func Get(params graphql.ResolveParams) (interface{}, error) {
  uuid := params.Args["uuid"].(string)

  device := models.Device{UUID: uuid}

  if err := models.Repo.Read(&device, "uuid"); err != nil {
    return nil, err
  }

  return device, nil
}

// List 获取设备列表
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
    return nil, err
  }

  return devices, nil
}

// Create 创建设备
func Create(params graphql.ResolveParams) (interface{}, error) {
  rootValue := params.Info.RootValue.(map[string]interface{})

  dType := params.Args["type"].(string)
  dName := params.Args["name"].(string)
  token := utils.GenRandomToken(8)
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

// Update 更新设备
func Update(params graphql.ResolveParams) (interface{}, error) {
  uuid := params.Args["uuid"].(string)
  dType := params.Args["type"]
  name := params.Args["name"]
  status := params.Args["status"]
  description := params.Args["description"]

  device := models.Device{UUID: uuid}
  if err := models.Repo.Read(&device, "uuid"); err != nil {
    return nil, err
  }

  if dType != nil {
    device.Type = dType.(string)
  }

  if name != nil {
    device.Name = name.(string)
  }

  if status != nil {
    device.Status = status.(int)
  }

  if description != nil {
    device.Description = description.(string)
  }

  if _, err := models.Repo.Update(&device); err != nil {
    return nil, err
  }

  return device, nil
}

// Delete 更新设备
func Delete(params graphql.ResolveParams) (interface{}, error) {
  uuid := params.Args["uuid"].(string)

  device := models.Device{UUID: uuid}
  if err := models.Repo.Read(&device, "uuid"); err != nil {
    return nil, err
  }

  if _, err := models.Repo.Delete(&device); err != nil {
    return nil, err
  }

  return "ok", nil
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
