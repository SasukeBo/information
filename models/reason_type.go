package models

import ()

// ReasonType 设备停机原因类型
type ReasonType struct {
	ID      int                 `orm:"auto;pk;column(id)"` // PKey 主键
	Name    string              `orm:"unique;index"`       // 类型名称
	Reasons []*DeviceStopReason `orm:"reverse(many)"`      // 设备停机原因
}
