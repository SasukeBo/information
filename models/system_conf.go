package models

import (
	// "github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

// 配置项值类型
var systemConfValueType = struct {
	String int
	Int    int
	Float  int
	Time   int
}{0, 1, 2, 3}

// SystemConf 系统配置项
type SystemConf struct {
	ID        int       `orm:"auto;pk;column(id)"` // PKey 主键
	Name      string    // 配置项名称
	Value     string    // 配置项值
	ValueType int       `orm:"default(0)"` // 配置项值类型
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

// ParseValue 根据配置项值类型返回配置项的值
// 类型解析失败时返回此类型的零值。
func (sc *SystemConf) ParseValue() interface{} {
	switch sc.ValueType {
	case systemConfValueType.String:
		return sc.Value
	case systemConfValueType.Int:
		v, err := strconv.Atoi(sc.Value)
		if err != nil {
			return 0
		}
		return v
	case systemConfValueType.Float:
		v, err := strconv.ParseFloat(sc.Value, 64)
		if err != nil {
			return float64(0)
		}
		return v
	case systemConfValueType.Time:
		v, err := time.Parse(time.RFC3339, sc.Value)
		if err != nil {
			var t time.Time
			return t
		}
		return v
	}

	return nil
}
