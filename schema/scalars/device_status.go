package scalars

import (
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/logs"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// DeviceStatusMap 主要用于做graphql自定义的DeviceStatus类型转换
var DeviceStatusMap = VMap{
	"prod":    models.DeviceStatus.Prod,
	"offline": models.DeviceStatus.OffLine,
	"stop":    models.DeviceStatus.Stop,
}

// DeviceStatus is a scalars graphql type
var DeviceStatus = graphql.NewScalar(graphql.ScalarConfig{
	Name: "DeviceStatus",
	Description: `DeviceStatus is represent device current status,
  it convert string to int for DB,
  and convert int to string for output`,
	Serialize: func(value interface{}) interface{} {
		rs, ok := value.(int)
		if !ok {
			return nil
		}
		key := DeviceStatusMap.rMap(rs)
		return key
	},
	// ParseValue 用于转换通过 variables 形式传递给 gquery 的值
	ParseValue: func(value interface{}) interface{} {
		key, ok := value.(string)
		if !ok {
			logs.Error("value is not a string")
			return nil
		}
		if value := DeviceStatusMap[key]; value != nil {
			return value
		}
		logs.Error("value is not a DeviceStatus type")
		return nil
	},
	// ParseLiteral 用于转换 gql inline 变量值
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST.(type) {
		case *ast.StringValue:
			if value := DeviceStatusMap[valueAST.GetValue().(string)]; value != nil {
				return value
			}
		}
		logs.Error("value is not a DeviceStatus type")
		return nil
	},
})
