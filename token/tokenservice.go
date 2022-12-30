package token

import (
	"auth-server/authorization"
	"auth-server/client"
	"auth-server/common"
	"auth-server/token/types"
	"crypto/sha256"
	"encoding/base64"
	"github.com/golang-jwt/jwt"
	"time"
)

func processGetTokensRequest(request types.GetTokenRequest) types.GetTokensResponse {
	switch request.GrantType {
	case types.AuthCodeGrant:
		return processAuthCodeGrantRequest(request)
	case types.RefreshTokenGrant:
		return processRefreshTokenGrantRequest(request)
	}
	panic("Illegal state when processing token request")
}

func processAuthCodeGrantRequest(request types.GetTokenRequest) types.GetTokensResponse {
	persistedCode := authorization.GetPersistedAuthorizationCode(request.Code)
	authorization.DeleteAuthorizationCode(persistedCode)
	return createTokens(persistedCode.UserEmail, persistedCode.ClientId)
}

func processRefreshTokenGrantRequest(request types.GetTokenRequest) types.GetTokensResponse {
	persistedClient := client.GetClientByIdAndSecret(request.ClientId, request.ClientSecret)
	subject := ExtractSubjectFromToken(request.RefreshToken)
	return createTokens(subject, persistedClient.ClientId)
}

func createTokens(subject string, clientId string) types.GetTokensResponse {
	return types.GetTokensResponse{
		AccessToken:  createAccessToken(subject),
		RefreshToken: createRefreshToken(subject),
		IdToken:      createIdToken(subject, clientId),
		TokenType:    "bearer",
		ExpiresIn:    300,
		Scope:        "subject",
	}
}

func ExtractSubjectFromToken(token string) string {
	var claims jwt.StandardClaims
	key, _ := jwt.ParseRSAPublicKeyFromPEM(common.GetPublicKey())

	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		panic("Error extracting claims from token : " + err.Error())
	}
	return claims.Subject
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
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
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

func createRefreshToken(subject string) string {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 100).Unix(),
		Issuer:    "http://auth-backend-service:8089",
		Subject:   subject,
		NotBefore: time.Now().Unix(),
		IssuedAt:  time.Now().Unix(),
	}
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(common.GetPrivateKey())
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedString, _ := token.SignedString(key)

	return signedString
}

func createIdToken(subject string, clientId string) string {
	claims := types.CustomIdClaims{
		Email: subject,
		Id:    subject,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    "http://auth-backend-service:8089",
			Subject:   subject,
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
