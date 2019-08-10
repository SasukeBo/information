package utils

import (
	"github.com/astaxie/beego/orm"
)

/*
EnumMap like Enum.map function in language elixir

Usage example:
```golang
	var lists = []orm.ParamsList{
		orm.ParamsList{1},
		orm.ParamsList{2},
		orm.ParamsList{3},
	}

	results := EnumMap(lists, func(item interface{}) interface{} {
		return item.([]interface{})[0]
	})

	fmt.Println(results)
```
=> [1 2 3]
*/
func EnumMap(lists []orm.ParamsList, f func(interface{}) interface{}) []interface{} {
	var result []interface{}

	for _, item := range lists {
		result = append(result, f(item))
	}

	return result
}
