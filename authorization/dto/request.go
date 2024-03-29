package dto

type AuthorizeRequest struct {
	ResponseType        string `json:"response_id" form:"response_type"`
	ClientId            string `json:"client_id" form:"client_id"`
	RedirectUri         string `json:"redirect_uri" form:"redirect_uri"`
	Scope               string `json:"scope" form:"scope"`
	State               string `json:"state" form:"state"`
	CodeChallenge       string `json:"code_challenge" form:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method" form:"code_challenge_method"`
	ResponseMode        string `json:"response_mode" form:"response_mode"`
}

type AuthorizationCodeRequest struct {
	Email               string `json:"email"`
	ClientId            string `json:"client_id"`
	CodeChallenge       string `json:"code_challenge" form:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method" form:"code_challenge_method"`
	State               string `json:"state"`
}
