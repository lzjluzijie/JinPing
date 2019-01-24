package models

import "time"

type Topic struct {
	ID     int64 `xorm:"pk autoincr"`
	UserID int64
	Name   string
	URL    string

	CreatedAt time.Time `xorm:"created"`
}

func CreateTopic(topic *Topic) (err error) {
	_, err = x.Insert(topic)
	return
}
