package models

import (
	"time"
)

type Position struct {
	Longitude string `json:"rep_longitude"`
	Latitude  string `json:"rep_latitude"`
}

type Report struct {
	ID                  string        `json:"_id"`
	RepType             uint          `json:"rep_type"`
	RepAffected         string        `json:"rep_affected"`
	RepPosition         Position      `json:"rep_position"`
	RepCreatedAt        time.Time     `json:"rep_created_at"`
	RepIncidentDatetime time.Time     `json:"rep_incident_datetime"`
	RepDescription      string        `json:"rep_description"`
	RepStatus           uint          `json:"rep_status"`
	RepInfo             []Information `json:"rep_info,omitempty"`
}
