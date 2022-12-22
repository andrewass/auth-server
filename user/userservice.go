package user

import (
	"auth-server/user/dto"
	"golang.org/x/crypto/bcrypt"
)

func getUserBySubject(subject string) User {
	return *findUserBySubject(subject)
}

func signInUser(request dto.SignInUserRequest) {
	existingUser := findUserByEmail(request.Email)
	if existingUser == nil || !matchingPassword(request.Password, existingUser.Password) {
		panic("User not found or incorrect password")
	}
}

func signUpUser(request dto.SignUpUserRequest) {
	verifyNotPreviouslyExistingUser(request.Email)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 8)
	newUser := User{
		Email:    request.Email,
		Password: string(hashedPassword),
		Subject:  request.Email,
	}
	saveUser(newUser)
}

func matchingPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func verifyNotPreviouslyExistingUser(email string) {
	if existsUserByEmail(email) {
		panic("Email already registered")
	}
}
