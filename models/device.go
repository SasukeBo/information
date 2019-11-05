package models

import (
	"fmt"
	"github.com/SasukeBo/information/utils"
	"github.com/astaxie/beego/orm"
	"regexp"
	"strconv"
	"time"
)

var accErr = Error{Message: "only device register can make this operation!"}

// Device 设备模型
type Device struct {
	ID       int    `orm:"auto;pk;column(id)"` // PKey 主键
	Type     string // 类型
	Name     string // 设备名称
	Address  string `orm:"null"`                             // 设备地址
	Number   string `orm:"null"`                             // 设备编号
	RemoteIP string `orm:"null;column(remote_ip)"`           // 接入IP
	Token    string `orm:"unique;index"`                     // 设备Token，用于数据加密
	Status   int    `orm:"default(0)"`                       // 离线状态
	User     *User  `orm:"rel(fk);null;on_delete(set_null)"` // 注册人
	// DeviceChargers []*DeviceCharger `orm:"reverse(many)"`
	CreatedAt      time.Time `orm:"auto_now_add;type(datetime)"`
	StatusChangeAt time.Time `orm:"auto_now;type(datetime)"`
	UpdatedAt      time.Time `orm:"auto_now;type(datetime)"`
}

// LoadUser _
func (d *Device) LoadUser() (*User, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(d, "User"); err != nil {
		return nil, Error{Message: "load related user failed.", OriErr: err}
	}

	return d.User, nil
}

// ValidateAccess _
func (d *Device) ValidateAccess(u *User) error {
	if d.User.ID != u.ID {
		return accErr
	}

	return nil
}

// DeviceClassOEE 设备班次生产指标
type DeviceClassOEE struct {
	Activation float64
	Yield      float64
	OEE        float64
}

// GetCurrentClassOEE 获取设备当班生产指标数据
func (d *Device) GetCurrentClassOEE() (interface{}, error) {
	o := orm.NewOrm()
	var (
		response   DeviceClassOEE
		dayBegin   time.Time
		nightBegin time.Time
		beginTime  time.Time
		ok         bool
	)

	now := utils.TruncateTime(time.Now().UTC())

	scDay := &SystemConf{Name: "白班交接时间"}
	if err := o.Read(scDay, "name"); err != nil {
		return nil, Error{Message: fmt.Sprintf("system config for %s not found.", scDay.Name)}
	}
	if dayBegin, ok = scDay.ParseValue().(time.Time); !ok {
		return nil, Error{Message: fmt.Sprintf("%s is not a time value", scDay.Name)}
	}
	dayBegin = utils.TruncateTime(dayBegin)

	scNight := &SystemConf{Name: "夜班交接时间"}
	if err := o.Read(scNight, "name"); err != nil {
		return nil, Error{Message: fmt.Sprintf("system config for %s not found.", scNight.Name)}
	}
	if nightBegin, ok = scNight.ParseValue().(time.Time); !ok {
		return nil, Error{Message: fmt.Sprintf("%s is not a time value", scNight.Name)}
	}
	nightBegin = utils.TruncateTime(nightBegin)

	totalInsideDurationSQL := `
	SELECT SUM(dsl.finish_at - dsl.begin_at) AS duration
	FROM device_status_log AS dsl
	WHERE device_id = ? AND status = ? AND dsl.begin_at > ? AND dsl.finish_at < ?
	`

	type duration struct {
		Duration string
	}

	var (
		totalInsideProdDuration duration        // 完全发生在时间区间内的 生产 总时长
		totalInsideStopDuration duration        // 完全发生在时间区间内的 停机 总时长
		frontThroughLog         DeviceStatusLog // 起止时间跨越交接班时间的状态日志
		currentLog              DeviceStatusLog // 当前未结束的状态日志
		prodSeconds             float64         // 生产状态持续秒数
		stopSeconds             float64         // 停机状态持续秒数
		good                    int             // 良品数
		bad                     int             // 不良品数
	)

	if now.After(dayBegin) && now.Before(nightBegin) {
		beginTime = dayBegin // 当前时间为白班，起始时间为早班交接班时间
	} else {
		beginTime = nightBegin // 当前时间为晚班，起始时间为晚班交接班时间
	}

	o.Raw(totalInsideDurationSQL, d.ID, DeviceStatus.Prod, beginTime, now).QueryRow(&totalInsideProdDuration)
	o.Raw(totalInsideDurationSQL, d.ID, DeviceStatus.Stop, beginTime, now).QueryRow(&totalInsideStopDuration)

	o.QueryTable("device_status_log").Filter("device_id", d.ID).Filter("begin_at__lt", beginTime).Filter("finish_at__gt", beginTime).One(&frontThroughLog)
	o.QueryTable("device_status_log").Filter("device_id", d.ID).OrderBy("-id").Limit(1).One(&currentLog)

	prodSeconds = parseDuration(totalInsideProdDuration.Duration)
	stopSeconds = parseDuration(totalInsideStopDuration.Duration)

	if frontThroughLog.ID != 0 {
		seconds := frontThroughLog.FinishAt.Sub(beginTime).Seconds()
		switch frontThroughLog.Status {
		case DeviceStatus.Prod:
			prodSeconds += seconds
		case DeviceStatus.Stop:
			stopSeconds += seconds
		}
	}

	if currentLog.ID != 0 {
		seconds := now.Sub(currentLog.BeginAt).Seconds()
		switch currentLog.Status {
		case DeviceStatus.Prod:
			prodSeconds += seconds
		case DeviceStatus.Stop:
			stopSeconds += seconds
		}
	}

	response.Activation = prodSeconds / (prodSeconds + stopSeconds)

	var piCount struct {
		Count int
	}

	productInsCountSQL := `
	SELECT COUNT(*) FROM product_ins pi
	JOIN device_product_ship dps ON pi.device_product_ship_id = dps.id
	WHERE dps.device_id = ? AND pi.created_at > ? AND pi.created_at < ? AND pi.qualified = ?
	`

	if err := o.Raw(productInsCountSQL, d.ID, beginTime, now, true).QueryRow(&piCount); err == nil {
		good = piCount.Count
	}

	if err := o.Raw(productInsCountSQL, d.ID, beginTime, now, false).QueryRow(&piCount); err == nil {
		bad = piCount.Count
	}

	response.Yield = float64(good) / float64(good+bad)

	return &response, nil
}

///////////////////////////////////////// private /////////////////////////////////////////

// parseDuration 解析数据库查询的时长字符串，返回整型秒数
func parseDuration(dbduration string) float64 {
	seconds := 0
	pattern := `^(\d*)( days? )?(\d+):(\d{2}):(\d{2})(\.\d*)?$`
	reg := regexp.MustCompile(pattern)
	matches := reg.FindStringSubmatch(dbduration)

	if len(matches) > 1 {
		days, err := strconv.Atoi(matches[1])
		if err == nil {
			seconds += days * 24 * 60 * 60
		}
	}

	if len(matches) > 3 {
		hours, err := strconv.Atoi(matches[3])
		if err == nil {
			seconds += hours * 60 * 60
		}
	}

	if len(matches) > 4 {
		minutes, err := strconv.Atoi(matches[4])
		if err == nil {
			seconds += minutes * 60
		}
	}

	if len(matches) > 5 {
		second, err := strconv.Atoi(matches[5])
		if err == nil {
			seconds += second
		}
	}

	return float64(seconds)
}
