package controllers

import (
	"acci-backend/models"
	"acci-backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	insertedReportID, success := insertedReport.InsertedID.(primitive.ObjectID)
	if !success {
		fmt.Println("Error: insertedReport.InsertedID is not of type primitive.ObjectID")
		return
	}

	insertedInformation, err := services.InitializeInformationService(insertedReportID)
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
