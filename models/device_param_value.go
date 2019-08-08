package models

import (
	"time"
	// "github.com/SasukeBo/information/utils"
)

// DeviceParamValue 设备参数值模型
type DeviceParamValue struct {
	Value       string       // 参数值
	ID          int          `orm:"auto;pk;column(id)"`
	DeviceParam *DeviceParam `orm:"rel(fk);on_delete()"`
	CreatedAt   time.Time    `orm:"auto_now_add;type(datetime)"`
}
