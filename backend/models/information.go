package models

type ContactInfo struct {
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	PhoneNum string `json:"phone,omitempty"`
}

type Information struct {
	ID               string      `bson:"id"`
	InfoType         int         `json:"inf_type"`
	InfoFilepath     string      `json:"info_filepath"`
	InfoContactInfo  ContactInfo `json:"info_contact_info,omitempty"`
	InfoLinkedReport string      `json:"info_linked_report_id"`
}
