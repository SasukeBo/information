package scalars

import (
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/logs"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// PrivTypeMap 主要用于做graphql自定义的PrivType类型转换
var PrivTypeMap = VMap{
	"default": models.PrivType.Default,
	"device":  models.PrivType.Device,
	"admin":   models.PrivType.Admin,
}

// PrivType is a scalars graphql type
var PrivType = graphql.NewScalar(graphql.ScalarConfig{
	Name: "PrivType",
	Description: `PrivType is represent privilege type,
  it convert string to int for DB,
  and convert int to string for output`,
	Serialize: func(value interface{}) interface{} {
		rs, ok := value.(int)
		if !ok {
			return nil
		}
		key := PrivTypeMap.rMap(rs)
		return key
	},
	// ParseValue 用于转换通过 variables 形式传递给 gquery 的值
	ParseValue: func(value interface{}) interface{} {
		key, ok := value.(string)
		if !ok {
			logs.Error("value is not a string")
			return nil
		}
		if value := PrivTypeMap[key]; value != nil {
			return value
		}
		logs.Error("value is not a PrivType type")
		return nil
	},
	// ParseLiteral 用于转换 gql inline 变量值
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST.(type) {
		case *ast.StringValue:
			if value := PrivTypeMap[valueAST.GetValue().(string)]; value != nil {
				return value
			}
		}
		logs.Error("value is not a PrivType type")
		return nil
	},
})
