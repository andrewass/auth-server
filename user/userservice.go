package user

import (
	"auth-server/user/dto"
	"golang.org/x/crypto/bcrypt"
)

func signUpUser(request dto.SignUpUserRequest) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 8)
	newUser := User{
		Email:    request.Email,
		Password: string(hashedPassword),
	}
	verifyNotPreviouslyExistingUser(newUser)
	saveUser(newUser)
	return request.Email
}

func verifyNotPreviouslyExistingUser(user User) {
	if existsUserByEmail(user.Email) {
		panic("Email already registered")
	}
}
