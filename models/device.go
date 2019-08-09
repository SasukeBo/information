package models

import (
	"time"

	"github.com/SasukeBo/information/utils"
)

// Device 设备模型
type Device struct {
	Type        string    // 类型
	Name        string    // 设备名称
	Mac         string    // 设备Mac地址
	Token       string    `orm:"unique;index"`              // 设备Token，用于数据加密
	Status      int       `orm:"default(0)"`                // 基础状态
	ID          int       `orm:"auto;pk;column(id)"`        // PKey 主键
	UUID        string    `orm:"column(uuid);unique;index"` // 通用唯一标识符
	User        *User     `orm:"rel(fk)"`                   // 注册人
	Description string    `orm:"null"`                      // 描述
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`
}

// Get get device by id
func (d *Device) Get() error {
	if err := Repo.Read(d); err != nil {
		return utils.ORMError{
			Message: "device get error",
			OrmErr:  err,
		}
	}

	return nil
}

// LoadUser _
func (d *Device) LoadUser() (*User, error) {
	if _, err := Repo.LoadRelated(d, "User"); err != nil {
		return nil, utils.ORMError{
			Message: "devcice load related user error",
			OrmErr:  err,
		}
	}

	return d.User, nil
}
