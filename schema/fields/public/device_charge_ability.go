package public

import (
  "github.com/graphql-go/graphql"

  "github.com/SasukeBo/information/schema/fields"
  "github.com/SasukeBo/information/schema/resolvers/device"
  "github.com/SasukeBo/information/schema/types"
)

// DeviceChargePrivGetField doc false
var DeviceChargePrivGetField = &graphql.Field{
  Type: types.DeviceChargeAbility,
  Args: graphql.FieldConfigArgument{
    "id": fields.GenArg(graphql.Int, "负责人权限ID", false),
  },
  Resolve: device.ChargePrivGet,
}

// DeviceChargePrivListField doc false
var DeviceChargePrivListField = &graphql.Field{
  Type: graphql.NewList(types.DeviceChargeAbility),
  Args: graphql.FieldConfigArgument{
    "deviceChargeID": fields.GenArg(graphql.Int, "负责人关系ID", false),
  },
  Description: "根据设备负责人关系获取权限list",
  Resolve:     device.ChargePrivList,
}

// DeviceChargePrivCreateField doc false
var DeviceChargePrivCreateField = &graphql.Field{
  Type: types.DeviceChargeAbility,
  Args: graphql.FieldConfigArgument{
    "deviceChargeID": fields.GenArg(graphql.Int, "负责人关系ID", false),
    "privilegeID":    fields.GenArg(graphql.Int, "权限ID", false),
  },
  Description: "为设备负责人添加权限",
  Resolve:     device.ChargePrivCreate,
}

// DeviceChargePrivDeleteField doc false
var DeviceChargePrivDeleteField = &graphql.Field{
  Type: graphql.String,
  Args: graphql.FieldConfigArgument{
    "id": fields.GenArg(graphql.Int, "设备负责人权限ID", false),
  },
  Description: "删除设备负责人的权限",
  Resolve:     device.ChargePrivDelete,
}
