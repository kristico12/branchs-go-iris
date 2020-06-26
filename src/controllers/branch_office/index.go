package branch_office

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"go-return/src/model"
	"go-return/src/utils"
	"strconv"
)

func IndexGet(ctx iris.Context) {
	var branchOffice model.BranchOffice
	branchOffices, err := branchOffice.Filter("")
	if err != nil {
		ctx.ViewData("error", err.Error())
	}
	if len(branchOffices) == 0 {
		ctx.ViewData("error", "No hay oficinas Registradas")
	}
	for i := range branchOffices {
		branchOffices[i].MascarateHours("03:04:05 PM", "03:04:05 PM")
		branchOffices[i].MascarateProvinceCity()
	}

	ctx.ViewData("branchOffices", branchOffices)
	ctx.View("branch-office.html")
	return
}
func IndexDashboardGet(ctx iris.Context) {
	var branchOffice model.BranchOffice
	listAttributes, _ := branchOffice.Select("SELECT column_name FROM information_schema.columns\nWHERE table_catalog = 'Appointment' AND table_name = 'branch_office'")
	for i, value := range listAttributes {
		utils.ListAttributesBranchOfficess[i].Key = value
	}
	ctx.ViewData("listAttributes", utils.ListAttributesBranchOfficess)
	ctx.View("dashboard/branch-office.html")
	return
}
func ApiPost(ctx iris.Context) {
	var validatorBranchOffice model.BranchOfficeValidator
	err := ctx.ReadJSON(&validatorBranchOffice)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	existError := validatorBranchOffice.Validate()
	if len(existError) > 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"errors": existError})
		return
	}
	var branchOffice = model.BranchOffice{City: validatorBranchOffice.City, Province: validatorBranchOffice.Province,
		Address: validatorBranchOffice.Address, CheckInTime: validatorBranchOffice.CheckInTime, ExitTime: validatorBranchOffice.ExitTime}
	err = branchOffice.Save()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"data": branchOffice})
	return
}
func ApiGet(ctx iris.Context) {
	var (
		branchOffices model.BranchOffice
		paginate      utils.Paginate
		titleOrder    string
		orderAscDesc  string
		filterCity []string
		filterProvince []string
		err           error
		strIn string
	)
	// asing id
	branchOffices.Id, _ = strconv.ParseUint(ctx.URLParam("branchOffice"), 10, 64)
	// FILTER
	if len(ctx.URLParam("filter")) > 0 {
		filterCity = model.FilterProvinceCity(ctx.URLParam("filter"),"city")
		filterProvince = model.FilterProvinceCity(ctx.URLParam("filter"), "province")
		if len(filterCity) > 0 && len(filterProvince) > 0 {
			strIn = "AND (city IN (?) OR province IN (?))"
		} else if len(filterCity) > 0 && len(filterProvince) == 0 {
			strIn = "AND city IN (?)"
		} else if len(filterCity) == 0 && len(filterProvince) > 0 {
			strIn = "AND province IN (?)"
		} else { strIn = "" }
	}
	// assign ORDER BY
	if len(ctx.URLParam("titleOrder")) > 0 {
		titleOrder = ctx.URLParam("titleOrder")
	} else { titleOrder = "id" }
	// define ASC or DESC
	if len(ctx.URLParam("orderAscDesc")) > 0 {
		orderAscDesc = ctx.URLParam("orderAscDesc")
	} else { orderAscDesc = "ASC" }
	// execute query count
	count, err := branchOffices.In(
		fmt.Sprintf("SELECT * FROM branch_office WHERE id<>%d %s",
			branchOffices.Id, strIn),filterCity,filterProvince)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	paginate.Filtered = uint32(len(count))
	// execute query paginate
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
	result, err := branchOffices.In(
		fmt.Sprintf("SELECT * FROM branch_office WHERE id<>%d %s ORDER BY %s OFFSET %d ROWS FETCH FIRST %d ROW ONLY",
			branchOffices.Id,strIn,titleOrder+" "+orderAscDesc,
			(paginate.Page*paginate.NumberForPage)-paginate.NumberForPage, paginate.NumberForPage),filterCity,filterProvince)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	for i := range result {
		result[i].MascarateHours("03:04:05 PM", "03:04:05 PM")
		result[i].MascarateProvinceCity()
	}
	paginate.Data = result
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(paginate)
	return
}
func ApiPut(ctx iris.Context)  {
	var validatorBranchOffice model.BranchOfficeValidator
	err := ctx.ReadJSON(&validatorBranchOffice)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	existError := validatorBranchOffice.Validate()
	if len(existError) > 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"errors": existError})
		return
	}
	id, _ := strconv.ParseUint(validatorBranchOffice.Id, 10, 64)
	var branchOffice = model.BranchOffice{ Id: id, City: validatorBranchOffice.City, Province: validatorBranchOffice.Province,
		Address: validatorBranchOffice.Address, CheckInTime: validatorBranchOffice.CheckInTime, ExitTime: validatorBranchOffice.ExitTime}
	err = branchOffice.Update()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": branchOffice})
	return
}
func ApiDelete(ctx iris.Context) {
	var branchOffice model.BranchOffice
	err := ctx.ReadJSON(&branchOffice)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	err = branchOffice.Delete()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": branchOffice})
	return
}