package services

import (
	"acci-backend/models"
	"acci-backend/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetReportService() (interface{}, error) {
	reports, err := repositories.GetAllReports()
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func AddReportService(report models.Report) (*mongo.InsertOneResult, error) {
	insertedResult, err := repositories.AddReport(report)
	if err != nil {
		return nil, err
	}
	return insertedResult, err
}
