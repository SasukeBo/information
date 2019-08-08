package public

import (
  "github.com/graphql-go/graphql"

  "github.com/SasukeBo/information/schema/fields"
  "github.com/SasukeBo/information/schema/resolvers/device"
  "github.com/SasukeBo/information/schema/types"
)

// DeviceChargeCreateField doc false
var DeviceChargeCreateField = &graphql.Field{
  Type: types.DeviceCharge,
  Args: graphql.FieldConfigArgument{
    "uuid":     fields.GenArg(graphql.String, "设备UUID", false),
    "userUUID": fields.GenArg(graphql.String, "指派人UUID", false),
  },
  Resolve: device.ChargeCreate,
}

// DeviceChargeDeleteField doc false
var DeviceChargeDeleteField = &graphql.Field{
  Type: graphql.String,
  Args: graphql.FieldConfigArgument{
    "id": fields.GenArg(graphql.Int, "设备指派ID", false),
  },
  Resolve: device.ChargeDelete,
}

// DeviceChargeUpdateField doc false
var DeviceChargeUpdateField = &graphql.Field{
  Type: types.DeviceCharge,
  Args: graphql.FieldConfigArgument{
    "id":       fields.GenArg(graphql.Int, "设备指派ID", false),
    "userUUID": fields.GenArg(graphql.String, "指派人UUID", false),
  },
  Resolve: device.ChargeUpdate,
}

// DeviceChargeGetField doc false
var DeviceChargeGetField = &graphql.Field{
  Type: types.DeviceCharge,
  Args: graphql.FieldConfigArgument{
    "id": fields.GenArg(graphql.Int, "设备负责关系ID", false),
  },
  Description: "通过id获取设备负责关系",
  Resolve:     device.ChargeGet,
}

// DeviceChargeListField doc false
var DeviceChargeListField = &graphql.Field{
  Type: graphql.NewList(types.DeviceCharge),
  Args: graphql.FieldConfigArgument{
    "userUUID":   fields.GenArg(graphql.String, "设备负责人uuid"),
    "deviceUUID": fields.GenArg(graphql.String, "设备uuid"),
  },
  Description: "通过负责人uuid或设备uuid获取设备负责关系列表",
  Resolve:     device.ChargeList,
}
