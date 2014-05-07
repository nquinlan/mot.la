package main

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func() string {
		return "mot.la - a home of <a href='http://scottmotte.com'>scottmotte</a>"
	})

	m.Run()
}
