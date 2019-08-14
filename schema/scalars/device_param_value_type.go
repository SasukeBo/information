package scalars

import (
	"github.com/SasukeBo/information/models"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// DeviceParamValueTypeMap 主要用于做graphql自定义的DeviceParamValueType类型转换
var DeviceParamValueTypeMap = VMap{
	"String":  models.DeviceParamValueType.String,
	"Boolean": models.DeviceParamValueType.Boolean,
	"Integer": models.DeviceParamValueType.Integer,
	"Float":   models.DeviceParamValueType.Float,
}

// DeviceParamValueType is a scalars graphql type
var DeviceParamValueType = graphql.NewScalar(graphql.ScalarConfig{
	Name: "DeviceParamValueType",
	Description: `DeviceParamValueType is represent device param value type,
	can be "String"/"Boolean"/"Integer"/"Float",
  it convert string to int for DB,
  and convert int to string for output`,
	Serialize: func(value interface{}) interface{} {
		rs, ok := value.(int)
		if !ok {
			return nil
		}
		key := DeviceParamValueTypeMap.rMap(rs)
		return key
	},
	// ParseValue 用于转换通过 variables 形式传递给 query 的值
	ParseValue: func(value interface{}) interface{} {
		key, ok := value.(string)
		if !ok {
			return nil
		}
		if value := DeviceParamValueTypeMap[key]; value != nil {
			return value
		}
		return nil
	},
	// ParseLiteral 用于转换 gql inline 变量值
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST.(type) {
		case *ast.StringValue:
			if value := DeviceParamValueTypeMap[valueAST.GetValue().(string)]; value != nil {
				return value
			}
		}
		return nil
	},
})
