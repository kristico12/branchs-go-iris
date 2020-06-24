package utils

import (
	"bytes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"unicode"
)

//------------------- client HTTP HEADER --------------------------|
func ClientHttp(api ApiPublic) (*http.Request, error) {
	return http.NewRequest(api.Type, api.Url, bytes.NewBuffer(api.Body))
}

//--------------------------- loads environment variable ----------------------------------|
func LoadEnvironmentEnv(key string) (string, error) {
	err := godotenv.Load(HomeDir + "/.env")
	if err != nil {
		return "", err
	}
	return os.Getenv(key), nil
}

//----------------------------- encode y decode token ---------------------------------------|
func EncodeToken(userName string, password string) (string, error) {
	var payload []CustomMessage
	payload = append(payload,CustomMessage{"userName",userName},	CustomMessage{"password", password})
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Payload: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(privateKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func DecodeToken(token string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, fmt.Errorf("Token no valido")
	}
	return claims, nil
}

//------------------------------ funcs cookies -----------------------------------------------|
func SetCookie(data string, name string, duration time.Duration) *http.Cookie {
	c := &http.Cookie{}
	c.Name = name
	c.HttpOnly = true
	c.Value = url.QueryEscape(data)
	c.Expires = time.Now().Add(duration)
	return c
}
//-------------------------------- func clear special characters ------------------------------|
func isMn (r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
func ClearSpecialCharacteres(info string) (string, error) {
	b := make([]byte,len(info))
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	_, _, err := t.Transform(b,[]byte(info),true)
	if err != nil { return "", err}
	return strings.ToLower(string(b)), nil
}