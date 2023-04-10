package user

import (
	"auth-server/user/dto"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repository *UserRepository
}

func (s *UserService) getUserBySubject(subject string) User {
	return *s.Repository.findUserBySubject(subject)
}

func (s *UserService) signInUser(request dto.SignInUserRequest) error {
	existingUser := s.Repository.findUserByEmail(request.Email)
	passwordError := s.matchingPassword(request.Password, existingUser.Password)
	if existingUser == nil || passwordError != nil {
		return errors.New("invalid password or username")
	}
	return nil
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

func (s *UserService) matchingPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (s *UserService) verifyNotPreviouslyExistingUser(email string) {
	if s.Repository.existsUserByEmail(email) {
		panic("Email already registered")
	}
}
