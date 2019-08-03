package custom

import (
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/logs"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// VMap interface to interface map type
type VMap map[interface{}]interface{}

func (m VMap) rMap(value int) interface{} {
	for k, v := range m {
		switch vt := v.(type) {
		case int:
			if value == vt {
				return k
			}
		}
	}
	return nil
}

// StatusMap 主要用于做graphql自定义的BaseStatus类型转换
var StatusMap = VMap{
	"default": models.BaseStatus.Default,
	"publish": models.BaseStatus.Publish,
	"block":   models.BaseStatus.Block,
	"deleted": models.BaseStatus.Deleted,
}

// BaseStatus is a custom graphql type
var BaseStatus = graphql.NewScalar(graphql.ScalarConfig{
	Name: "BaseStatus",
	Description: `BaseStatus is represent role current status,
	it convert string to int for DB,
	and convert int to string for output`,
	// Serialize 用于将 BaseStatus 类型值转换为 string 类型从gql接口输出
	Serialize: func(value interface{}) interface{} {
		rs, ok := value.(int)
		if !ok {
			return nil
		}
		key := StatusMap.rMap(rs)
		return key
	},
	// ParseValue 用于转换通过 variables 形式传递给 gquery 的值
	ParseValue: func(value interface{}) interface{} {
		key, ok := value.(string)
		if !ok {
			logs.Error("value is not a string")
			return nil
		}
		if value := StatusMap[key]; value != nil {
			return value
		}
		logs.Error("value is not a BaseStatus type")
		return nil
	},
	// ParseLiteral 用于转换 gql inline 变量值
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST.(type) {
		case *ast.StringValue:
			if value := StatusMap[valueAST.GetValue().(string)]; value != nil {
				return value
			}
		}
		logs.Error("value is not a BaseStatus type")
		return nil
	},
})
