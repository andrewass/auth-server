package token

import "auth-server/token/dto"

func IntrospectToken(request dto.IntrospectTokenRequest) dto.IntrospectTokenResponse {
	return dto.IntrospectTokenResponse{
		Active:   true,
		Username: "testuser",
	}
}

func RevokeToken(request dto.RevokeTokenRequest) {
}
