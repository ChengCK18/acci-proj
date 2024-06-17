package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Position struct {
	RepLatitude  float64 `bson:"repLatitude" json:"repLatitude"`
	RepLongitude float64 `bson:"repLongitude" json:"repLongitude"`
}

type Report struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	RepCategory         uint               `bson:"repCategory" json:"repCategory"`
	RepPosition         Position           `bson:"repPosition" json:"repPosition"`
	RepCreatedAt        time.Time          `bson:"repCreatedAt" json:"repCreatedAt"`
	RepIncidentDatetime time.Time          `bson:"repIncidentDatetime" json:"repOncidentDatetime"`
	RepDescription      string             `bson:"repDescription" json:"repDescription"`
	RepStatus           uint               `bson:"repStatus" json:"repStatus"`
}
