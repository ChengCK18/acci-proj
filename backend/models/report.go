package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Position struct {
	Longitude string `json:"rep_longitude"`
	Latitude  string `json:"rep_latitude"`
}

type Report struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"` //omitempty because POST request will not provide this attr, mongodb will generate this
	RepType             uint               `json:"rep_type"`
	RepAffected         string             `json:"rep_affected"`
	RepPosition         Position           `json:"rep_position"`
	RepCreatedAt        time.Time          `json:"rep_created_at"`
	RepIncidentDatetime time.Time          `json:"rep_incident_datetime"`
	RepDescription      string             `json:"rep_description"`
	RepStatus           uint               `json:"rep_status"`
}
