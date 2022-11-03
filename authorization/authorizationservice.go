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

func DeleteAuthorizationCode(authCode AuthCode) {
	deleteAuthorizationCode(authCode)
}

func GetPersistedAuthorizationCode(clientId string, code string) AuthCode {
	return getAuthorizationCode(clientId, code)
}

func createFrontendUrl(request dto.AuthorizeRequest) string {
	frontendUrl, _ := decideFrontendUrl(request.ClientId)
	values := frontendUrl.Query()
	values.Add("client_id", request.ClientId)
	values.Add("state", request.State)
	values.Add("redirect_uri", request.RedirectUri)
	values.Add("code_challenge", request.CodeChallenge)
	values.Add("code_challenge_method", request.CodeChallengeMethod)
	frontendUrl.RawQuery = values.Encode()

	return frontendUrl.String()
}

func decideFrontendUrl(clientId string) (*url.URL, error) {
	requestedClient := client.GetClient(clientId)
	if requestedClient.ClientType == client.Internal {
		return url.Parse(viper.Get("FRONTEND_URL").(string) + "/authentication")
	}
	return url.Parse(viper.Get("FRONTEND_URL").(string) + "/confirmation")
}

func createAndSaveAuthorizationCode(request dto.AuthorizationCodeRequest) string {
	code := generateRandomString()
	authorizationCode := AuthCode{
		Code:                code,
		CodeChallenge:       request.CodeChallenge,
		CodeChallengeMethod: CodeChallengeMethod(request.CodeChallengeMethod),
		ClientId:            request.ClientId,
		UserEmail:           request.Email,
		ExpirationTime:      time.Now().Add(time.Minute * 10).Unix(),
	}
	saveAuthorizationCode(authorizationCode)
	return code
}

func generateRandomString() string {
	b := make([]rune, 30)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
