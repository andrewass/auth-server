package user

import (
	"auth-server/user/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repository *UserRepository
}

func (s *UserService) getUserBySubject(subject string) User {
	return *s.Repository.findUserBySubject(subject)
}

func (s *UserService) signInUser(request dto.SignInUserRequest) {
	existingUser := s.Repository.findUserByEmail(request.Email)
	if existingUser == nil || !s.matchingPassword(request.Password, existingUser.Password) {
		panic("User not found or incorrect password")
	}
}

func (s *UserService) signUpUser(request dto.SignUpUserRequest) {
	s.verifyNotPreviouslyExistingUser(request.Email)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 8)
	newUser := User{
		Email:    request.Email,
		Password: string(hashedPassword),
		Subject:  request.Email,
	}
	s.Repository.saveUser(newUser)
}

func (s *UserService) matchingPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *UserService) verifyNotPreviouslyExistingUser(email string) {
	if s.Repository.existsUserByEmail(email) {
		panic("Email already registered")
	}
}
