package authorization

import (
	"auth-server/authorization/dto"
	"auth-server/client"
	"github.com/spf13/viper"
	"math/rand"
	"net/url"
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

func (s *AuthorizationService) createFrontendUrl(request dto.AuthorizeRequest) string {
	frontendUrl, _ := s.decideFrontendUrl(request.ClientId)
	values := frontendUrl.Query()
	values.Add("client_id", request.ClientId)
	values.Add("state", request.State)
	values.Add("redirect_uri", request.RedirectUri)
	values.Add("code_challenge", request.CodeChallenge)
	values.Add("code_challenge_method", request.CodeChallengeMethod)
	frontendUrl.RawQuery = values.Encode()

	return frontendUrl.String()
}

func (s *AuthorizationService) decideFrontendUrl(clientId string) (*url.URL, error) {
	requestedClient := s.ClientService.GetClientById(clientId)
	if requestedClient.ClientType == client.Internal {
		return url.Parse(viper.Get("FRONTEND_URL").(string) + "/authentication/internal")
	}
	return url.Parse(viper.Get("FRONTEND_URL").(string) + "/authentication/external")
}

func (s *AuthorizationService) createAndSaveAuthorizationCode(request dto.AuthorizationCodeRequest) string {
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
	return code
}

func generateRandomString() string {
	b := make([]rune, 30)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
