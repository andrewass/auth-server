package authorization

import (
	"auth-server/authorization/dto"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
)

func SetUpAuthorizationRoutes(router *gin.Engine) {
	router.GET("/authorize", authorizeUserHandler)
	router.POST("/authorization-response", authorizationResponseHandler)
}

func authorizeUserHandler(context *gin.Context) {
	var request dto.AuthorizeRequest
	err := context.BindQuery(&request)
	if err != nil {
		panic(err)
	}
	frontendUrl := constructFrontendUrl(request)
	context.Redirect(http.StatusFound, frontendUrl)
}

func authorizationResponseHandler(context *gin.Context) {
	var request dto.AuthorizationCodeRequest
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}

}

func constructFrontendUrl(request dto.AuthorizeRequest) string {
	frontendUrl, _ := url.Parse(viper.Get("FRONTEND_URL").(string) + "/confirmation")
	values := frontendUrl.Query()
	values.Add("client_id", request.ClientId)
	values.Add("state", request.State)
	values.Add("redirect_uri", request.RedirectUri)
	frontendUrl.RawQuery = values.Encode()

	return frontendUrl.String()
}
