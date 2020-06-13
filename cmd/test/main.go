package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	fmt.Println("---- Hello Kingdom! ************")

	app := iris.New()
	app.Get("/", func(ctx context.Context) {
		_, _ = ctx.WriteString("Hello Kingdom! ********")
	})
	_ = app.Run(iris.Addr(":8080"))
}
