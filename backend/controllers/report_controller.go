package controllers

import (
	"acci-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReport(c *gin.Context) {
	report, err := services.GetReportService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

func AddReport(c *gin.Context) {
	insertedReport, err := services.AddReport(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	insertedInformation, err := services.AddInformationList(insertedReport)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the inserted result including the inserted ID
	c.JSON(http.StatusCreated, gin.H{
		"message":   "Report created successfully",
		"report_id": insertedReport.InsertedID,
		"info_id":   insertedInformation.InsertedID,
	})
}
