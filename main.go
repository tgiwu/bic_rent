package main

import (
	"fmt"

	"github.com/kataras/iris/v12"

)

func main () {
	fmt.Printf("new project")
	app:= iris.New()

	app.Use(iris.Compression)

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("Helow  <strong>%s</strong>", "world")
	})

	app.Listen(":8081")
}