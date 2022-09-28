package token

import (
	"auth-server/authorization"
	"auth-server/token/dto"
	"github.com/golang-jwt/jwt"
	"log"
)

func getToken(request dto.GetTokenRequest) dto.GetTokenResponse {
	return dto.GetTokenResponse{
		AccessToken: createAccessToken(),
		IdToken:     createAccessToken(),
		TokenType:   "Bearer",
		ExpiresIn:   324355346,
		Scope:       "email",
	}
}

func introspectToken(request dto.IntrospectTokenRequest) dto.IntrospectTokenResponse {
	return dto.IntrospectTokenResponse{
		Active:   true,
		Username: "testuser",
	}
}

func revokeToken(request dto.RevokeTokenRequest) {
}

func createAccessToken() string {
	claims := jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
		Subject:   "1234567890",
		NotBefore: 34234534565,
		Audience:  "test-audience",
	}
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(authorization.GetPrivateKey())
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedString, err := token.SignedString(key)
	if err != nil {
		log.Println(err)
	}
	return signedString
}
