package domain

import "time"

type Patient struct {
	Name string
	Id   string
}

type Doctor struct {
	Id string
}

type Prescription struct {
	Patient        Patient
	Doctor         Doctor
	Medications    []string
	Description    string
	ExpirationDate time.Time
}

type CreatePrescriptionCommand struct {
	ExpirationDate     time.Time
	PatientName        string
	DosageInstructions string
	Medications        []string
}

func NewPrescription(
	patient Patient,
	medications []string,
	description string,
	expiration time.Time,
	doctor Doctor,
) *Prescription {
	return &Prescription{
		Patient:        patient,
		Medications:    medications,
		Description:    description,
		ExpirationDate: expiration,
		Doctor:         doctor,
	}
}
