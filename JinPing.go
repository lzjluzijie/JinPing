package main

import (
	"net/http"

	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.New()

	err := http.ListenAndServe(":80", m)
	if err != nil {
		panic(err)
	}
}
