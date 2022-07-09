package dto

type ConfigResponse struct {
	RegistrationEndpoint  string `json:"registration_endpoint"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	IntrospectionEndpoint string `json:"introspection_endpoint"`
}
