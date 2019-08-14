package public

import (
  "github.com/graphql-go/graphql"

  "github.com/SasukeBo/information/schema/fields"
  "github.com/SasukeBo/information/schema/resolvers/device"
  "github.com/SasukeBo/information/schema/types"
)

// DeviceParamValueListField doc false
var DeviceParamValueListField = &graphql.Field{
  Type: graphql.NewList(types.DeviceParamValue),
  Args: graphql.FieldConfigArgument{
    "deviceParamID": fields.GenArg(graphql.Int, "参数ID", false),
    "beforeTime":    fields.GenArg(graphql.DateTime, "开始时间"),
    "afterTime":     fields.GenArg(graphql.DateTime, "结束时间"),
  },
  Resolve: device.ParamValueList,
}
