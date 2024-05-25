package models

import "time"

type Position struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type Report struct {
	ID                  string    `json:"id"`
	RepType             uint      `json:"rep_type"`
	RepAffected         string    `json:"rep_affected"`
	RepPosition         Position  `json:"rep_position"`
	RepCreatedAt        time.Time `json:"rep_created_at"`
	RepIncidentDatetime time.Time `json:"rep_incident_datetime"`
	RepDescription      string    `json:"rep_description"`
	RepStatus           uint      `json:"rep_status"`
}
