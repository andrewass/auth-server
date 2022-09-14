package dto

import "auth-server/resources"

type WellKnownEndpointsResponse struct {
	Issuer                string `json:"issuer"`
	RegistrationEndpoint  string `json:"registration_endpoint"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	IntrospectionEndpoint string `json:"introspection_endpoint"`
	JwksUri               string `json:"jwks_uri"`
}

type JWKResponse struct {
	Keys []resources.JWK `json:"keys"`
}
