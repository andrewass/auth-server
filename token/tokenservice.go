package token

import (
	"auth-server/token/dto"
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
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
