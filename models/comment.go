package models

import (
	"errors"
	"fmt"
	"time"
)

type Comment struct {
	ID      int64 `xorm:"pk autoincr"`
	UserID  int64
	TopicID int64
	Content string

	CreatedAt time.Time `xorm:"created"`
}

func NewComment(comment *Comment) (err error) {
	_, err = x.Insert(comment)
	return
}

func GetCommentsByURL(url string) (comments []Comment, err error) {
	topic := &Topic{
		URL: url,
	}

	has, err := x.Get(topic)
	if err != nil {
		return
	}
	if !has {
		err = errors.New(fmt.Sprintf("cannot find topic: %s", url))
		return
	}

	comments = []Comment{}
	err = x.Where("TopicID = ?", topic.ID).Find(&comments)
	if err != nil {
		return
	}
	return
}
