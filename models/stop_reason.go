package models

import (
	"github.com/astaxie/beego/orm"
)

// StopReason 设备停机原因
// 为了兼容PLC，准确收集停机原因，采用事先注册设备停机原因的方式。
// 设备发送停机原因时以约定好的指定字节指定位设置0或1表示该原因是否存在。
// 约定好固定的某几位是所有设备统一的停机原因，例如：人为停机。
// 其余位为自定义注册的异常停机原因。
// 当发生人为停机时，将忽略同时发生的异常停机原因。
type StopReason struct {
	ID        int     `orm:"auto;pk;column(id)"`  // PKey 主键
	Device    *Device `orm:"rel(fk);on_delete()"` // 所属设备
	WordIndex int     // 第几个字节
	BitPos    int     // 第几个bit位
	Content   string  // 原因内容
}

// TableUnique 自定义唯一键
func (sr *StopReason) TableUnique() [][]string {
	return [][]string{
		[]string{"device_id", "word_index", "bit_pos"},
	}
}

// LoadDevice _
func (sr *StopReason) LoadDevice() (*Device, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(sr, "Device"); err != nil {
		return nil, nil
	}

	return sr.Device, nil
}
