package structs

import "github.com/golang-jwt/jwt/v4"

type UserToken struct {
	jwt.StandardClaims
	Username string `json:"Username"`
}
