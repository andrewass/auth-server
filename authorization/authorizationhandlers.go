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
}

func constructFrontendUrl(request dto.AuthorizeRequest) string {
	frontendUrl, _ := url.Parse(viper.Get("FRONTEND_URL").(string) + "/authentication")
	values := frontendUrl.Query()
	values.Add("client_id", request.ClientId)
	frontendUrl.RawQuery = values.Encode()

	return frontendUrl.String()
}
