package client

import "time"

type clientType string

const (
	internal clientType = "INTERNAL"
	external clientType = "EXTERNAL"
)

type Client struct {
	UserEmail               string     `bson:"user_email"`
	ClientId                string     `bson:"client_id"`
	ClientSecret            string     `bson:"client_secret"`
	ClientIdIssuedAt        time.Time  `bson:"client_id_issued_at"`
	ClientType              clientType `bson:"client_type"`
	LogoUri                 string     `bson:"logo_uri"`
	ApplicationType         string     `bson:"application_type"`
	ClientName              string     `bson:"client_name"`
	ClientUri               string     `bson:"client_uri"`
	InitiateLoginUri        string     `bson:"initiate_login_uri"`
	TokenEndpointAuthMethod string     `bson:"token_endpoint_auth_method"`
	RedirectUris            []string   `bson:"redirect_uris"`
	ResponseTypes           []string   `bson:"response_types"`
	GrantTypes              []string   `bson:"grant_types"`
	PostLogoutRedirectUris  []string   `bson:"post_logout_redirect_uris"`
}
