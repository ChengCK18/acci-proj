package controllers

import (
	"acci-backend/models"
	"acci-backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReport(c *gin.Context) {
	report, err := services.GetReportService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // terminate
	}

	c.JSON(http.StatusOK, report)
}

func CreateReport(c *gin.Context) {
	var report models.Report

	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Parsed result here => ", report)

	insertedReport, err := services.AddReportService(report)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the inserted result including the inserted ID
	c.JSON(http.StatusCreated, gin.H{
		"message": "Report created successfully",
		"id":      insertedReport.InsertedID,
	})
}
