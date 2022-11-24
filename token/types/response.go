package types

type GetTokensResponse struct {
	AccessToken string `json:"access_token"`
	IdToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
}

type IntrospectTokenResponse struct {
	Active    bool   `json:"active"`
	Username  string `json:"username"`
	Scope     string `json:"scope"`
	ClientId  string `json:"client_id"`
	ExpiresIn int64  `json:"exp"`
}
