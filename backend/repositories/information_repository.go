package repositories

import (
	"acci-backend/models"
	"acci-backend/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindInformationById(id primitive.ObjectID) (models.Information, error) {
	var info models.Information

	return info, nil

}

func AddInformation(info models.Information) (*mongo.InsertOneResult, error) {
	collection := utils.MongoClient.Database("acciProj").Collection("submittedInformation")

	info.InfoCreatedAt = time.Now()

	insertResult, err := collection.InsertOne(context.TODO(), info)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func AddInformationList(infoList models.InformationList) (*mongo.InsertOneResult, error) {
	collection := utils.MongoClient.Database("acciProj").Collection("submittedInformation")

	insertResult, err := collection.InsertOne(context.TODO(), infoList)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}
