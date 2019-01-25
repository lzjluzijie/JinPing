package routers

import (
	"github.com/lzjluzijie/JinPing/routers/api"
	"gopkg.in/macaron.v1"
)

func RegisterRouters(m *macaron.Macaron) {
	m.Use(macaron.Static("public"))

	api.RegisterRouters(m)
}
