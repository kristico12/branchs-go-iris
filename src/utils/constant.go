package utils

import (
	"os"
)
//------------ paht dir proyect ---------------------------------------------------------|
var HomeDir,_ = os.Getwd()
//------------------------- urls publics apis -------------------------------------------|
var UrlApiCityProvinceColombia, _ = LoadEnvironmentEnv("URL_CITY_PROVINCE_COLOMBIA")
//----------------------- type fecth --------------------------------------------------|
var TypeMethod = typeMethod {"POST", "GET", "PUT", "DELETE"}
const ContentTypeJson = "application/json"
// ------------------ private key token ------------------------------------|
var privateKey, _ = LoadEnvironmentEnv("ENCODE_TOKEN")
// ------------------ cookie name -----------------------|
var CookieName, _ = LoadEnvironmentEnv("COKIE_NAME")
//------------------ list attributes for Model ---------------|
var ListAttributesBranchOfficess = []CustomMessage{
	{"","Id" }, {"","Ciudad"}, {"","Provincia"},
	{"","Direccion"}, {"","Hora Entrada"}, {"","Hora Salida"}}
var ListAttributesPermission = []CustomMessage{
	{"","Id" }, {"","Nombre"}}

//------------------ NumberForPage ----------------|
const NumberForPage = uint32(7)