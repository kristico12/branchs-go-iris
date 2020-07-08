package model

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"go-return/src/utils"
	"strings"
	"time"
)

// constant
var provinceCityColombia []utils.ProviceCityColombiaApi

type PermissionValidator struct {
	Id string `json:"id"`
	Name string `json:"name" validate:"required,min=3,max=100,existSpacing"`
}
type RoleValidator struct {
	Id uint64 `json:"id"`
	Name        string `json:"name" validate:"required,min=3,max=30,existSpacing"`
	Description string `json:"description" validate:"max=400"`
}
type UserAuthValidator struct {
	Username string         `json:"userName" validate:"required,min=5,max=30"`
	Password string         `json:"password" validate:"required,min=6,max=500"`
	IdRol    *RoleValidator `json:"idRol" validate:"required,dive,required"`
}
type UserValidator struct {
	Name           string             `json:"name" validate:"required,min=3,max=30"`
	LastName       string             `json:"lastName" validate:"required,min=5,max=50"`
	Identification string             `json:"identification" validate:"required,min=5,max=20,numeric"`
	Birthday       *time.Time         `json:"birthday"`
	IdUserAuth     *UserAuthValidator `json:"idUserAuth" validate:"dive"`
}
type ClientValidator struct {
	Address string         `json:"address" validate:"min=3,max=80"`
	IdUser  *UserValidator `json:"idUser" validate:"required,dive,required"`
}
type EmployeeValidator struct {
	Salary float32        `json:"salary" validate:"required,min=1,numeric"`
	IdUser *UserValidator `json:"idUser" validate:"required,dive,required"`
}
type ServiceValidator struct {
	Name        string  `json:"name" validate:"required,min=5,max=30"`
	Description string  `json:"description" validate:"max=100"`
	Price       float32 `json:"price" validate:"required,min=1,numeric"`
}
type EmployeeServiceValidator struct {
	IdEmployee []*EmployeeValidator `json:"idEmployee" validate:"required,dive,required"`
	IdService  []*ServiceValidator  `json:"idService" validate:"required,dive,required"`
}
type BranchOfficeValidator struct {
	Id string `json:"id"`
	City        string `json:"city" validate:"required,min=1,max=50,numeric,customInvalidCodeCity=Province"`
	Province    string `json:"province" validate:"required,min=1,max=50,numeric,customInvalidCodeProvince"`
	Address     string `json:"address" validate:"required,min=5"`
	CheckInTime string `json:"checkInTime" validate:"required,min=8,max=8"`
	ExitTime    string `json:"exitTime" validate:"required,min=8,max=8,customValidateHour=CheckInTime"`
}
type BranchOfficeServiceValidator struct {
	IdBranchOffice []*BranchOfficeValidator `json:"idBranchOffice" validate:"required,dive,required"`
	IdService      []*ServiceValidator      `json:"idService" validate:"required,dive,required"`
}
type AppointmentValidator struct {
	TotalPrice     float32                  `json:"totalPrice" validate:"required,min=1,numeric"`
	StartTime      *time.Time               `json:"startTime" validate:"required"`
	EndTime        *time.Time               `json:"endTime" validate:"required"`
	IdClient       []*ClientValidator       `json:"idClient" validate:"required,dive,required"`
	IdEmployee     []*EmployeeValidator     `json:"idEmployee" validate:"required,dive,required"`
	IdBranchOffice []*BranchOfficeValidator `json:"idBranchOffice" validate:"required,dive,required"`
}
type ServiceAppointmentValidator struct {
	IdAppointment []*AppointmentValidator `json:"idAppointment" validate:"required,dive,required"`
	IdService     []*ServiceValidator     `json:"idService" validate:"required,dive,required"`
}
type EmployeeAppointmenValidator struct {
	IdAppointment []*AppointmentValidator `json:"idAppointment" validate:"required,dive,required"`
	IdEmployee    []*EmployeeValidator    `json:"idEmployee" validate:"required,dive,required"`
}

//-------------- custom func validators ----------------------------------|
func customValidateHour(fl validator.FieldLevel) bool {
	var exitHour = strings.SplitN(fl.Field().String(), ":", 3)
	var entryHour = strings.SplitN(fl.Parent().FieldByName(fl.Param()).String(), ":", 3)
	var entryHourDuration, _ = time.ParseDuration(fmt.Sprintf("%sh%sm%ss", entryHour[0], entryHour[1], entryHour[2]))
	var exitHourDuration, _ = time.ParseDuration(fmt.Sprintf("%sh%sm%ss", exitHour[0], exitHour[1], exitHour[2]))
	var isValid = true
	if exitHourDuration <= entryHourDuration {
		isValid = false
	}
	return isValid
}
func customInvalidCodeProvince(fl validator.FieldLevel) bool {
	var codeProvince = fl.Field().String()
	var isValid = false
	// search value
	for _, value := range provinceCityColombia {
		if codeProvince == value.CodiDepartment {
			isValid = true
		}
	}
	return isValid
}
func customInvalidCodeCity(fl validator.FieldLevel) bool {
	var selectedCity utils.ProviceCityColombiaApi
	var codeCity = fl.Field().String()
	var codeProvince = fl.Parent().FieldByName(fl.Param()).String()
	var isValid = false
	// search value
	for _, value := range provinceCityColombia {
		if codeCity == value.CodiCity {
			selectedCity = value
		}
	}
	if selectedCity.CodiDepartment == codeProvince {
		isValid = true
	}
	return isValid
}
func existSpacing(fl validator.FieldLevel) bool {
	var isValid = false
	var i = strings.Index(fl.Field().String(), " ")
	if i == -1 {
		isValid = true
	}
	return isValid
}

//-------------- interfaces validator ------------------------------------|
type ValidatorsInfo interface {
	Validate() []utils.CustomMessage
}

func validate(err interface{}) []utils.CustomMessage {
	var AllErrorMessage []utils.CustomMessage
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			AllErrorMessage = append(AllErrorMessage, utils.CustomMessage{
				Key: strcase.ToLowerCamel(err.Field()), Value: "error: " + err.Tag()})
		}
	}
	return AllErrorMessage
}

func (self PermissionValidator) Validate() []utils.CustomMessage {
	v := validator.New()
	_ = v.RegisterValidation("existSpacing", existSpacing)
	return validate(v.Struct(self))
}
func (self RoleValidator) Validate() []utils.CustomMessage {
	v := validator.New()
	_ = v.RegisterValidation("existSpacing", existSpacing)
	return validate(v.Struct(self))
}
func (self UserAuthValidator) Validate() []utils.CustomMessage {
	err := validator.New().Struct(self)
	return validate(err)
}
func (self UserValidator) Validate() []utils.CustomMessage {
	err := validator.New().Struct(self)
	return validate(err)
}
func (self EmployeeValidator) Validate() []utils.CustomMessage {
	err := validator.New().Struct(self)
	return validate(err)
}
func (self BranchOfficeValidator) Validate() []utils.CustomMessage {
	v := validator.New()
	_ = v.RegisterValidation("customValidateHour", customValidateHour)
	_ = v.RegisterValidation("customInvalidCodeProvince", customInvalidCodeProvince)
	_ = v.RegisterValidation("customInvalidCodeCity", customInvalidCodeCity)
	return validate(v.Struct(self))
}
func init() {
	var call utils.ApiPublic
	// fecht api
	call.Url = utils.UrlApiCityProvinceColombia
	dataApiProvince, _ := call.GetDataApi(utils.TypeMethod.GET, utils.ContentTypeJson)
	// convert json to strung value
	json.Unmarshal([]byte(dataApiProvince), &provinceCityColombia)

}
