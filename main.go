package main

import (
	"auth-server/client"
	"auth-server/persistence"
	"github.com/gin-gonic/gin"
)

func main() {
	configurePersistenceLayer()
	configureRoutes()
}

func configurePersistenceLayer() {
	persistence.ConnectToDatabase()
}

func configureRoutes() {
	router := gin.Default()
	client.SetUpClientRoutes(router)
	router.Run(":8089")
}
