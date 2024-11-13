// api/router.go

package api

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/julienschmidt/httprouter"
	"github.com/pinkskirts/medchain-backend/adapter"
	"github.com/pinkskirts/medchain-backend/controler"
	"github.com/pinkskirts/medchain-backend/core/utils"
	"github.com/pinkskirts/medchain-backend/usecase"
)

const (
	API_V1 = "/api/v1"
)

var logger = utils.BuildStdoutLogger("Router")

// api/router.go

func gettPrescriptions(controler *controler.PrescriptionController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		controler.GetPrescription(w, r)
	}
}

func createPrescription(controler *controler.PrescriptionController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		controler.CreatePrescription(w, r)
	}
}

func InitRouter(client *ethclient.Client, storageAddr, factoryAddr common.Address) error {
	logger.Info("Initializing router")
	router := httprouter.New()

	adapter, err := adapter.NewEthereumPrescription(client, storageAddr, factoryAddr)
	if err != nil {
		return err
	}

	usecase := usecase.NewPrescriptionUseCase(adapter)
	controler := controler.NewPrescriptionController(usecase)

	router.GET(API_V1+"/prescription/", CORS(gettPrescriptions(controler)))
	router.POST(API_V1+"/prescription/", CORS(createPrescription(controler)))
	router.OPTIONS(API_V1+"/prescription/", CORS(gettPrescriptions(controler)))
	router.OPTIONS(API_V1+"/prescription/", CORS(createPrescription(controler)))

	logger.Fatal(http.ListenAndServe(":6969", router))
	return nil
}

// CORS middleware function
func CORS(next http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
