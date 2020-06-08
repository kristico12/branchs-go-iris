package model

import (
	_ "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-return/src/utils"
)

//------------------ database model connection -------------------------------------|
var shemaDatabase, _ = utils.LoadEnvironmentEnv("SHEMA_DATABASE")
var typeDatabase, _ = utils.LoadEnvironmentEnv("TYPE_DATABASE")
var userDatabase, _ = utils.LoadEnvironmentEnv("USER_DATABASE")
var passwordDatabase, _ = utils.LoadEnvironmentEnv("PASSWORD_DATABASE")
var sslMode, _ = utils.LoadEnvironmentEnv("SSL_MODE")
/*var nConxIdle, _ = utils.LoadEnvironmentEnv("MAX_IDLE_COX_DATABASE")
var nConx, _ = utils.LoadEnvironmentEnv("MAX_COX_DATABASE")
var maxCoxDatabase, _ = strconv.Atoi(nConx)
var maxIdleCoxDatabase, _ = strconv.Atoi(nConxIdle)*/
/*var hostDatabase, _ = utils.LoadEnvironmentEnv("HOST_DATABASE")
var portDatabase, _ = utils.LoadEnvironmentEnv("PORT_DATABASE")*/

//--------------------- interface conexion DB ----------------------------------|
func ConnectionDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Connect(typeDatabase, fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		userDatabase, passwordDatabase, shemaDatabase, sslMode))
	if err != nil {
		return nil, err
	}
	return db, nil
}