package token

import (
	"auth-server/authorization"
	"auth-server/client"
	"auth-server/common"
	"auth-server/token/types"
	"crypto/sha256"
	"encoding/base64"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

func createTokenResponse(request types.GetTokenRequest) types.GetTokensResponse {
	persistedCode := authorization.GetPersistedAuthorizationCode(request.Code)
	authorization.DeleteAuthorizationCode(persistedCode)

	return types.GetTokensResponse{
		AccessToken: createAccessToken(persistedCode.UserEmail),
		IdToken:     createIdToken(persistedCode.UserEmail, persistedCode.ClientId),
		TokenType:   "Bearer",
		ExpiresIn:   324355346,
		Scope:       "email",
	}
}

func ExtractSubjectFromToken(bearerToken string) string {
	splitToken := strings.Split(bearerToken, "Bearer ")
	var claims jwt.StandardClaims
	key, _ := jwt.ParseRSAPublicKeyFromPEM(common.GetPublicKey())

	_, err := jwt.ParseWithClaims(splitToken[1], &claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		panic("Error extracting claims from token : " + err.Error())
	}
	return claims.Subject
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

func createAccessToken(subject string) string {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Issuer:    "http://auth-backend-service:8089",
		Subject:   subject,
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
			Issuer:    "http://auth-backend-service:8089",
			Subject:   email,
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
