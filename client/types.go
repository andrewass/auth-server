package client

type Client struct {
	UserEmail               string   `json:"user_email" bson:"user_email"`
	ClientId                string   `json:"clientId" bson:"client_id"`
	ClientSecret            string   `json:"clientSecret" bson:"client_secret"`
	ClientIdIssuedAt        int64    `json:"client_id_issued_at" bson:"client_id_issued_at"`
	LogoUri                 string   `json:"logo_uri" bson:"logo_uri"`
	ApplicationType         string   `json:"application_type" bson:"application_type"`
	ClientName              string   `json:"client_name" bson:"client_name"`
	ClientUri               string   `json:"client_uri" bson:"client_uri"`
	InitiateLoginUri        string   `json:"initiate_login_uri" bson:"initiate_login_uri"`
	TokenEndpointAuthMethod string   `json:"token_endpoint_auth_method" bson:"token_endpoint_auth_method"`
	RedirectUris            []string `json:"redirect_uris" bson:"redirect_uris"`
	ResponseTypes           []string `json:"response_types" bson:"response_types"`
	GrantTypes              []string `json:"grant_types" bson:"grant_types"`
	PostLogoutRedirectUris  []string `json:"post_logout_redirect_uris" bson:"post_logout_redirect_uris"`
}
