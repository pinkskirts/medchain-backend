package usecase

import "github.com/pinkskirts/medchain-backend/port"

type PrescriptionUseCase struct {
	bcPort port.BCPrescription
}

func NewPrescriptionUseCase(bcPort port.BCPrescription) *PrescriptionUseCase {
	return &PrescriptionUseCase{
		bcPort: bcPort,
	}
}

func (p *PrescriptionUseCase) FetchPrescription() ([]*PrescriptionDTO, error) {
	prescriptions, err := p.bcPort.FetchPrescriptions()
	if err != nil {
		return nil, err
	}
	var dtos []*PrescriptionDTO
	for _, prescription := range prescriptions {

		dtos = append(dtos, ToDTO(prescription))
	}

	return dtos, nil
}

func (p *PrescriptionUseCase) CreatePrescription(command *CreatePrescriptionRequest) (id string, err error) {
	domainCommand, err := ToDomain(command)
	if err != nil {
		return "", err
	}
	return p.bcPort.CreatePresecription(domainCommand)
}
