package authorization

import (
	"auth-server/authorization/dto"
	"auth-server/client"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type AuthorizationService struct {
	Repository    *AuthorizationRepository
	ClientService *client.ClientService
}

func (s *AuthorizationService) DeleteAuthorizationCode(authCode AuthCode) {
	s.Repository.deleteAuthorizationCode(authCode)
}

func (s *AuthorizationService) GetPersistedAuthorizationCode(code string) AuthCode {
	return s.Repository.getAuthorizationCode(code)
}

func (s *AuthorizationService) createAndSaveAuthorizationCode(request dto.AuthorizationCodeRequest) dto.AuthorizationCodeResponse {
	code := generateRandomString()
	authorizationCode := AuthCode{
		Code:                code,
		CodeChallenge:       request.CodeChallenge,
		CodeChallengeMethod: CodeChallengeMethod(request.CodeChallengeMethod),
		ClientId:            request.ClientId,
		UserEmail:           request.Email,
		ExpirationTime:      time.Now().Add(time.Minute * 10).Unix(),
	}
	s.Repository.saveAuthorizationCode(authorizationCode)
	return dto.AuthorizationCodeResponse{
		Code:  code,
		State: request.State,
	}
}

func generateRandomString() string {
	b := make([]rune, 30)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
