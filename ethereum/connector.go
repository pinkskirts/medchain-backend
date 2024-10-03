package ethereum

import (
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pinkskirts/medchain-backend/core/utils"
)

var (
	logger  = utils.BuildStdoutLogger("ETH Connector")
	ETH_URL = os.Getenv("ETH_URL")
    WALLET_ADDRESS = os.Getenv("MAIN_WALLET_ADDRESS")
    WALLET_KEY = os.Getenv("MAIN_WALLET_PRIVATE_KEY") // This is the worst
)

func Connect() (*ethclient.Client, error) {
	logger.Info("Trying to connected with ETH client...")
	if ETH_URL == "" || WALLET_ADDRESS == "" || WALLET_KEY == "" {
		logger.Error("Failed to recieve Eth account information from the enviroment")
		return nil, errors.New("ENVs not provided")
	}

	client, err := ethclient.Dial(ETH_URL)
	if err != nil {
		logger.Error("Failed to connect with ETH Client")
		return nil, errors.New("Failed to connect with client")
	}
	defer client.Close()
	logger.Info("Connected with ETH client")

	return client, nil
}
