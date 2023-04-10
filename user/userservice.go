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

func (s *UserService) signInUser(request dto.SignInUserRequest) error {
	existingUser := s.Repository.findUserByEmail(request.Email)
	if existingUser == nil || s.getPasswordMismatchError(request.Password, existingUser.Password) != nil {
		return &InvalidCredentialsError{msg: "invalid email or password"}
	}
	return nil
}

func (s *UserService) signUpUser(request dto.SignUpUserRequest) error {
	if s.Repository.existsUserByEmail(request.Email) {
		return &CredentialsConflictError{msg: "email already registered to an account"}
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 8)
	newUser := User{
		Email:    request.Email,
		Password: string(hashedPassword),
		Subject:  request.Email,
	}
	s.Repository.saveUser(newUser)
	return nil
}

func (s *UserService) getPasswordMismatchError(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
