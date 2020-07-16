package role

import (
	"database/sql"
	"fmt"
	"github.com/kataras/iris/v12"
	"go-return/src/model"
	"go-return/src/utils"
	"strconv"
)

func IndexGet(ctx iris.Context) {
	var role model.Role
	result, _ := role.Select("SELECT column_name FROM information_schema.columns\nWHERE table_catalog = 'Appointment' AND table_name = 'role'")
	for i, value := range result {
		utils.ListAttributesRoles[i].Key = value
	}
	ctx.ViewData("listAttributes", utils.ListAttributesRoles)
	ctx.StatusCode(iris.StatusOK)
	ctx.View("dashboard/role.html")
	return
}
func ApiGet(ctx iris.Context) {
	var (
		role model.Role
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
	// count filtered
	if count, err := role.Select(fmt.Sprintf("SELECT count(*) FROM permission %s",filter)); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	} else {
		if count, err := strconv.Atoi(count[0]); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"message": err.Error()})
			return
		} else {
			paginate.Filtered = uint32(count)
		}
	}
	if paginate.Data, err = role.Filter(fmt.Sprintf("%s ORDER BY %s OFFSET $1 ROWS FETCH FIRST $2 ROW ONLY", filter,titleOrder+" "+orderAscDesc),
		(paginate.Page*paginate.NumberForPage)-paginate.NumberForPage, paginate.NumberForPage); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(paginate)
	return
}
func ApiPost(ctx iris.Context) {
	var roleValidator model.RoleValidator
	var err error
	ctx.ReadJSON(&roleValidator)
	roleValidator.Validate()
	existError := roleValidator.Validate()
	if len(existError) > 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"errors": existError})
		return
	}
	var role = model.Role{Name: roleValidator.Name, Description: sql.NullString{roleValidator.Description, true}}
	err = role.Save()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"data": role})
	return
}
func ApiPut(ctx iris.Context) {
	var roleValidator model.RoleValidator
	var err error
	ctx.ReadJSON(&roleValidator)
	existError := roleValidator.Validate()
	if len(existError) > 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"errors": existError})
		return
	}
	var role = model.Role{Id: roleValidator.Id, Name: roleValidator.Name, Description: sql.NullString{roleValidator.Description, true} }
	err = role.Update()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": role})
	return
}
