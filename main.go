package main

import (
	"auth-server/client"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	client.SetUpClientRoutes(router)
	router.Run(":8089")
}
