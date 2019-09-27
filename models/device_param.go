package models

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"time"

	"github.com/SasukeBo/information/models/errors"
)

// DeviceParam 设备参数模型
type DeviceParam struct {
	Name      string    // 参数名称
	Sign      string    // 参数签名（标识），要求英文及数字组合的字符串
	Type      int       // 参数值类型，string？int？bool？
	Device    *Device   `orm:"rel(fk);on_delete()"`
	ID        int       `orm:"auto;pk;column(id)"`
	Author    *User     `orm:"rel(fk);null;on_delete(set_null)"` // 创建人，删除时置空
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// DeviceParamValueType 设备参数值类型
var DeviceParamValueType = struct {
	String  int
	Boolean int
	Integer int
	Float   int
}{0, 1, 2, 3}

// TableUnique 自定义唯一键
func (dp *DeviceParam) TableUnique() [][]string {
	return [][]string{
		[]string{"device_id", "sign"},
	}
}

// GetBy get device_param by col
func (dp *DeviceParam) GetBy(col string) error {
	if err := Repo.Read(dp, col); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Field:   col,
			Message: fmt.Sprintf("get device_param by %s error", col),
			OriErr:  err,
		}
	}

	return nil
}

// Get get device_param by id
func (dp *DeviceParam) Get() error {
	if err := Repo.Read(dp); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "get device_param error",
			OriErr:  err,
		}
	}

	return nil
}

// Delete delete device_param by id
func (dp *DeviceParam) Delete() error {
	if _, err := Repo.Delete(dp); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "delete device_param error",
			OriErr:  err,
		}
	}

	return nil
}

// Insert insert device_param
func (dp *DeviceParam) Insert() error {
	if _, err := Repo.Insert(dp); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "insert device_param error",
			OriErr:  err,
		}
	}

	return nil
}

// Update update device_param
func (dp *DeviceParam) Update(cols ...string) error {
	if _, err := Repo.Update(dp, cols...); err != nil {
		return errors.LogicError{
			Type:    "Model",
			Message: "update device_param error",
			OriErr:  err,
		}
	}

	return nil
}

// LoadAuthor _
func (dp *DeviceParam) LoadAuthor() (*User, error) {
	if _, err := Repo.LoadRelated(dp, "Author"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "device_param load author error",
			OriErr:  err,
		}
	}

	return dp.Author, nil
}

// LoadDevice _
func (dp *DeviceParam) LoadDevice() (*Device, error) {
	if _, err := Repo.LoadRelated(dp, "Device"); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "device_param load device error",
			OriErr:  err,
		}
	}

	return dp.Device, nil
}

// LoadDeviceParamValues _
func (dp *DeviceParam) LoadDeviceParamValues(params graphql.ResolveParams) ([]*DeviceParamValue, error) {
	qs := Repo.QueryTable("device_param_value").Filter("device_param_id", dp.ID).OrderBy("-created_at")

	if limit := params.Args["limit"]; limit != nil {
		qs = qs.Limit(limit)
	}

	if offset := params.Args["offset"]; offset != nil {
		qs = qs.Offset(offset)
	}

	if beforeTime := params.Args["beforeTime"]; beforeTime != nil {
		qs = qs.Filter("created_at__lt", beforeTime)
	}

	if afterTime := params.Args["afterTime"]; afterTime != nil {
		qs = qs.Filter("created_at__gt", afterTime)
	}

	var dpvs []*DeviceParamValue
	if _, err := qs.All(&dpvs); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "device_param load device_param_value error",
			OriErr:  err,
		}
	}

	return dpvs, nil
}

// ValidateAccess _
func (dp *DeviceParam) ValidateAccess(params graphql.ResolveParams, sign ...string) error {
	var device *Device
	var err error
	if device, err = dp.LoadDevice(); err != nil {
		return err
	}

	if err := device.ValidateAccess(params, sign...); err != nil {
		return err
	}

	return nil
}
