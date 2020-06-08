package login

import (
	"github.com/kataras/iris/v12"
	"go-return/src/model"
	"go-return/src/utils"
	"time"
)

func IndexGet(ctx iris.Context)  {
	ctx.View("login.html")
	return
}
func IndexPost(ctx iris.Context)  {
	var auth model.UserAuth
	var password string
	getFormData := ctx.FormValues()
	for key, value := range getFormData {
		if key == "userName" { auth.Username = value[0] }
		if key == "password" { password = value[0] }
	}
	err := auth.Get("user_name=$1",auth.Username)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.ViewData("error", err.Error())
		ctx.View("login.html")
		return
	}
	err = auth.IsValidPassword(password)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.ViewData("error", err.Error())
		ctx.View("login.html")
		return
	}
	token, err := utils.EncodeToken(auth.Username, auth.Password)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.ViewData("error", err.Error())
		ctx.View("login.html")
		return
	}
	// set cookie
	ctx.SetCookie(utils.SetCookie(token,utils.CookieName,time.Hour * 24))
	ctx.Redirect("/", iris.StatusFound)
	return
}