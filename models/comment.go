package models

import "time"

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

func GetCommentsByTopic(topicID int64) (comments []Comment, err error) {
	comments = []Comment{}
	err = x.Where("ID = ?", topicID).Find(&comments)
	if err != nil {
		return
	}
	return
}
