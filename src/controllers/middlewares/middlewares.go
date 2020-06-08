package middlewares

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"go-return/src/model"
	"go-return/src/utils"
	"strconv"
	"strings"
)

func Auth(ctx iris.Context) {
	cookie := ctx.GetCookie(utils.CookieName)
	token, err := utils.DecodeToken(cookie)
	if err != nil {
		ctx.RemoveCookie(utils.CookieName, iris.CookieCleanPath)
		ctx.Redirect("/login")
		return
	}
	for _, value := range token.Payload {
		ctx.Values().Set(value.Key, value.Value)
	}
	ctx.Next()
}
func NotAuth(ctx iris.Context) {
	cookie := ctx.GetCookie(utils.CookieName)
	_, err := utils.DecodeToken(cookie)
	if err != nil {
		ctx.RemoveCookie(utils.CookieName, iris.CookieCleanPath)
		ctx.Next()
	}
	ctx.Redirect("/")
	return
}
func IsValidUser(ctx iris.Context) {
	var userAuth model.UserAuth
	userAuth.Username = ctx.Values().GetString("userName")
	userAuth.Password = ctx.Values().GetString("password")
	err := userAuth.Get("user_name=$1 AND password=$2", userAuth.Username, userAuth.Password)
	ctx.Values().Set("id", userAuth.Id)
	if err != nil {
		ctx.RemoveCookie(utils.CookieName, iris.CookieCleanPath)
		ctx.Redirect("/login")
		return
	}
	ctx.Next()
}
func IsPermision(ctx iris.Context) {
	var contentType = ctx.GetHeader("Content-Type")
	var permision string
	var isPermission = false
	permision = ctx.Request().URL.String()[strings.LastIndex(ctx.Request().URL.String(), "/")+1:]
	var userAuth model.UserAuth
	userAuth.Id, _ = strconv.ParseUint(ctx.Values().GetString("id"), 10, 64)
	result, err := userAuth.Select("SELECT p.name FROM user_auth\nLEFT JOIN role r on user_auth.id_role = r.id\nLEFT JOIN permission_role pr on r.id = pr.id_role\nLEFT JOIN permission p on pr.id_permission = p.id\nWHERE user_auth.id = $1", userAuth.Id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		if contentType == "application/json" {
			ctx.JSON(iris.Map{"message": err.Error()})
		} else {
			ctx.Request().URL.Query().Set("error", err.Error())
		}
		return
	}
	for _, value := range result {
		if value == "all" || value == permision {
			isPermission = true
		}
	}
	if isPermission {
		ctx.Next()
	} else {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.Redirect("/permission303")
		return
	}
}
func IsValidBranchOffice(ctx iris.Context) {
	var contentType = ctx.GetHeader("Content-Type")
	idBranchOffice := ctx.Params().GetUint64Default("id", 0)
	if idBranchOffice == uint64(0) {
		ctx.StatusCode(iris.StatusInternalServerError)
		if contentType == "application/json" {
			ctx.JSON(iris.Map{"message": "Invalid ID"})
		} else {
			ctx.URLParamDefault("error", "Invalid ID")
		}
		return
	} else {
		var branchOffice model.BranchOffice
		err := branchOffice.Get("id=$1", idBranchOffice)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			if contentType == "application/json" {
				ctx.JSON(iris.Map{"message": err.Error()})
			} else {
				ctx.URLParamDefault("error", err.Error())
			}
			return
		}
		branchOffice.MascarateProvinceCity()
		// validates permission
		var userAuth model.UserAuth
		userAuth.Id, _ = strconv.ParseUint(ctx.Values().GetString("id"), 10, 64)
		result, err := userAuth.Select("SELECT p.name FROM user_auth\nLEFT JOIN role r on user_auth.id_role = r.id\nLEFT JOIN permission_role pr on r.id = pr.id_role\nLEFT JOIN permission p on pr.id_permission = p.id\nWHERE user_auth.id = $1", userAuth.Id)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			if contentType == "application/json" {
				ctx.JSON(iris.Map{"message": err.Error()})
			} else {
				ctx.Request().URL.Query().Set("error", err.Error())
			}
			return
		}
		ctx.ViewData("permission", result)
		ctx.ViewData("branch", branchOffice)
		ctx.ViewData("urlHome", fmt.Sprintf("/%d",idBranchOffice))
		ctx.Next()
	}
}
