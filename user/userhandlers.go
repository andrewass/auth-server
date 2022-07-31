package user

import (
	"auth-server/user/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpUserRoutes(router *gin.Engine) {
	router.POST("/user/sign-up", signUpUserHandler)
}

func signUpUserHandler(context *gin.Context) {
	request := dto.SignUpUserRequest{}
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	email := signUpUser(request)
	context.JSON(http.StatusOK, email)
}
