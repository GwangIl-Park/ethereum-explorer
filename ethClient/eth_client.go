package ethClient

import (
	"ethereum-explorer/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	Eth *ethclient.Client
}

func NewEthClient(cfg *config.Config) *EthClient {
	client, err := ethclient.Dial(cfg.ChainUrl)
	if err != nil {

	}
	return &EthClient{
		client,
	}
}