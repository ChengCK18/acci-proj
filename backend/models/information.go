package models

type ContactInfo struct {
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	PhoneNum string `json:"phone,omitempty"`
}

type Information struct {
	ID              string      `json:"id"`
	InfoType        int         `json:"inf_type"`
	InfoFilepath    string      `json:"info_filepath"`
	InfoContactInfo ContactInfo `json:"info_contact_info,omitempty"`
}
