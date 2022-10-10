package user

import (
	"auth-server/user/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpUserRoutes(router *gin.Engine) {
	router.POST("/user/sign-in", signInUserHandler)
	router.POST("/user/sign-up", signUpUserHandler)
}

func signInUserHandler(context *gin.Context) {
	request := dto.SignInUserRequest{}
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	signInUser(request)
	context.Status(http.StatusOK)
}

func signUpUserHandler(context *gin.Context) {
	request := dto.SignUpUserRequest{}
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	signUpUser(request)
	context.Status(http.StatusOK)
}
