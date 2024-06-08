package routes

import (
	"acci-backend/controllers"
	"acci-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	reportGroup := router.Group("api/report")
	{
		reportGroup.GET("/:id", middleware.LoggerMiddleware(), controllers.GetReport)
	}
}
