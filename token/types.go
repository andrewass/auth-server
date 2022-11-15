package token

import "github.com/golang-jwt/jwt"

type CustomIdClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type TypeHint string

const (
	AccessToken  TypeHint = "AccessToken"
	RefreshToken          = "RefreshToken"
)
