package client

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Client struct {
	ID                      primitive.ObjectID `bson:"_id"`
	UserEmail               string             `bson:"user_email"`
	ClientId                string             `bson:"client_id"`
	ClientSecret            string             `bson:"client_secret"`
	ClientIdIssuedAt        time.Time          `bson:"client_id_issued_at"`
	ClientSecretIssuedAt    time.Time          `bson:"client_secret_issued_at"`
	ClientDescription       string             `bson:"client_description"`
	LogoUri                 string             `bson:"logo_uri"`
	ApplicationType         string             `bson:"application_type"`
	ClientName              string             `bson:"client_name"`
	ClientUri               string             `bson:"client_uri"`
	InitiateLoginUri        string             `bson:"initiate_login_uri"`
	TokenEndpointAuthMethod string             `bson:"token_endpoint_auth_method"`
	RedirectUris            []string           `bson:"redirect_uris"`
	ResponseTypes           []string           `bson:"response_types"`
	GrantTypes              []string           `bson:"grant_types"`
	PostLogoutRedirectUris  []string           `bson:"post_logout_redirect_uris"`
}
