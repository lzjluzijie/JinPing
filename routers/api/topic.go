package api

import (
	"strconv"

	"github.com/lzjluzijie/JinPing/models"
	"gopkg.in/macaron.v1"
)

func NewTopic(c *macaron.Context) {
	userID, err := strconv.ParseInt(c.Query("userID"), 10, 64)
	if err != nil {
		c.Error(500, err.Error())
		return
	}
	name := c.Query("name")
	url := c.Query("url")

	topic := &models.Topic{
		UserID: userID,
		Name:   name,
		URL:    url,
	}

	err = models.CreateTopic(topic)
	if err != nil {
		c.Error(500, err.Error())
		return
	}

	c.JSON(200, topic)
	return
}

func GetCommentsByTopic(c *macaron.Context) {
	topicID, err := strconv.ParseInt(c.Params("topicID"), 10, 64)
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
