package controllers

import (
	"acci-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateInformationList(c *gin.Context) {
	infoListIdStr := c.Param("id")

	_, err := services.UpdateInformationList(c, infoListIdStr) // pass in the relevant
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Information added successfully",
	})
}

func GetInformationList(c *gin.Context) {
	infoList, err := services.GetInformationList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, infoList)
}
