package device

import (
  "github.com/SasukeBo/information/models"
  // "github.com/SasukeBo/information/utils"
  "github.com/graphql-go/graphql"
)

// ChargePrivGet 获取设备负责人权限
func ChargePrivGet(params graphql.ResolveParams) (interface{}, error) {
  id := params.Args["id"].(int)

  deviceChargeAbility := models.DeviceChargeAbility{ID: id}
  if err := models.Repo.Read(&deviceChargeAbility); err != nil {
    return nil, err
  }

  return deviceChargeAbility, nil
}

// ChargePrivList 根据设备负责人关系ID获取权限list
func ChargePrivList(params graphql.ResolveParams) (interface{}, error) {
  deviceChargeID := params.Args["deviceChargeID"].(int)
  qs := models.Repo.QueryTable("device_charge_ability")

  var chargePrivs []*models.DeviceChargeAbility
  if _, err := qs.Filter("device_charge_id", deviceChargeID).All(&chargePrivs); err != nil {
    return nil, err
  }

  return chargePrivs, nil
}

// ChargePrivCreate 为设备负责人添加权限
func ChargePrivCreate(params graphql.ResolveParams) (interface{}, error) {
  deviceChargeID := params.Args["deviceChargeID"].(int)
  privilegeID := params.Args["privilegeID"].(int)

  deviceCharge := models.DeviceCharge{ID: deviceChargeID}
  if err := models.Repo.Read(&deviceCharge); err != nil {
    return nil, err
  }

  privilege := models.Privilege{ID: privilegeID}
  if err := models.Repo.Read(&privilege); err != nil {
    return nil, err
  }

  deviceChargeAbility := models.DeviceChargeAbility{
    DeviceCharge: &deviceCharge,
    Privilege:    &privilege,
  }

  if _, err := models.Repo.Insert(&deviceChargeAbility); err != nil {
    return nil, err
  }

  return deviceChargeAbility, nil
}

// ChargePrivDelete 删除设备负责人的权限
func ChargePrivDelete(params graphql.ResolveParams) (interface{}, error) {
  id := params.Args["id"].(int)

  deviceChargeAbility := models.DeviceChargeAbility{ID: id}
  if _, err := models.Repo.Delete(&deviceChargeAbility); err != nil {
    return nil, err
  }

  return "ok", nil
}
