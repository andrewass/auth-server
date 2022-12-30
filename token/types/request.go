package types

type GetTokenRequest struct {
	GrantType    GrantType `json:"grant_type" form:"grant_type"`
	RedirectUri  string    `json:"redirect_uri" form:"redirect_uri"`
	Code         string    `json:"code" form:"code"`
	CodeVerifier string    `json:"code_verifier" form:"code_verifier"`
	ClientId     string    `json:"client_id" form:"client_id"`
	ClientSecret string    `json:"client_secret" form:"client_secret"`
	RefreshToken string    `json:"refresh_token" form:"refresh_token"`
}

type IntrospectTokenRequest struct {
	Token         string   `json:"token"`
	TokenTypeHint TypeHint `json:"token_type_hint"`
}

type RevokeTokenRequest struct {
	Token         string   `json:"token"`
	TokenTypeHint TypeHint `json:"token_type_hint"`
}
