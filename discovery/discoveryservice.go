package discovery

import (
	"auth-server/common"
	"auth-server/discovery/dto"
	"github.com/spf13/viper"
)

func getOpenIdConfig() dto.WellKnownEndpointsResponse {
	return getConfig()
}

func getOauthConfig() dto.WellKnownEndpointsResponse {
	return getConfig()
}

func getJwks() dto.JWKResponse {
	return dto.JWKResponse{
		Keys: common.GetJwks(),
	}
}

func getConfig() dto.WellKnownEndpointsResponse {
	frontendUrl := viper.Get("FRONTEND_URL").(string)
	backendUrl := viper.Get("BACKEND_URL").(string)
	return dto.WellKnownEndpointsResponse{
		Issuer:                backendUrl,
		RegistrationEndpoint:  backendUrl + "/clients",
		AuthorizationEndpoint: frontendUrl + "/authentication",
		TokenEndpoint:         backendUrl + "/server/token/token",
		IntrospectionEndpoint: backendUrl + "/server/token/introspect",
		JwksUri:               frontendUrl + "/server/.well-known/jwks.json",
	}
}
