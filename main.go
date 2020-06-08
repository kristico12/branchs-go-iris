// package
package main

// imports
import (
	"github.com/kataras/iris/v12"
	"go-return/src/model"
	"go-return/src/routes"
	"go-return/src/utils"
)

func main() {
	serverApp := routes.Routes()
	//---------------------------- validate conection .enviroment -----------------------------------------|
	_, err := utils.LoadEnvironmentEnv("load")
	if err != nil {
		serverApp.Logger().Fatalf("load enviroment variable")
	}
	// ---------------------------- Sync connection DB ---------------------------------------------------|
	err = model.Migrate()
	if err != nil {
		serverApp.Logger().Fatalf("error sync database ", err)
	}
	// add template comfig using django views
	tmpl := iris.Django("./src/views", ".html")
	tmpl.Reload(true)  // Enable re-build on local template files changes
	// register the all views engine
	serverApp.RegisterView(tmpl)
	// add route public public files
	serverApp.HandleDir("/public", "./src/public")
	// Run application
	serverApp.Run(iris.Addr(":8000"))
}
