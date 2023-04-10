package user

import (
	"auth-server/token"
	"auth-server/user/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type UserHandlers struct {
	UserService  *UserService
	TokenService *token.TokenService
}

func (h *UserHandlers) SetUpUserRoutes(router *gin.Engine) {
	router.POST("/user/sign-in", h.signInUserHandler)
	router.POST("/user/sign-up", h.signUpUserHandler)
	router.GET("/user/info", h.getUserInfoHandler)
}

func (h *UserHandlers) getUserInfoHandler(context *gin.Context) {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, "Bearer ")
	subject := h.TokenService.ExtractSubjectFromToken(splitToken[1])
	user := h.UserService.getUserBySubject(subject)
	userInfo := h.mapToUserInfo(user)

	context.JSON(http.StatusOK, userInfo)
}

func (h *UserHandlers) signInUserHandler(context *gin.Context) {
	request := dto.SignInUserRequest{}
	err := context.BindJSON(&request)

	if err != nil {
		panic(err)
	}
	signInError := h.UserService.signInUser(request)
	if signInError == nil {
		context.Status(http.StatusOK)
	} else {
		context.Status(http.StatusUnauthorized)
	}
}

func (h *UserHandlers) signUpUserHandler(context *gin.Context) {
	request := dto.SignUpUserRequest{}
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	h.UserService.signUpUser(request)
	context.Status(http.StatusOK)
}

func (h *UserHandlers) mapToUserInfo(user User) dto.UserInfoResponse {
	return dto.UserInfoResponse{
		Email: user.Email,
	}
}
