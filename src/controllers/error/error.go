package error

import (
	"github.com/kataras/iris/v12"
)

func Index303Get(ctx iris.Context) {
	ctx.ViewData("code", 303)
	ctx.ViewData("message", "User Not Permission")
	ctx.View("error.html")
	return
}
func Index404Get(ctx iris.Context)  {
	ctx.ViewData("code", 404)
	ctx.ViewData("message", "Not Found")
	ctx.View("error.html")
	return
}
func Index500Get(ctx iris.Context) {
	var message string
	if len(ctx.URLParam("error")) > 0 {
		message = ctx.URLParam("error")
	} else {
		message = "Oups something went wrong, try again"
	}
	ctx.ViewData("code", 500)
	ctx.ViewData("message", message)
	ctx.View("error.html")
	return
}