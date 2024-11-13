package usecase

import (
	"time"

	"github.com/pinkskirts/medchain-backend/domain"
)

func ToDTO(domain *domain.Prescription) *PrescriptionDTO {
	return &PrescriptionDTO{
		Address:            domain.Patient.Id,
		SigningDoctor:      domain.Doctor.Id,
		ExpirationDate:     domain.ExpirationDate.Format("2006-01-02T15:04:05.000Z"),
		Patient:            domain.Patient.Name,
		DosageInstructions: domain.Description,
		Medications:        domain.Medications,
		IsValid:            true,
	}
}

func ToDomain(command *CreatePrescriptionRequest) (*domain.CreatePrescriptionCommand, error) {
	expiration, err := time.Parse("2006-01-02T15:04:05.000Z", command.ExpirationDate)
	if err != nil {
		return nil, err
	}
	return &domain.CreatePrescriptionCommand{
		ExpirationDate:     expiration,
		PatientName:        command.Patient,
		DosageInstructions: command.DosageInstructions,
		Medications:        command.Medications,
	}, nil
}
