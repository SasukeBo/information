package models

import (
	"time"
	// "github.com/SasukeBo/information/utils"
)

// DeviceStatus 设备状态
var DeviceStatus = struct {
	Prod    int // 生产
	Stop    int // 停机
	OffLine int // 离线
}{0, 1, 2}

// DeviceStatusLog 设备运行状态变更模型
type DeviceStatusLog struct {
	Status   int       // 设备运行状态
	ID       int       `orm:"auto;pk;column(id)"`
	Device   *Device   `orm:"rel(fk);on_delete()"`
	ChangeAt time.Time `orm:"auto_now_add;type(datetime)"`
}
