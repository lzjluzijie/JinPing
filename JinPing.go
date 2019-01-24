package main

import (
	"log"
	"net/http"

	"github.com/lzjluzijie/JinPing/routers"

	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Renderer())

	routers.RegisterRouters(m)

	log.Println("ok")

	err := http.ListenAndServe("127.0.0.1:80", m)
	if err != nil {
		panic(err)
	}
}
