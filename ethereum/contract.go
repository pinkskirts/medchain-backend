package ethereum

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pinkskirts/medchain-backend/contracts"
)

func DeployContracts() {
	// Connect to the Ethereum client
	client, err := Connect()
	if err != nil {
		logger.Fatal(err)
	}

	// Load the private key
	pk, err := crypto.HexToECDSA(WALLET_KEY)
	if err != nil {
		logger.Fatal(err)
	}

	// Get the public key and address
	pubk := pk.Public()
	pubkECDSA, ok := pubk.(*ecdsa.PublicKey)
	if !ok {
		logger.Fatal("Failed to cast public key to ECDSA")
	}
	fromAddr := crypto.PubkeyToAddress(*pubkECDSA)

	// Get the nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		logger.Fatal(err)
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Fatal(err)
	}

	// Get the chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		logger.Fatal(err)
	}

	// Create a keyed transactor
	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		logger.Fatal(err)
	}

	auth.GasLimit = uint64(6000000)
	auth.GasPrice = gasPrice

	// Read the Storage contract's ABI and bytecode from files
	storageABI, err := os.ReadFile("build/Storage.abi")
	if err != nil {
		logger.Fatal(err)
	}

	storageBin, err := os.ReadFile("build/Storage.bin")
	if err != nil {
		logger.Fatal(err)
	}

	// Parse the Storage contract's ABI
	storageAbiParsed, err := abi.JSON(strings.NewReader(string(storageABI)))
	if err != nil {
		logger.Fatal(err)
	}

	// Deploy the Storage contract
	auth.Nonce = big.NewInt(int64(nonce))
	storageAddr, tx, _, err := bind.DeployContract(
		auth,
		storageAbiParsed,
		common.FromHex(string(storageBin)),
		client,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Storage contract deployed at:", storageAddr.Hex())
	fmt.Println("Transaction hash:", tx.Hash().Hex())

	// Wait for the transaction to be mined (optional)
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}

	// Increment the nonce for the next transaction
	nonce++

	// Read the Factory contract's ABI and bytecode from files
	factoryABI, err := os.ReadFile("build/Factory.abi")
	if err != nil {
		log.Fatal(err)
	}

	factoryBin, err := os.ReadFile("build/Factory.bin")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the Factory contract's ABI
	factoryAbiParsed, err := abi.JSON(strings.NewReader(string(factoryABI)))
	if err != nil {
		log.Fatal(err)
	}

	// Deploy the Factory contract, passing the Storage contract's address to the constructor
	auth.Nonce = big.NewInt(int64(nonce))
	factoryAddr, tx, _, err := bind.DeployContract(
		auth,
		factoryAbiParsed,
		common.FromHex(string(factoryBin)),
		client,
		storageAddr, // Constructor parameter for Factory contract
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Factory contract deployed at:", factoryAddr.Hex())
	fmt.Println("Transaction hash:", tx.Hash().Hex())
}

func InteractWithFactoryContract(factoryAddress common.Address, client *ethclient.Client) {
	// Instantiate the Factory contract
	factoryInstance, err := contracts.NewFactory(factoryAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// For methods that modify the state, use a transact
	auth, err := PrepareAuth(client)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare parameters for the createPrescription function
	expirationDate := "2025-01-01"
	patient := "John Doe"
	dosageInstructions := "Take one pill daily"
	medications := []string{"MedicationA", "MedicationB"}
	isValid := true

	tx, err := factoryInstance.CreatePrescription(
		auth,
		expirationDate,
		patient,
		dosageInstructions,
		medications,
		isValid,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction submitted:", tx.Hash().Hex())

	// Optionally, wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction mined in block", receipt.BlockNumber)
}

func PrepareAuth(client *ethclient.Client) (*bind.TransactOpts, error) {
	pk, err := crypto.HexToECDSA(WALLET_KEY)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		return nil, err
	}

	// Customizar as opções de transação
	auth.GasLimit = uint64(6000000)

	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	return auth, nil
}

type PrescriptionDetails struct {
	Address            string   `json:"address"`
	SigningDoctor      string   `json:"signingDoctor"`
	ExpirationDate     string   `json:"expirationDate"`
	Patient            string   `json:"patient"`
	DosageInstructions string   `json:"dosageInstructions"`
	Medications        []string `json:"medications"`
	IsValid            bool     `json:"isValid"`
}

func InteractWithStorageContract(storageAddress common.Address, client *ethclient.Client) ([]*PrescriptionDetails, error) {
	// Instantiate the Storage contract
	storageInstance, err := contracts.NewStorage(storageAddress, client)
	if err != nil {
		return nil, err
	}

	// Get all stored prescription addresses
	prescriptionAddresses, err := storageInstance.GetAllPrescriptionAddresses(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	// For each prescription address, get the details
	var prescriptions []*PrescriptionDetails
	for _, prescriptionAddress := range prescriptionAddresses {
		// Get the details of the prescription
		details, err := GetPrescriptionDetailsByAddress(prescriptionAddress, client)
		if err != nil {
			log.Printf("Error getting prescription details at %s: %v", prescriptionAddress.Hex(), err)
			continue
		}

		prescriptions = append(prescriptions, details)
	}

	// Return the list of prescriptions
	return prescriptions, nil
}

func GetPrescriptionDetailsByAddress(prescriptionAddress common.Address, client *ethclient.Client) (*PrescriptionDetails, error) {
	// Instantiate the Prescription contract
	prescriptionInstance, err := contracts.NewPrescription(prescriptionAddress, client)
	if err != nil {
		return nil, err
	}

	// Call getPrescriptionDetails to get all details
	signingDoctor, expirationDate, patient, dosageInstructions, medications, isValid, err := prescriptionInstance.GetPrescriptionDetails(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	// Build the PrescriptionDetails struct
	details := &PrescriptionDetails{
		Address:            prescriptionAddress.Hex(),
		SigningDoctor:      signingDoctor.Hex(),
		ExpirationDate:     expirationDate,
		Patient:            patient,
		DosageInstructions: dosageInstructions,
		Medications:        medications,
		IsValid:            isValid,
	}

	return details, nil
}

type CreatePrescriptionRequest struct {
	ExpirationDate     string   `json:"expirationDate"`
	Patient            string   `json:"patient"`
	DosageInstructions string   `json:"dosageInstructions"`
	Medications        []string `json:"medications"`
	IsValid            bool     `json:"isValid"`
}

func CreatePrescription(factoryAddress common.Address, client *ethclient.Client, request *CreatePrescriptionRequest) (common.Address, error) {
	// Instanciar o contrato Factory
	factoryInstance, err := contracts.NewFactory(factoryAddress, client)
	if err != nil {
		return common.Address{}, err
	}

	// Preparar o transactor
	auth, err := PrepareAuth(client)
	if err != nil {
		return common.Address{}, err
	}

	// Chamar a função CreatePrescription do contrato Factory
	tx, err := factoryInstance.CreatePrescription(
		auth,
		request.ExpirationDate,
		request.Patient,
		request.DosageInstructions,
		request.Medications,
		request.IsValid,
	)
	if err != nil {
		return common.Address{}, err
	}

	fmt.Println("Transação submetida:", tx.Hash().Hex())

	// Aguardar a transação ser minerada
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return common.Address{}, err
	}

	fmt.Println("Transação minerada no bloco", receipt.BlockNumber)

	// Obter o endereço da nova prescrição a partir dos logs de eventos
	// Assumindo que o contrato Factory emite um evento PrescriptionCreated(address)
	factoryAbi, err := abi.JSON(strings.NewReader(contracts.FactoryABI))
	if err != nil {
		return common.Address{}, err
	}

	// Hash da assinatura do evento
	eventSignature := []byte("PrescriptionCreated(address)")
	eventSignatureHash := crypto.Keccak256Hash(eventSignature)

	var prescriptionAddress common.Address

	for _, vLog := range receipt.Logs {
		if len(vLog.Topics) > 0 && vLog.Topics[0] == eventSignatureHash {
			// Encontrou o evento
			var eventData struct {
				PrescriptionAddress common.Address
			}

			err := factoryAbi.UnpackIntoInterface(&eventData, "PrescriptionCreated", vLog.Data)
			if err != nil {
				return common.Address{}, err
			}

			// O endereço da prescrição pode estar em Topics[1]
			if len(vLog.Topics) > 1 {
				prescriptionAddress = common.HexToAddress(vLog.Topics[1].Hex())
			} else {
				prescriptionAddress = eventData.PrescriptionAddress
			}
			break
		}
	}

	if prescriptionAddress == (common.Address{}) {
		return common.Address{}, errors.New("evento PrescriptionCreated não encontrado nos logs")
	}

	return prescriptionAddress, nil
}
