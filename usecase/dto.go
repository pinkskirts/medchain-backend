package usecase

type PrescriptionDTO struct {
	Address            string   `json:"address"`
	SigningDoctor      string   `json:"signingDoctor"`
	ExpirationDate     string   `json:"expirationDate"`
	Patient            string   `json:"patient"`
	DosageInstructions string   `json:"dosageInstructions"`
	Medications        []string `json:"medications"`
	IsValid            bool     `json:"isValid"`
}

type CreatePrescriptionRequest struct {
	ExpirationDate     string   `json:"expirationDate"`
	Patient            string   `json:"patient"`
	DosageInstructions string   `json:"dosageInstructions"`
	Medications        []string `json:"medications"`
	IsValid            bool     `json:"isValid"`
}
