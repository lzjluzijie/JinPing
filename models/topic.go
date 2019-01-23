package models

import "time"

type Topic struct {
	ID        int64 `xorm:"pk autoincr"`
	UserID    int64
	URL       string
	Name      string
	CreatedAt time.Time `xorm:"created"`

	comments []Comment
}

func NewTopic(topic *Topic) (err error) {
	_, err = x.Insert(topic)
	return
}

func (topic *Topic) GetComments() (comments []Comment, err error) {
	if topic.comments != nil {
		comments = topic.comments
		return
	}

	comments = []Comment{}
	err = x.Where("ID = ?", topic.ID).Find(&comments)
	if err != nil {
		return
	}

	return
}
