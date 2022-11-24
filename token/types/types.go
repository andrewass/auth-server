package types

import "github.com/golang-jwt/jwt"

type CustomIdClaims struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	jwt.StandardClaims
}

type TypeHint string

const (
	AccessToken  TypeHint = "AccessToken"
	RefreshToken          = "RefreshToken"
)
