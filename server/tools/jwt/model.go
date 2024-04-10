package jwt

import "github.com/golang-jwt/jwt"

type TokenData struct {
	Id uint32 `json:"id"`
	jwt.StandardClaims
}
