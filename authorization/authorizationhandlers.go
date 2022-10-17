package authorization

import (
	"auth-server/authorization/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpAuthorizationRoutes(router *gin.Engine) {
	router.GET("/authorize", authorizeUserHandler)
}

func authorizeUserHandler(context *gin.Context) {
	request := dto.AuthorizeRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	var response = authorizeClient(request)
	context.JSON(http.StatusOK, response)
}
