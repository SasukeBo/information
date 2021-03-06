package schema

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

/*									base status
------------------------------------------------- */

// StatusMap 主要用于做graphql自定义的BaseStatus类型转换
var StatusMap = VMap{
	"default": models.BaseStatus.Default,
	"publish": models.BaseStatus.Publish,
	"block":   models.BaseStatus.Block,
	"deleted": models.BaseStatus.Deleted,
}

// BaseStatus is a scalars graphql type
var BaseStatus = graphql.NewScalar(graphql.ScalarConfig{
	Name: "BaseStatus",
	Description: `
#### 基础状态
可选值为：
- **default** 默认状态
- **publish** 发布状态
- **block** 禁用状态
- **deleted** 删除状态`,
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

/*									device status
------------------------------------------------- */

// DeviceStatusMap 主要用于做graphql自定义的DeviceStatus类型转换
var DeviceStatusMap = VMap{
	"prod":    models.DeviceStatus.Prod,
	"offline": models.DeviceStatus.OffLine,
	"stop":    models.DeviceStatus.Stop,
}

// DeviceStatus is a scalars graphql type
var DeviceStatus = graphql.NewScalar(graphql.ScalarConfig{
	Name: "DeviceStatus",
	Description: `
#### 设备状态
可选值：
- **prod** 生产状态
- **offline** 离线状态
- **stop** 停机状态`,
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

/*								privilege type
------------------------------------------------- */

// PrivTypeMap 主要用于做graphql自定义的PrivType类型转换
var PrivTypeMap = VMap{
	"default": models.PrivType.Default,
	"device":  models.PrivType.Device,
	"admin":   models.PrivType.Admin,
}

// PrivType is a scalars graphql type
var PrivType = graphql.NewScalar(graphql.ScalarConfig{
	Name: "PrivType",
	Description: `
#### 权限类型
可选值：
- **default** 默认权限
- **device** 设备相关权限
- **admin** 管理员相关权限`,
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
