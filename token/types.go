package token

import "github.com/golang-jwt/jwt"

type CustomIdClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
