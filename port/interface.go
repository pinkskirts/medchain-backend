package port

import "github.com/pinkskirts/medchain-backend/domain"

type BCPrescription interface {
	FetchPrescriptions() ([]*domain.Prescription, error)
	CreatePresecription(prescription *domain.CreatePrescriptionCommand) (id string, err error)
}
