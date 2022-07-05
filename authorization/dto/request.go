package dto

type AuthorizeRequest struct {
	ResponseType string `json:"responseType"`
	ClientId     string `json:"clientId"`
}
