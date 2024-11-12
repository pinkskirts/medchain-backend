package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pinkskirts/medchain-backend/api"
	"github.com/pinkskirts/medchain-backend/ethereum"
)

func main() {
	// ethereum.DeployContracts()

	client, err := ethereum.Connect()
	if err != nil {
		log.Fatal(err)
	}


	ethereum.InteractWithStorageContract(storageAddress, client)

	api.InitAPI(client, storageAddress, factoryAddress)
	api.InitRouter()
}
