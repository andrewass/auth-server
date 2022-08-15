package dto

type AddClientRequest struct {
	LogoUri                 string   `json:"logo_uri"`
	ApplicationType         string   `json:"application_type"`
	ClientName              string   `json:"client_name" binding:"required"`
	ClientUri               string   `json:"client_uri"`
	InitiateLoginUri        string   `json:"initiate_login_uri"`
	TokenEndpointAuthMethod string   `json:"token_endpoint_auth_method"`
	RedirectUris            []string `json:"redirect_uris"`
	ResponseTypes           []string `json:"response_types"`
	GrantTypes              []string `json:"grant_types"`
	PostLogoutRedirectUris  []string `json:"post_logout_redirect_uris"`
}
