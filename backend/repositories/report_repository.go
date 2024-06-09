package repositories

import (
	"acci-backend/models"
	"acci-backend/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func FindReportById(id uint) ([]models.Report, error) {
	var reports []models.Report

	collection := utils.MongoClient.Database("acciProj").Collection("submittedReport")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &reports); err != nil {
		panic(err)
	}

	return reports, nil

}
