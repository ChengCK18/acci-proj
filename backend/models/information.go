package models

type ContactInfo struct {
	Name     string `bson:"name"`
	Email    string `bson:"email,omitempty"`
	PhoneNum string `bson:"phone,omitempty"`
}

type Information struct {
	ID               string      `bson:"id"`
	InfoType         int         `bson:"inf_type"`
	InfoFilepath     string      `bson:"info_filepath"`
	InfoContactInfo  ContactInfo `bson:"info_contact_info,omitempty"`
	InfoLinkedReport Report      `bson:"info_linked_report"`
}
