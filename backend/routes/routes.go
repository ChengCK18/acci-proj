package routes

import (
	"acci-backend/controllers"
	"acci-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	reportGroup := router.Group("api/report")
	{
		reportGroup.GET("/", middleware.LoggerMiddleware(), controllers.GetReport)
		reportGroup.POST("/", middleware.LoggerMiddleware(), controllers.AddReport)
	}

	informationGroup := router.Group("api/information")
	{
		informationGroup.GET("/", middleware.LoggerMiddleware(), controllers.GetInformationList)
		informationGroup.PATCH("/:id", middleware.LoggerMiddleware(), controllers.UpdateInformationList)
	}
}
