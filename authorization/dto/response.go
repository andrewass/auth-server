package dto

type AuthorizationCodeResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
}
