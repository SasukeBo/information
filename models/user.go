package models

import (
	"time"
)

// User 用户模型
type User struct {
	ID         int    `orm:"auto;pk;column(id)"`
	UUID       string `orm:"column(uuid);unique;index"`
	Phone      string `orm:"unique"`
	Password   string
	AvatarURL  string      `orm:"column(avatar_url);null"`
	Role       *Role       `orm:"rel(fk);null"`
	UserExtend *UserExtend `orm:"rel(one)"`
	Status     BaseStatus  `orm:"default(0)"`
	CreatedAt  time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time   `orm:"auto_now;type(datetime)"`
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

	// 开始事务
	if err := repo.Begin(); err != nil {
		hasError <- true
		return err
	}

	if err := a.Role.GetByID(); err != nil {
		hasError <- true
		return err
	}
	if err := a.UserExtend.Insert(); err != nil {
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
