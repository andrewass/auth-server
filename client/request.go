package client

type AddClientRequest struct {
	UserEmail               string   `json:"user_email"`
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
	ClientId                string
	ClientSecret            string
	Type                    ClientType
}

type UpdateClientRequest struct {
	ClientID                string   `json:"client_id"`
	ClientSecret            string   `json:"client_secret"`
	LogoUri                 string   `json:"logo_uri"`
	ApplicationType         string   `json:"application_type"`
	ClientName              string   `json:"client_name" binding:"required"`
	ClientUri               string   `json:"client_uri"`
	ClientType              string   `json:"client_type"`
	InitiateLoginUri        string   `json:"initiate_login_uri"`
	TokenEndpointAuthMethod string   `json:"token_endpoint_auth_method"`
	RedirectUris            []string `json:"redirect_uris"`
	ResponseTypes           []string `json:"response_types"`
	GrantTypes              []string `json:"grant_types"`
	PostLogoutRedirectUris  []string `json:"post_logout_redirect_uris"`
}

type DeleteClientRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
