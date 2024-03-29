package client

import "time"

type ClientInformationDto struct {
	ClientId                string    `json:"client_id"`
	ClientSecret            string    `json:"client_secret"`
	ClientIdIssuedAt        time.Time `json:"client_id_issued_at"`
	ClientSecretIssuedAt    time.Time `json:"client_secret_issued_at"`
	LogoUri                 string    `json:"logo_uri"`
	ApplicationType         string    `json:"application_type"`
	ClientName              string    `json:"client_name" binding:"required"`
	ClientUri               string    `json:"client_uri"`
	InitiateLoginUri        string    `json:"initiate_login_uri"`
	TokenEndpointAuthMethod string    `json:"token_endpoint_auth_method"`
	RedirectUris            []string  `json:"redirect_uris"`
	ResponseTypes           []string  `json:"response_types"`
	GrantTypes              []string  `json:"grant_types"`
	PostLogoutRedirectUris  []string  `json:"post_logout_redirect_uris"`
	UserEmail               string    `json:"user_email"`
}
