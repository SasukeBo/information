package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// DeviceStatus 设备状态
var DeviceStatus = struct {
	OffLine int // 离线
	Prod    int // 运行
	Stop    int // 在线但是未运行，就是停机
}{0, 1, 2}

// DeviceStatusLog 设备运行状态变更模型
type DeviceStatusLog struct {
	ID       int           `orm:"auto;pk;column(id)"`
	Status   int           // 设备运行状态
	Reasons  []*StopReason `orm:"null;rel(m2m)"` // 停机原因IDs
	Device   *Device       `orm:"rel(fk);on_delete()"`
	BeginAt  time.Time     `orm:"type(datetime)"`
	FinishAt time.Time     `orm:"type(datetime)"`
}

// LoadDevice _
func (dsl *DeviceStatusLog) LoadDevice() (*Device, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(dsl, "Device"); err != nil {
		return nil, Error{Message: "load related device failed.", OriErr: err}
	}
	return dsl.Device, nil
}
