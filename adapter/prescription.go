package adapter

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pinkskirts/medchain-backend/domain"
	"github.com/pinkskirts/medchain-backend/ethereum"
	"github.com/pinkskirts/medchain-backend/port"
)

type EthereumPrescription struct {
	client      *ethclient.Client
	storageAddr common.Address
	factoryAddr common.Address
}

func NewEthereumPrescription(
	client *ethclient.Client,
	storageAddr, factoryAddr common.Address,
) (*EthereumPrescription, error) {
	if client == nil {
		return nil, fmt.Errorf("ethereum client is not initialized")
	}
	return &EthereumPrescription{
		client:      client,
		storageAddr: storageAddr,
		factoryAddr: factoryAddr,
	}, nil
}

var _ port.BCPrescription = &EthereumPrescription{}

func (e *EthereumPrescription) FetchPrescriptions() (p []*domain.Prescription, er error) {

	prescriptions, err := ethereum.InteractWithStorageContract(e.storageAddr, e.client)
	if err != nil {
		log.Println("Error getting prescriptions:", err)
		return nil, err
	}
	var domainPrescription []*domain.Prescription
	for _, prescription := range prescriptions {
		val, err := ToDomain(prescription)
		if err != nil {
			er = err
			return
		}
		domainPrescription = append(domainPrescription, val)
	}

	return domainPrescription, nil
}

func (e *EthereumPrescription) CreatePresecription(domain *domain.CreatePrescriptionCommand) (id string, err error) {
	request := ToEthereum(domain)
	prescriptionAddress, err := ethereum.CreatePrescription(e.factoryAddr, e.client, request)
	if err != nil {
		log.Println("Error creating prescription:", err)
		return "", err
	}
	id = string(prescriptionAddress.Bytes())
	return id, nil
}
