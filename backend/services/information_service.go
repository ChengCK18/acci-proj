package services

import (
	"acci-backend/models"
	"acci-backend/repositories"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetInformationList() ([]models.InformationList, error) {
	infoList, err := repositories.GetInformationList()
	if err != nil {
		return nil, err
	}

	return infoList, nil
}

func AddInformationList(insertedReport *mongo.InsertOneResult) (*mongo.InsertOneResult, error) {
	insertedReportID, success := insertedReport.InsertedID.(primitive.ObjectID)
	if !success {
		return nil, fmt.Errorf("insertedReport.InsertedID is not of type primitive.ObjectID")
	}

	info := models.InformationList{
		InfoID:        insertedReportID,
		InfoList:      []models.Information{},
		InfoCreatedAt: time.Now(),
	}

	result, err := repositories.AddInformationList(info)
	if err != nil {
		return nil, err
	}

	return result, err
}

func UpdateInformationList(c *gin.Context, infoListIdStr string) (*mongo.UpdateResult, error) {
	infoListId, err := primitive.ObjectIDFromHex(infoListIdStr)
	if err != nil {
		return nil, err
	}

	var info models.Information
	if err := c.ShouldBindJSON(&info); err != nil {
		return nil, err //failed parsing
	}
	info.InfoCreatedAt = time.Now()

	result, err := repositories.UpdateInformationListById(c, infoListId, info)
	if err != nil {
		return nil, err
	}

	return result, nil

}
