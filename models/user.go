package models

import "time"

type User struct {
	ID   int64 `xorm:"pk autoincr"`
	Name string

	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func CreateUser(user *User) (err error) {
	_, err = x.Insert(user)
	return
}
