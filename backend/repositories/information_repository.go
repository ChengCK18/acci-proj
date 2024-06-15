package repositories

import (
	"acci-backend/models"
	"acci-backend/utils"
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetInformationList() ([]models.InformationList, error) {
	var infoList []models.InformationList
	collection := utils.MongoClient.Database("acciProj").Collection("submittedInformation")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &infoList); err != nil {
		return nil, err
	}
	return infoList, nil

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

	result, err := collection.InsertOne(context.TODO(), infoList)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateInformationListById(c *gin.Context, infoListId primitive.ObjectID, info models.Information) (*mongo.UpdateResult, error) {
	collection := utils.MongoClient.Database("acciProj").Collection("submittedInformation")

	update := bson.M{"$push": bson.M{"infoList": info}}
	filter := bson.M{"_id": infoListId}

	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 { //UpdateOne does not retunr error if no record match filter, need to check here :O
		return nil, fmt.Errorf("no record found with ID: %s", infoListId.Hex())
	}

	return result, nil
}
