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

	storageAddress := common.HexToAddress("0x186D9cc85dD0F4f1F93Fc6E1326098312ACd5566")
	factoryAddress := common.HexToAddress("0x9d04576D334D06f604A2Af226650B603ecf278A9")

	ethereum.InteractWithStorageContract(storageAddress, client)
	ethereum.InteractWithFactoryContract(factoryAddress, client)

	api.InitAPI(client, storageAddress, factoryAddress)
	api.InitRouter()
}
