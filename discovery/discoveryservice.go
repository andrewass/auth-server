package discovery

import "auth-server/discovery/dto"

func getOpenIdConfig() dto.ConfigResponse {
	return getConfig()
}

func getOauthConfig() dto.ConfigResponse {
	return getConfig()
}

func getConfig() dto.ConfigResponse {
	return dto.ConfigResponse{
		RegistrationEndpoint:  "",
		AuthorizationEndpoint: "",
		TokenEndpoint:         "",
		IntrospectionEndpoint: "",
	}
}
