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

type TokenService struct {
	AuthorizationService *authorization.AuthorizationService
	ClientService        *client.ClientService
}

func (s *TokenService) processGetTokensRequest(request types.GetTokenRequest) types.GetTokensResponse {
	switch request.GrantType {
	case types.AuthCodeGrant:
		return s.processAuthCodeGrantRequest(request)
	case types.RefreshTokenGrant:
		return s.processRefreshTokenGrantRequest(request)
	}
	panic("Illegal state when processing token request")
}

func (s *TokenService) processAuthCodeGrantRequest(request types.GetTokenRequest) types.GetTokensResponse {
	persistedCode := s.AuthorizationService.GetPersistedAuthorizationCode(request.Code)
	s.AuthorizationService.DeleteAuthorizationCode(persistedCode)
	return s.createTokens(persistedCode.UserEmail, persistedCode.ClientId)
}

func (s *TokenService) processRefreshTokenGrantRequest(request types.GetTokenRequest) types.GetTokensResponse {
	persistedClient := s.ClientService.GetClientByIdAndSecret(request.ClientId, request.ClientSecret)
	subject := s.ExtractSubjectFromToken(request.RefreshToken)
	return s.createTokens(subject, persistedClient.ClientId)
}

func (s *TokenService) createTokens(subject string, clientId string) types.GetTokensResponse {
	return types.GetTokensResponse{
		AccessToken:  s.createAccessToken(subject),
		RefreshToken: s.createRefreshToken(subject),
		IdToken:      s.createIdToken(subject, clientId),
		TokenType:    "bearer",
		ExpiresIn:    300,
		Scope:        "subject",
	}
}

func (s *TokenService) ExtractSubjectFromToken(token string) string {
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

func (s *TokenService) validateCodeChallenge(persistedCode authorization.AuthCode, verifier string) {
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

func (s *TokenService) createAccessToken(subject string) string {
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

func (s *TokenService) createRefreshToken(subject string) string {
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

func (s *TokenService) createIdToken(subject string, clientId string) string {
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

func (s *TokenService) introspectToken(request types.IntrospectTokenRequest) types.IntrospectTokenResponse {
	return types.IntrospectTokenResponse{
		Active:   true,
		Username: "testuser",
	}
}

func (s *TokenService) revokeToken(request types.RevokeTokenRequest) {
}
