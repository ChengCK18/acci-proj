package services

import (
	"acci-backend/models"
	"acci-backend/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetInformationByIdService(id primitive.ObjectID) (interface{}, error) {
	info, err := repositories.FindInformationById(id)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func InitializeInformationService(id primitive.ObjectID) (*mongo.InsertOneResult, error) {

	info := models.InformationList{
		InfoID:        id,
		InfoList:      []models.Information{},
		InfoCreatedAt: time.Now(),
	}

	insertedResult, err := repositories.AddInformationList(info)
	if err != nil {
		return nil, err
	}
	return insertedResult, err
}
