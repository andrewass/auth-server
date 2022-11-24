package token

import (
	"auth-server/authorization"
	"auth-server/client"
	"auth-server/common"
	"auth-server/token/types"
	"crypto/sha256"
	"encoding/base64"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"time"
)

func getTokens(request types.GetTokenRequest) types.GetTokensResponse {
	persistedCode := authorization.GetPersistedAuthorizationCode(request.Code)
	authorization.DeleteAuthorizationCode(persistedCode)
	userEmail := persistedCode.UserEmail

	return types.GetTokensResponse{
		AccessToken: createAccessToken(),
		IdToken:     createIdToken(userEmail, persistedCode.ClientId),
		TokenType:   "Bearer",
		ExpiresIn:   324355346,
		Scope:       "email",
	}
}

func validateClientSecret(clientId string, secret string) {
	if secret != "" {
		persistedClient := client.GetClient(clientId)
		if secret != persistedClient.ClientSecret {
			panic("Client secret mismatch")
		}
	}
}

func validateCodeChallenge(persistedCode authorization.AuthCode, verifier string) {
	if verifier != "" {
		challenge := persistedCode.CodeChallenge
		hash := sha256.New()
		hash.Write([]byte(verifier))
		result := base64.RawURLEncoding.EncodeToString(hash.Sum(nil))
		if challenge != result {
			panic("Challenge mismatch. Expected : " + challenge + " , Received : " + result)
		}
	}
}

func createAccessToken() string {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Issuer:    viper.Get("BASE_URL").(string),
		Subject:   "1234567890",
		NotBefore: time.Now().Unix(),
		Audience:  "test-audience",
		IssuedAt:  time.Now().Unix(),
	}
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(common.GetPrivateKey())
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedString, _ := token.SignedString(key)

	return signedString
}

func createIdToken(email string, clientId string) string {

	claims := types.CustomIdClaims{
		Email: email,
		Id:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    viper.Get("BASE_URL").(string),
			Subject:   "1234567890-id",
			NotBefore: time.Now().Unix(),
			Audience:  clientId,
			IssuedAt:  time.Now().Unix(),
		}}
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(common.GetPrivateKey())
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedString, _ := token.SignedString(key)

	return signedString
}

func introspectToken(request types.IntrospectTokenRequest) types.IntrospectTokenResponse {
	return types.IntrospectTokenResponse{
		Active:   true,
		Username: "testuser",
	}
}

func revokeToken(request types.RevokeTokenRequest) {
}
