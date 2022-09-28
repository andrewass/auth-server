package discovery

import (
	"auth-server/authorization"
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
		Keys: authorization.GetJwks(),
	}
}

func getConfig() dto.WellKnownEndpointsResponse {
	baseUrl := viper.Get("BASE_URL").(string)
	frontEndUrl := viper.Get("FRONTEND_URL").(string)
	return dto.WellKnownEndpointsResponse{
		Issuer:                baseUrl,
		RegistrationEndpoint:  baseUrl + "/clients",
		AuthorizationEndpoint: frontEndUrl + "/authentication",
		TokenEndpoint:         baseUrl + "/token",
		IntrospectionEndpoint: baseUrl + "/token/introspect",
		JwksUri:               baseUrl + "/.well-known/jwks",
	}
}
