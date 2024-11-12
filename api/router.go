// api/router.go

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/julienschmidt/httprouter"
	"github.com/pinkskirts/medchain-backend/core/utils"
	"github.com/pinkskirts/medchain-backend/ethereum"
)

const (
	API_V1 = "/api/v1"
)

var (
	logger         = utils.BuildStdoutLogger("Router")
	storageAddress common.Address
	factoryAddress common.Address
	ethClient      *ethclient.Client
)

// Initialize the Ethereum client and contract addresses
func InitAPI(client *ethclient.Client, storageAddr, factoryAddr common.Address) {
	ethClient = client
	storageAddress = storageAddr
	factoryAddress = factoryAddr
}

// api/router.go

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Lidar com requisições OPTIONS
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodGet:
		prescriptions, err := getPrescriptions()
		if err != nil {
			http.Error(w, "Failed to get prescriptions", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(prescriptions)

	case http.MethodPost:
		// Ler o corpo da requisição para obter os detalhes da prescrição
		var request ethereum.CreatePrescriptionRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Chamar a função para criar a prescrição
		prescriptionAddress, err := ethereum.CreatePrescription(factoryAddress, ethClient, &request)
		if err != nil {
			log.Println("Error creating prescription:", err)
			http.Error(w, "Failed to create prescription", http.StatusInternalServerError)
			return
		}

		// Retornar resposta de sucesso com o endereço da prescrição
		response := map[string]string{
			"message":             "Prescription created successfully",
			"prescriptionAddress": prescriptionAddress.Hex(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func InitRouter() {
	logger.Info("Initializing router")
	router := httprouter.New()

	router.GET(API_V1+"/", Index)
	router.POST(API_V1+"/", Index)
	router.OPTIONS(API_V1+"/", Index)

	logger.Fatal(http.ListenAndServe(":6969", router))
}

func getPrescriptions() ([]*ethereum.PrescriptionDetails, error) {
	if ethClient == nil {
		return nil, fmt.Errorf("ethereum client is not initialized")
	}

	prescriptions, err := ethereum.InteractWithStorageContract(storageAddress, ethClient)
	if err != nil {
		log.Println("Error getting prescriptions:", err)
		return nil, err
	}

	return prescriptions, nil
}
