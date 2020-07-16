package routes

import (
	"github.com/kataras/iris/v12"
	"go-return/src/controllers/branch_office"
	viewError "go-return/src/controllers/error"
	"go-return/src/controllers/home"
	"go-return/src/controllers/login"
	"go-return/src/controllers/middlewares"
	"go-return/src/controllers/permission"
	"go-return/src/controllers/role"
)


func Routes() *iris.Application {
	app := iris.New()
	// route branch office
	app.Get("/", middlewares.Auth,middlewares.IsValidUser,branch_office.IndexGet)
	// route for dashboard
	dashboard := app.Party("/{id:uint64}", middlewares.Auth, middlewares.IsValidUser,middlewares.IsValidBranchOffice)
	dashboard.Get("/",home.IndexGet)
	dashboard.Get("/branch_office", middlewares.IsPermision,branch_office.IndexDashboardGet)
	dashboard.Get("/permission", middlewares.IsPermision,permission.IndexGet)
	dashboard.Get("/role", middlewares.IsPermision,role.IndexGet)
	//login
	routeLogin := app.Party("/login", middlewares.NotAuth)
	routeLogin.Get("/", login.IndexGet)
	routeLogin.Post("/", login.IndexPost)
	// route 303 permission
	app.Get("/permission303", viewError.Index303Get)
	// route 404 not found
	app.OnErrorCode(iris.StatusNotFound, viewError.Index404Get)
	// internal server error
	app.OnErrorCode(iris.StatusInternalServerError, viewError.Index500Get)
	//----------------------- routes for apis -----------------------------------|
	routeApi := app.Party("/api", middlewares.Auth, middlewares.IsValidUser)
	// route for branch offices
	routeApiBranchOffice := routeApi.Party("/branch_office", middlewares.IsPermision)
	routeApiBranchOffice.Post("/",branch_office.ApiPost)
	routeApiBranchOffice.Get("/",branch_office.ApiGet)
	routeApiBranchOffice.Put("/",branch_office.ApiPut)
	routeApiBranchOffice.Delete("/",branch_office.ApiDelete)
	// route for permission
	routeApiPermission := routeApi.Party("/permission", middlewares.IsPermision)
	routeApiPermission.Get("/", permission.ApiGet)
	routeApiPermission.Post("/", permission.ApiPost)
	routeApiPermission.Delete("/", permission.ApiDelete)
	routeApiPermission.Put("/", permission.ApiPut)
	// route for roles
	routeApiRole := routeApi.Party("/role", middlewares.IsPermision)
	routeApiRole.Get("/", role.ApiGet)
	routeApiRole.Post("/", role.ApiPost)
	routeApiRole.Put("/", role.ApiPut)
	return app
}
