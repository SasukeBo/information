package models

import (
	"time"
)

// Enum the User status

// User 用户模型
type User struct {
	Password    string
	ID          int          `orm:"auto;pk;column(id)"`
	UUID        string       `orm:"column(uuid);unique;index"`
	Account     string       `orm:"unique"`
	UserProfile *UserProfile `orm:"rel(one)"`
	Role        *Role        `orm:"rel(fk)"`
	Status      BaseStatus   `orm:"default(0)"`
	CreatedAt   time.Time    `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time    `orm:"auto_now;type(datetime)"`
}

// Insert 创建用户
func (a *User) Insert() error {
	hasError := make(chan bool, 1)

	defer func() {
		if has := <-hasError; has {
			repo.Rollback()
			return
		}
		repo.Commit()
	}()

	if err := repo.Begin(); err != nil {
		hasError <- true
		return err
	}
	if err := a.Role.GetByID(); err != nil {
		hasError <- true
		return err
	}
	if err := a.UserProfile.Insert(); err != nil {
		hasError <- true
		return err
	}
	if _, err := repo.Insert(a); err != nil {
		hasError <- true
		return err
	}
	hasError <- false
	return nil
}
