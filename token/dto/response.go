package dto

type IntrospectTokenResponse struct {
	Active   bool   `json:"active"`
	Username string `json:"username"`
}
