package controllers

import (
	"acci-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetReport(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid report ID"})
		return // terminate
	}

	report, err := services.GetReportService(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // terminate
	}

	c.JSON(http.StatusOK, report)
}
