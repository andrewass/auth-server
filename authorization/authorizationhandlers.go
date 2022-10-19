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
	router.POST("/authorization-confirmation", authorizationConfirmationHandler)
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

func authorizationConfirmationHandler(context *gin.Context) {
	var request dto.AuthorizationConfirmationRequest
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	redirectUrl := constructResponseRedirectUrl(request)
	context.JSON(200, redirectUrl)
}

func constructResponseRedirectUrl(request dto.AuthorizationConfirmationRequest) string {
	redirectUrl, _ := url.Parse(request.RedirectUri)
	values := redirectUrl.Query()
	values.Add("code", "fdsfsdtgertewtfdgfgyh")
	values.Add("state", request.State)
	redirectUrl.RawQuery = values.Encode()

	return redirectUrl.String()
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
