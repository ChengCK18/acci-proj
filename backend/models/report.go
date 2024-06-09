package models

import (
	"time"
)

type Position struct {
	Longitude string `bson:"rep_longitude"`
	Latitude  string `bson:"rep_latitude"`
}

type Report struct {
	ID                  string        `bson:"_id"`
	RepType             uint          `bson:"rep_type"`
	RepAffected         string        `bson:"rep_affected"`
	RepPosition         Position      `bson:"rep_position"`
	RepCreatedAt        time.Time     `bson:"rep_created_at"`
	RepIncidentDatetime time.Time     `bson:"rep_incident_datetime"`
	RepDescription      string        `bson:"rep_description"`
	RepStatus           uint          `bson:"rep_status"`
	RepInfo             []Information `bson:"rep_info,omitempty"`
}
