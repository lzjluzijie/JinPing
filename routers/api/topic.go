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

func GetCommentsByURL(c *macaron.Context) {
	url := c.Query("url")

	comments, err := models.GetCommentsByURL(url)
	if err != nil {
		c.Error(500, err.Error())
		return
	}

	c.JSON(200, comments)
}
