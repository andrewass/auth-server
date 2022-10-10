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
	authorization.ConfigureKeys()
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

	user.SetUpUserRoutes(router)
	discovery.SetUpDiscoveryRoutes(router)
	client.SetUpClientRoutes(router)
	authorization.SetUpAuthorizationRoutes(router)
	token.SetUpTokenRoutes(router)

	if err := router.Run(":8089"); err != nil {
		panic(err)
	}
}
