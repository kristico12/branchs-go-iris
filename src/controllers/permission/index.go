package permission

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"go-return/src/model"
	"go-return/src/utils"
	"strconv"
)

func IndexGet(ctx iris.Context) {
	var permission model.Permission
	result, _ := permission.Select("SELECT column_name FROM information_schema.columns\nWHERE table_catalog = 'Appointment' AND table_name = 'permission'")
	for i, value := range result {
		utils.ListAttributesPermission[i].Key = value
	}
	ctx.ViewData("listAttributes", utils.ListAttributesPermission)
	ctx.View("dashboard/permission.html")
	return
}
func ApiGet(ctx iris.Context) {
	var (
		permission model.Permission
		paginate      utils.Paginate
		filter string
		titleOrder    string
		orderAscDesc  string
		err           error
	)
	// FILTER
	if len(ctx.URLParam("filter")) > 0 {
		filter = fmt.Sprintf("WHERE name LIKE '%%%s%%'",ctx.URLParam("filter"))
	} else { filter = "" }
	// assign ORDER BY
	if len(ctx.URLParam("titleOrder")) > 0 {
		titleOrder = ctx.URLParam("titleOrder")
	} else { titleOrder = "id" }
	// define ASC or DESC
	if len(ctx.URLParam("orderAscDesc")) > 0 {
		orderAscDesc = ctx.URLParam("orderAscDesc")
	} else { orderAscDesc = "ASC" }
	if len(ctx.URLParam("page")) > 0 {
		if convInt, err := strconv.Atoi(ctx.URLParam("page")); err == nil {
			paginate.Page = uint32(convInt)
		} else {
			paginate.Page = uint32(1)
		}
	} else {
		paginate.Page = uint32(1)
	}
	paginate.NumberForPage = utils.NumberForPage
	result, err := permission.Filter("$1 ORDER BY $2 OFFSET $3 ROWS FETCH FIRST $4 ROW ONLY",filter,titleOrder+" "+orderAscDesc,
		(paginate.Page*paginate.NumberForPage)-paginate.NumberForPage, paginate.NumberForPage)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(result)
	return
}
