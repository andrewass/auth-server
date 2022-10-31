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
	baseUrl := viper.Get("BASE_URL").(string)
	return dto.WellKnownEndpointsResponse{
		Issuer:                baseUrl,
		RegistrationEndpoint:  baseUrl + "/clients",
		AuthorizationEndpoint: baseUrl + "/authorize",
		TokenEndpoint:         baseUrl + "/token",
		IntrospectionEndpoint: baseUrl + "/token/introspect",
		JwksUri:               baseUrl + "/.well-known/jwks.json",
	}
}
