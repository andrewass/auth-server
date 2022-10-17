package token

import (
	"auth-server/authorization"
	"auth-server/token/dto"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"time"
)

func getTokens(request dto.GetTokenRequest) dto.GetTokensResponse {
	return dto.GetTokensResponse{
		AccessToken: createAccessToken(),
		IdToken:     createIdToken(),
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
		ExpiresAt: getExpirationTime(),
		Issuer:    viper.Get("AUTH_SERVER_SERVICE").(string),
		Subject:   "1234567890",
		NotBefore: getCurrentTime(),
		Audience:  "test-audience",
	}
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(authorization.GetPrivateKey())
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedString, _ := token.SignedString(key)

	return signedString
}

func createIdToken() string {

	claims := CustomIdClaims{"", jwt.StandardClaims{
		ExpiresAt: getExpirationTime(),
		Issuer:    viper.Get("AUTH_SERVER_SERVICE").(string),
		Subject:   "1234567890-id",
		NotBefore: getCurrentTime(),
		Audience:  "test-audience-id-token",
	},
	}
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(authorization.GetPrivateKey())
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedString, _ := token.SignedString(key)

	return signedString
}

func getCurrentTime() int64 {
	return time.Now().Local().Unix()
}

func getExpirationTime() int64 {
	return time.Now().Local().Add(time.Hour * 1).Unix()
}
