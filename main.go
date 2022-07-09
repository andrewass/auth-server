package main

import (
	"auth-server/authorization"
	"auth-server/client"
	"auth-server/common"
	"auth-server/discovery"
	"auth-server/token"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
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
	discovery.SetUpDiscoveryRoutes(router)
	client.SetUpClientRoutes(router)
	authorization.SetUpAuthorizationRoutes(router)
	token.SetUpTokenRoutes(router)

	if err := router.Run(":8089"); err != nil {
		panic(err)
	}
}
