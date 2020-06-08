package utils

import "github.com/dgrijalva/jwt-go"

type ApiPublic struct {
	Url string
	Type string
	Body []byte
}

type typeMethod struct {
	POST string
	GET string
	PUT string
	DELETE string
}
type CustomMessage struct {
	Key string
	Value string
}
type Claims struct {
	Payload []CustomMessage
	jwt.StandardClaims
}
type ProviceCityColombiaApi struct {
	Region         string `json:"region"`
	CodiDepartment string `json:"c_digo_dane_del_departamento"`
	DepartmentName string `json:"departamento"`
	CodiCity       string `json:"c_digo_dane_del_municipio"`
	City           string `json:"municipio"`
}
type Paginate struct {
	Page uint32
	NumberForPage uint32
	Filtered uint32
	Data interface{}
}