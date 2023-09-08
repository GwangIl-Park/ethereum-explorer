package server

import (
	"ethereum-explorer/config"
	"ethereum-explorer/db"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/middlewares/handler"
	"ethereum-explorer/subscriber"
	"net/http"
	"time"
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

func (server *Server) Start(errorChan chan error, router *http.ServeMux) error {
	err := http.ListenAndServe(server.Config.Url, handler.GetHandler(router))
	if err != nil {
		return err
	}
	return nil
}
