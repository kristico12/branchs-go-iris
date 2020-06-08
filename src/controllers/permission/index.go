package permission

import (
	"github.com/kataras/iris/v12"
)

func IndexPost(ctx iris.Context) {
	ctx.StatusCode(iris.StatusAccepted)
	ctx.JSON(iris.Map{ "data": "hola"})
	return
}
