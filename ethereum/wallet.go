package ethereum

import (
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

type Wallet struct {
	address    *common.Address
	url        *accounts.URL
	walletInfo *[]byte
}

func GenerateWallet(passfrase string) (*Wallet, error) {
	logger.Info("Generating a keystore")
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(passfrase)
	if err != nil {
		return nil, errors.New("Failed to create a keystore")
	}
	file, err := os.ReadFile(account.URL.String())
	if err != nil {
		return nil, errors.New("Failed to create a keystore")
	}
	return &Wallet{&account.Address, &account.URL, &file}, nil
}
