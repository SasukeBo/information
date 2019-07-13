package models

// UserProfile 用户信息模型
type UserProfile struct {
	ID       int    `orm:"auto;pk;column(id)"`
	UUID     string `orm:"column(uuid)"`
	RealName string
	User     *User  `orm:"reverse(one)"`
	Phone    string `orm:"null"`
	Email    string `orm:"null"`
}

// Insert doc false
func (up *UserProfile) Insert() error {
	_, err := repo.Insert(up)
	return err
}
