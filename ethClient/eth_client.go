package ethClient

import (
	"ethereum-explorer/config"
	"ethereum-explorer/logger"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type EthClient struct {
	Http *ethclient.Client
	Ws *ethclient.Client
}

func NewEthClient(cfg *config.Config) (*EthClient, error) {
	logger.Logger.WithFields(logrus.Fields{
		"Http": cfg.ChainHttp,
		"Ws": cfg.ChainWs,
	}).Info("Connecting Eth Client")

	http, err := ethclient.Dial(cfg.ChainHttp)
	if err != nil {
		return nil, err
	}

	ws, err := ethclient.Dial(cfg.ChainWs)
	if err != nil {
		return nil, err
	}

	logger.Logger.WithFields(logrus.Fields{
		"Http": cfg.ChainHttp,
		"Ws": cfg.ChainWs,
	}).Info("Connecting Eth Client")

	return &EthClient{
		http,
		ws,
	}, nil
}