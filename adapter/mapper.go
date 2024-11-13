package adapter

import (
	"time"

	"github.com/pinkskirts/medchain-backend/domain"
	"github.com/pinkskirts/medchain-backend/ethereum"
)

func ToDomain(prescription *ethereum.PrescriptionDetails) (*domain.Prescription, error) {
	patient := domain.Patient{
		Name: prescription.Patient,
		Id:   prescription.Address,
	}
	doctor := domain.Doctor{
		Id: prescription.SigningDoctor,
	}
	expiration, err := time.Parse("2006-01-02T15:04:05.000Z", prescription.ExpirationDate)
	if err != nil {
		return nil, err
	}
	return domain.NewPrescription(
		patient,
		prescription.Medications,
		prescription.DosageInstructions,
		expiration,
		doctor,
	), nil
}

func ToEthereum(command *domain.CreatePrescriptionCommand) *ethereum.CreatePrescriptionRequest {
	return &ethereum.CreatePrescriptionRequest{
		ExpirationDate:     command.ExpirationDate.Format("2006-01-02T15:04:05.000Z"),
		Patient:            command.PatientName,
		DosageInstructions: command.DosageInstructions,
		Medications:        command.Medications,
		IsValid:            true,
	}
}
