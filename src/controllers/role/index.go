package role

import "github.com/kataras/iris/v12"

func IndexGet(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	ctx.View("dashboard/role.html")
	return
}
