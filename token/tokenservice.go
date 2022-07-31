package token

import "auth-server/token/dto"

func introspectToken(request dto.IntrospectTokenRequest) dto.IntrospectTokenResponse {
	return dto.IntrospectTokenResponse{
		Active:   true,
		Username: "testuser",
	}
}

func revokeToken(request dto.RevokeTokenRequest) {
}
