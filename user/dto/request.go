package dto

type SignUpUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
