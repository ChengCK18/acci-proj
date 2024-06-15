package services

import (
	"acci-backend/models"
	"acci-backend/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetReportService() ([]models.Report, error) {
	reports, err := repositories.GetAllReports()
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func AddReport(c *gin.Context) (*mongo.InsertOneResult, error) {
	var report models.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		return nil, err
	}

	insertedResult, err := repositories.AddReport(report)
	if err != nil {
		return nil, err
	}
	return insertedResult, err
}
