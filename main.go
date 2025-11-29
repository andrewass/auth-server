package main

import (
	"auth-server/authorization"
	"auth-server/client"
	"auth-server/common"
	"auth-server/discovery"
	"auth-server/token"
	"auth-server/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	common.ConfigureKeys()
	configureEnvironmentVariables()
	configurePersistenceLayer()
	configureRoutes()
}

func configureEnvironmentVariables() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func configurePersistenceLayer() {
	common.ConnectToDatabase()
}

func configureRoutes() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	userService := user.UserService{
		Repository: &user.UserRepository{Collection: common.Database.Collection("user")},
	}
	clientService := client.ClientService{
		Repository: &client.ClientRepository{Collection: common.Database.Collection("client")},
	}
	authorizationService := authorization.AuthorizationService{
		Repository:    &authorization.AuthorizationRepository{Collection: common.Database.Collection("authorizationCode")},
		ClientService: &clientService,
	}
	tokenService := token.TokenService{AuthorizationService: &authorizationService, ClientService: &clientService}

	discovery.SetUpDiscoveryRoutes(router)

	userHandlers := &user.UserHandlers{UserService: &userService, TokenService: &tokenService}
	userHandlers.SetUpUserRoutes(router)

	clientHandlers := &client.ClientHandlers{Service: &client.ClientService{
		Repository: &client.ClientRepository{Collection: common.Database.Collection("client")}},
	}
	clientHandlers.SetUpClientRoutes(router)
	//For testing during development only
	clientHandlers.AddAdminClient()

	authorizationHandlers := &authorization.AuthorizationHandler{Service: &authorizationService}
	authorizationHandlers.SetUpAuthorizationRoutes(router)

	tokenHandlers := &token.TokenHandlers{Service: &tokenService}
	tokenHandlers.SetUpTokenRoutes(router)

	if err := router.Run(":8089"); err != nil {
		panic(err)
	}
}
