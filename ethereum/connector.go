package ethereum

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	utils "github.com/pinkskirts/medchain-backend/core/utils"
)

var (
	logger  = utils.BuildStdoutLogger("ETH Connector")
	ETH_URL = os.Getenv("ETH_URL")
)

func Connect() (*ethclient.Client, error) {
	logger.Info("Trying to connected with ETH client...")
	if ETH_URL == "" {
		logger.Error("Failed to recieve URL from the enviroment")
		return nil, errors.New("URL not provided")
	}

	client, err := ethclient.Dial(ETH_URL)
	if err != nil {
		logger.Error("Failed to connect with ETH Client")
		return nil, errors.New("Failed to connect with client")
	}
	defer client.Close()
	logger.Info("Connected with ETH client")

    // Test
    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        logger.Panic(err)
    }
    fmt.Print(chainID)

	return client, nil
}
