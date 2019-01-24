package api

import (
	"github.com/lzjluzijie/JinPing/models"
	"gopkg.in/macaron.v1"
)

func NewUser(c *macaron.Context) {
	name := c.Query("name")

	user := &models.User{
		Name: name,
	}
	err := models.CreateUser(user)
	if err != nil {
		c.Error(500, err.Error())
		return
	}

	c.JSON(200, name)
}
