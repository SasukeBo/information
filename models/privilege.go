package models

// PrivType 权限类型
var PrivType = struct {
	Default int
	Admin   int
	Device  int
}{0, 1, 2}

// Privilege 权限模型
type Privilege struct {
	Name     string // 权限名称
	ID       int    `orm:"auto;pk;column(id)"`
	Sign     string // 权限签名
	PrivType int    `orm:"default(0)"` // 权限类型
}

// TableUnique 自定义唯一键
func (p *Privilege) TableUnique() [][]string {
	return [][]string{
		[]string{"sign", "priv_type"},
	}
}
