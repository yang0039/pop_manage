package middleware

import (
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
	jwt.StandardClaims
}
