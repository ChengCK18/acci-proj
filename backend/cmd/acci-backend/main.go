package main

import (
	"acci-backend/config"
	"acci-backend/routes"
	"acci-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	router := gin.Default()
	routes.SetupRoutes(router)

	utils.InitMongoDB()

	router.Run()
}
