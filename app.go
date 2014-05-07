package main

import (
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "<html><body><p>mot.la</p><ul><li>a home of <a href='http://scottmotte.com'>scottmotte</a></li><li>I'm on <a href='http://twitter.com/scottmotte'>twitter</a></ul><p><img src='https://avatars2.githubusercontent.com/u/3848?s=460'></p></body></html>"
	})

	m.Run()
}
