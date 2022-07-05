package authorization

import (
	"auth-server/authorization/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpAuthorizationRoutes(router *gin.Engine) {
	router.GET("/authorize", authorizeClientHandler)
}

func authorizeClientHandler(context *gin.Context) {
	request := dto.AuthorizeRequest{}
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	var response = AuthorizeClient(request)
	context.JSON(http.StatusOK, response)
}
