package main

import (
	"acci-backend/config"
	"acci-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run()
}
