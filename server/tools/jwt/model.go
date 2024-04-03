package jwt

import "github.com/golang-jwt/jwt"

type TokenData struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}
