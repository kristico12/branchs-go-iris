package home

import "github.com/kataras/iris/v12"

func IndexGet(ctx iris.Context)  {
	ctx.View("dashboard/home.html")
	return
}
