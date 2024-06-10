package repositories

import (
	"acci-backend/models"
	"acci-backend/utils"
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllReports() ([]models.Report, error) {
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

func AddReport(report models.Report) (*mongo.InsertOneResult, error) {
	collection := utils.MongoClient.Database("acciProj").Collection("submittedReport")

	report.ID = strconv.FormatInt(time.Now().UnixNano(), 10) // TODO: replace this later. Only for temp ids now
	report.RepCreatedAt = time.Now()

	insertResult, err := collection.InsertOne(context.TODO(), report)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}
