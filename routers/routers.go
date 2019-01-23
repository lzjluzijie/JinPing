package routers

import (
	"github.com/lzjluzijie/JinPing/routers/api"
	"gopkg.in/macaron.v1"
)

func RegisterRouters(m *macaron.Macaron) {
	api.RegisterRouters(m)
}
