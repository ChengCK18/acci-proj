package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactInfo struct {
	Name     string `bson:"infoName" json:"infoName"`
	Email    string `bson:"infoEmail,omitempty" json:"infoEmail,omitempty"`
	PhoneNum string `bson:"infoPhone,omitempty" json:"infoPhone,omitempty"`
}

type Information struct {
	InfoCreatedAt   time.Time   `bson:"infoCreatedAt" json:"infoCreatedAt"`
	InfoType        int         `bson:"infoType" json:"infoType"`
	InfoDescription string      `bson:"infoDescription" json:"infoDescription"`
	InfoContactInfo ContactInfo `bson:"infoContactInfo,omitempty" json:"infoContactInfo,omitempty"`
}

type InformationList struct {
	InfoID   primitive.ObjectID `bson:"_id" json:"_id"`
	InfoList []Information      `bson:"infoList" json:"infoList"`
}
