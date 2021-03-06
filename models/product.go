package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Product 产品模型
type Product struct {
	ID               int           `orm:"auto;pk;column(id)"`
	Name             string        `orm:"unique"`                           // 产品名称
	Token            string        `orm:"unique;index"`                     // 产品Token
	ProductorContact string        `orm:"null"`                             // 生产负责人联系方式
	Productor        string        `orm:"null"`                             // 生产负责人
	Register         *User         `orm:"rel(fk);null;on_delete(set_null)"` // 注册人
	DetectItems      []*DetectItem `orm:"reverse(many)"`                    // 产品检测项
	FinishTime       time.Time     `orm:"null;type(datetime)"`              // 交货日期
	Total            int           `orm:"null"`                             // 指标总量
	OrderNum         string        `orm:"null"`                             // 订单编号
	Customer         string        `orm:"null"`                             // 订货方
	CustomerContact  string        `orm:"null"`                             // 联系方式
	CreatedAt        time.Time     `orm:"auto_now_add;type(datetime)"`
	UpdatedAt        time.Time     `orm:"auto_now;type(datetime)"`
}

// LoadDetectItem _
func (p *Product) LoadDetectItem() ([]*DetectItem, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(p, "DetectItems"); err != nil {
		return nil, nil
	}

	return p.DetectItems, nil
}

// LoadUser _
func (p *Product) LoadUser() (*User, error) {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(p, "Register"); err != nil {
		return nil, nil
	}

	return p.Register, nil
}

// GetDevices _
func (p *Product) GetDevices(db orm.Ormer) ([]*Device, error) {
	var devices []*Device

	sql := `
	SELECT d.id, d.address, d.name, d.number, d.remote_ip, d.status, d.token, d.type, d.user_id, d.created_at, d.updated_at
	FROM device d
	JOIN device_product_ship dps ON d.id = dps.device_id
	WHERE dps.product_id = ?
	`

	if _, err := db.Raw(sql, p.ID).QueryRows(&devices); err != nil {
		return devices, Error{Message: "get product devices failed.", OriErr: err}
	}

	return devices, nil
}
