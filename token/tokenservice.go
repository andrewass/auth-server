package token

import (
	"auth-server/authorization"
	"auth-server/client"
	"auth-server/common"
	"auth-server/token/dto"
	"crypto/sha256"
	"encoding/base64"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"time"
)

func getTokens(request dto.GetTokenRequest) dto.GetTokensResponse {
	persistedCode := authorization.GetPersistedAuthorizationCode(request.ClientId, request.Code)
	authorization.DeleteAuthorizationCode(persistedCode)
	validateRequest(persistedCode, request)
	userEmail := persistedCode.UserEmail

	return dto.GetTokensResponse{
		AccessToken: createAccessToken(),
		IdToken:     createIdToken(userEmail),
		TokenType:   "Bearer",
		ExpiresIn:   324355346,
		Scope:       "email",
	}
}

func validateRequest(persistedCode authorization.AuthCode, request dto.GetTokenRequest) {
	if persistedCode.ClientId != request.ClientId {
		panic("Client id mismatch. Expected : " + persistedCode.ClientId + " , Received : " + request.ClientId)
	}
	if persistedCode.Code != request.Code {
		panic("Authorization code mismatch. Expected " + persistedCode.Code + " , Received : " + request.Code)
	}
	if persistedCode.ExpirationTime < time.Now().Unix() {
		panic("Expired authorization code")
	}
	validateClientSecret(request.ClientId, request.ClientSecret)
	validateCodeChallenge(persistedCode, request.CodeVerifier)
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
		Issuer:    viper.Get("AUTH_SERVER_SERVICE").(string),
		Subject:   "1234567890",
		NotBefore: time.Now().Unix(),
		Audience:  "test-audience",
	}
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(common.GetPrivateKey())
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedString, _ := token.SignedString(key)

	return signedString
}

func createIdToken(email string) string {

	claims := CustomIdClaims{email, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Issuer:    viper.Get("AUTH_SERVER_SERVICE").(string),
		Subject:   "1234567890-id",
		NotBefore: time.Now().Unix(),
		Audience:  "test-audience-id-token",
	}}
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(common.GetPrivateKey())
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedString, _ := token.SignedString(key)

	return signedString
}

func introspectToken(request dto.IntrospectTokenRequest) dto.IntrospectTokenResponse {
	return dto.IntrospectTokenResponse{
		Active:   true,
		Username: "testuser",
	}
}

func revokeToken(request dto.RevokeTokenRequest) {
}
