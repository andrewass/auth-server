package types

import (
	"github.com/golang-jwt/jwt"
)

type CustomIdClaims struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	jwt.StandardClaims
}

type TypeHint string
type GrantType string

const (
	AccessToken  TypeHint = "AccessToken"
	RefreshToken TypeHint = "RefreshToken"

	AuthCodeGrant     GrantType = "authorization_code"
	RefreshTokenGrant GrantType = "refresh_token"
)
