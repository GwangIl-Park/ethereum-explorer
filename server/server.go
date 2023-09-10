package server

import (
	"ethereum-explorer/config"
	"ethereum-explorer/db"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/logger"
	"ethereum-explorer/middleware/handler"
	"ethereum-explorer/subscriber"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	Db        *db.DB
	Config    *config.Config
	EthClient *ethClient.EthClient
	Sub       *subscriber.Subscriber
	Timeout   time.Duration
}

func NewServer(db *db.DB, cfg *config.Config, ethClient *ethClient.EthClient, sub *subscriber.Subscriber, timeout time.Duration) Server {
	return Server{
		db,
		cfg,
		ethClient,
		sub,
		timeout,
	}
}

func (server *Server) Start(errorChan chan error, router *http.ServeMux) {
	logger.Logger.WithFields(logrus.Fields{
		"URL": server.Config.Url,
	}).Info("Server Start")

	err := http.ListenAndServe(server.Config.Url, handler.GetHandler(router))
	if err != nil {
		logger.Logger.WithError(err).Error("Server Error")
		errorChan <- err
	}
}
