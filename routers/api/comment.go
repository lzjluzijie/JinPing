package api

import (
	"strconv"

	"github.com/lzjluzijie/JinPing/models"
	"gopkg.in/macaron.v1"
)

func NewComment(c *macaron.Context) {
	userID, err := strconv.ParseInt(c.Query("userID"), 10, 64)
	if err != nil {
		c.Error(500, err.Error())
		return
	}
	topicID, err := strconv.ParseInt(c.Query("topicID"), 10, 64)
	if err != nil {
		c.Error(500, err.Error())
		return
	}
	content := c.Query("content")

	comment := &models.Comment{
		UserID:  userID,
		TopicID: topicID,
		Content: content,
	}

	err = models.NewComment(comment)
	if err != nil {
		c.Error(500, err.Error())
		return
	}

	c.JSON(200, comment)
}

func GetComments(c *macaron.Context) {
	topicID, err := strconv.ParseInt(c.Query("topicID"), 10, 64)
	if err != nil {
		c.Error(500, err.Error())
		return
	}

	comments, err := models.GetCommentsByTopic(topicID)
	if err != nil {
		c.Error(500, err.Error())
		return
	}

	c.JSON(200, comments)
}
