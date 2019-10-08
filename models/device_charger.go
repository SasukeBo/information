package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// DeviceCharger 设备负责人
type DeviceCharger struct {
	ID         int `orm:"auto;pk;column(id)"`
	Name       string
	Phone      string
	Department string
	JobNumber  string
	Device     *Device   `orm:"rel(fk);on_delete()"` // 设备，删除时删除
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `orm:"auto_now;type(datetime)"`
}

// TableUnique 自定义唯一键
func (dc *DeviceCharger) TableUnique() [][]string {
	return [][]string{
		[]string{"device_id", "name", "phone"},
	}
}

// LoadDevice _
func (dc *DeviceCharger) LoadDevice() (*Device, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(dc, "Device"); err != nil {
		return nil, Error{Message: "load related device_charger failed", OriErr: err}
	}

	return dc.Device, nil
}
