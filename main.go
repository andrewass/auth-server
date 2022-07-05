package main

import (
	"auth-server/authorization"
	"auth-server/client"
	"auth-server/persistence"
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
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func configurePersistenceLayer() {
	persistence.ConnectToDatabase()
}

func configureRoutes() {
	router := gin.Default()

	client.SetUpClientRoutes(router)
	authorization.SetUpAuthorizationRoutes(router)
	token.SetUpTokenRoutes(router)

	err := router.Run(":8089")
	if err != nil {
		panic(err)
	}
}
