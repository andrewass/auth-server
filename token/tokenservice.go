package token

import (
	"auth-server/authorization"
	"auth-server/common"
	"auth-server/token/dto"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"time"
)

func getTokens(request dto.GetTokenRequest) dto.GetTokensResponse {
	persistedCode := authorization.GetPersistedAuthorizationCode(request.ClientId, request.Code)
	codeIsInvalid := isInvalidAuthorizationCode(persistedCode, request.ClientId, request.Code)
	userEmail := persistedCode.UserEmail
	authorization.DeleteAuthorizationCode(persistedCode)
	if codeIsInvalid {
		panic("Invalid authorization code")
	}
	return dto.GetTokensResponse{
		AccessToken: createAccessToken(),
		IdToken:     createIdToken(userEmail),
		TokenType:   "Bearer",
		ExpiresIn:   324355346,
		Scope:       "email",
	}
}

func isInvalidAuthorizationCode(persistedCode authorization.AuthCode, clientId string, requestCode string) bool {
	return persistedCode.ClientId != clientId ||
		persistedCode.Code != requestCode ||
		persistedCode.ExpirationTime < time.Now().Unix()
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
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(common.GetPrivateKey())
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedString, _ := token.SignedString(key)

	return signedString
}

func createIdToken(email string) string {

	claims := CustomIdClaims{email, jwt.StandardClaims{
		ExpiresAt: getExpirationTime(),
		Issuer:    viper.Get("AUTH_SERVER_SERVICE").(string),
		Subject:   "1234567890-id",
		NotBefore: getCurrentTime(),
		Audience:  "test-audience-id-token",
	},
	}
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(common.GetPrivateKey())
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedString, _ := token.SignedString(key)

	return signedString
}

func getCurrentTime() int64 {
	return time.Now().Unix()
}

func getExpirationTime() int64 {
	return time.Now().Add(time.Hour * 1).Unix()
}
