package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactInfo struct {
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	PhoneNum string `json:"phone,omitempty"`
}

type Information struct {
	InfoCreatedAt    time.Time   `json:"info_created_at"`
	InfoType         int         `json:"inf_type"`
	InfoFilepath     string      `json:"info_filepath"`
	InfoContactInfo  ContactInfo `json:"info_contact_info,omitempty"`
	InfoLinkedReport string      `json:"info_linked_report_id"`
}

type InformationList struct {
	InfoID        primitive.ObjectID `bson:"_id"`
	InfoList      []Information      `json:""`
	InfoCreatedAt time.Time          `json:"info_created_at"`
}
