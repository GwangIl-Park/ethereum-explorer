package server

import (
	"database/sql"
	"ethereum-explorer/config"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/subscriber"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Db *sql.DB
	Config *config.Config
	Gin *gin.Engine
	EthClient *ethClient.EthClient
	Sub *subscriber.Subscriber
	Timeout time.Duration
}

func NewServer(db *sql.DB, cfg *config.Config, gin *gin.Engine, ethClient *ethClient.EthClient, sub *subscriber.Subscriber, timeout time.Duration) Server {
	return Server{
		db,
		cfg,
		gin,
		ethClient,
		sub,
		timeout,
	}
}

func (server *Server) Start() {
	address := fmt.Sprintf("%s:%s", server.Config.Host, server.Config.Port)
	server.Gin.Run(address)
}