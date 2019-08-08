package models

// UserExtend 用户信息模型
type UserExtend struct {
	ID    int     `orm:"auto;pk;column(id)"`
	User  *User   `orm:"reverse(one);on_delete()"` // 用户删除时删除
	Name  string  `orm:"null"`                     // 真实姓名
	Email *string `orm:"unique;null"`
}
