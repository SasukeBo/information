package models

// UserExtend 用户信息模型
type UserExtend struct {
	ID    int    `orm:"auto;pk;column(id)"`
	User  *User  `orm:"reverse(one)"`
	Name  string `orm:"null"`
	Email string `orm:"unique;null"`
}

// Insert doc false
func (up *UserExtend) Insert() error {
	_, err := repo.Insert(up)
	return err
}
