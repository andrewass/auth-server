package dto

type SignUpUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
