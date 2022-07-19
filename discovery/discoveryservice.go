package discovery

import (
	"auth-server/discovery/dto"
	"github.com/spf13/viper"
)

func getOpenIdConfig() dto.ConfigResponse {
	return getConfig()
}

func getOauthConfig() dto.ConfigResponse {
	return getConfig()
}

func getConfig() dto.ConfigResponse {
	baseUrl := viper.Get("BASE_URL").(string)
	frontEndUrl := viper.Get("FRONTEND_URL").(string)
	return dto.ConfigResponse{
		RegistrationEndpoint:  baseUrl + "/clients",
		AuthorizationEndpoint: frontEndUrl + "/sign-in",
		TokenEndpoint:         baseUrl + "/token",
		IntrospectionEndpoint: baseUrl + "/token/introspect",
	}
}
