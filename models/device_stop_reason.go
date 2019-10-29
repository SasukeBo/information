package models

import (
// "github.com/astaxie/beego/orm"
)

// DeviceStopReason 设备停机原因
type DeviceStopReason struct {
	ID      int         `orm:"auto;pk;column(id)"`  // PKey 主键
	Type    *ReasonType `orm:"rel(fk);on_delete()"` // 原因类型
	Content string      // 原因内容
	Code    string      `orm:"unique;index;"` // 停机代码
}
