package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
)

//-------------------------- interface model Actions ------------------------------------------|
type ModelCrud interface {
	Save() error
	Update() error
	Delete()
	Get() error
	Filter()
	Select()
}
//-------------------------------- all methods SAVE ----------------------------------------|
func (self *BranchOffice) Save() error  {
	Db, err := ConnectionDatabase()
	if err != nil { return err }
	tx := Db.MustBegin()
	tx.NamedExec("INSERT INTO branch_office (city, province, address, check_in_time, exit_time) VALUES (:city,:province,:address,:check_in_time,:exit_time)",self)
	err = tx.Commit()
	defer Db.Close()
	if err != nil { return err }
	return nil
}
func (self *Permission) Save() error {
	Db, err := ConnectionDatabase()
	if err != nil { return err }
	tx := Db.MustBegin()
	tx.NamedExec("INSERT INTO permission (name) VALUES (:name)",self)
	err = tx.Commit()
	defer Db.Close()
	if err != nil { return err }
	return nil
}
//-------------------------------- all methods GET ----------------------------------------|
func (self *UserAuth) Get(where string, args ...interface{}) error {
	Db, err := ConnectionDatabase()
	if err != nil { return err }
	err = Db.Get(self, fmt.Sprintf("SELECT * FROM user_auth WHERE %s",where),args...)
	defer Db.Close()
	if err != nil { return err }
	return nil
}
func (self *BranchOffice) Get(where string, args ...interface{}) error {
	Db, err := ConnectionDatabase()
	if err != nil { return err }
	err = Db.Get(self, fmt.Sprintf("SELECT * FROM branch_office WHERE %s",where),args...)
	defer Db.Close()
	if err != nil { return err }
	return nil
}
//-------------------------------- all methods UPDATE ------------------------------------------|
func (self *BranchOffice) Update() error {
	Db, err := ConnectionDatabase()
	if err != nil { return err }
	tx := Db.MustBegin()
	tx.NamedExec("UPDATE branch_office SET city=:city, province=:province, address=:address, check_in_time=:check_in_time, exit_time=:exit_time WHERE id=:id",self)
	err = tx.Commit()
	defer Db.Close()
	if err != nil { return err }
	return nil
}
func (self *Permission) Update() error {
	Db, err := ConnectionDatabase()
	if err != nil { return err }
	tx := Db.MustBegin()
	tx.NamedExec("UPDATE permission SET name=:name WHERE id=:id",self)
	err = tx.Commit()
	defer Db.Close()
	if err != nil { return err }
	return nil
}
//------------------------------- all methods DELETE --------------------------------------------|
func (self BranchOffice) Delete() error {
	Db, err := ConnectionDatabase()
	if err != nil { return err }
	tx := Db.MustBegin()
	tx.NamedExec("DELETE FROM branch_office WHERE id=:id",self)
	err = tx.Commit()
	defer Db.Close()
	if err != nil { return err }
	return nil
}
func (self Permission) Delete() error {
	Db, err := ConnectionDatabase()
	if err != nil { return err }
	tx := Db.MustBegin()
	tx.NamedExec("DELETE FROM permission WHERE id=:id",self)
	err = tx.Commit()
	defer Db.Close()
	if err != nil { return err }
	return nil
}
//-------------------------------- all methods Filter -------------------------------------------|
func (self BranchOffice) Filter(where string, args ...interface{}) ([]BranchOffice, error) {
	var result []BranchOffice
	Db, err := ConnectionDatabase()
	if err != nil { return nil, err }
	err = Db.Select(&result, fmt.Sprintf("SELECT * FROM branch_office %s", where), args...)
	defer Db.Close()
	if err != nil { return nil, err }
	return result, nil
}
func (self Permission) Filter(where string, args ...interface{}) ([]Permission, error) {
	var result []Permission
	Db, err := ConnectionDatabase()
	if err != nil { return nil, err }
	err = Db.Select(&result, fmt.Sprintf("SELECT * FROM permission %s", where), args...)
	defer Db.Close()
	if err != nil { return nil, err }
	return result, nil
}

//------------------------------- all Filter IN -----------------------------------------------------|
func (self BranchOffice) In(filterIn string, args ...interface{}) ([]BranchOffice, error) {
	var (
		existSlice bool = false
		err        error
		Db         *sqlx.DB
		newSlice []interface{}
	)
	for i, arg := range args {
		switch reflect.TypeOf(arg).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(arg)
			if s.Len() > 0 {
				existSlice = true
			} else { args[i] = nil }
		}
	}
	for _, item := range args {
		if item != nil { newSlice = append(newSlice, item) }
	}
	args = newSlice
	var result []BranchOffice
	Db, err = ConnectionDatabase()
	if err != nil { return nil, err }
	if existSlice {
		filterIn, args, err = sqlx.In(filterIn,args...)
		if err != nil { return nil, err }
		filterIn = Db.Rebind(filterIn)
	} else { args = nil }
	Db.Select(&result,filterIn,args...)
	if err != nil { return nil, err }
	defer Db.Close()
	return result, nil
}
//------------------------------- all methods CustomQuerys -----------------------------------------------|
func (self UserAuth) Select(customQuery string, args ...interface{}) ([]string, error) {
	var result []string
	Db, err := ConnectionDatabase()
	if err != nil { return nil, err }
	err = Db.Select(&result, customQuery, args...)
	defer Db.Close()
	if err != nil { return nil, err }
	return result, nil
}
func (self BranchOffice) Select(customQuery string, args ...interface{}) ([]string, error) {
	var result []string
	Db, err := ConnectionDatabase()
	if err != nil { return nil, err }
	err = Db.Select(&result, customQuery, args...)
	defer Db.Close()
	if err != nil { return nil, err }
	return result, nil
}
func (self Permission) Select(customQuery string, args ...interface{}) ([]string, error) {
	var result []string
	Db, err := ConnectionDatabase()
	if err != nil { return nil, err }
	err = Db.Select(&result, customQuery, args...)
	defer Db.Close()
	if err != nil { return nil, err }
	return result, nil
}