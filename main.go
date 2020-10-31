package main

import "github.com/kataras/iris/v12"

func main(){
	app := iris.New()

	app.Get("/", func(ctx iris.Context){
		ctx.ServeFile("index.html",false)
	})

	app.Run(iris.Addr(":5000"))
}