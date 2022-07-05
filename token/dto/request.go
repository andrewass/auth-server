package dto

type IntrospectTokenRequest struct {
	Token         string `json:"token"`
	TokenTypeHint string `json:"token_type_hint"`
}

type RevokeTokenRequest struct {
	Token         string `json:"token"`
	TokenTypeHint string `json:"token_type_hint"`
}
