package user

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Subject  string `json:"subject"`
}
