package api

import "gopkg.in/macaron.v1"

func RegisterRouters(m *macaron.Macaron) {
	m.Group("/api", func() {
		m.Post("/comment", NewComment)
		m.Post("/user", NewUser)
	})
}
