package models

// UserProfile 用户信息模型
type UserProfile struct {
	Uuid     string `orm:"pk"`
	RealName string
	User     *User  `orm:"reverse(one)"`
	Phone    string `orm:"null"`
	Email    string `orm:"null"`
}
