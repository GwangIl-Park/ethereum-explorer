package server

import (
	"database/sql"
	"ethereum-explorer/config"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Db *sql.DB
	Config *config.Config
	Gin *gin.Engine
	Timeout time.Duration
}

func NewServer(cfg *config.Config, db *sql.DB, gin *gin.Engine, timeout time.Duration) Server {
	return Server{
		db,
		cfg,
		gin,
		timeout,
	}
}

func (server *Server) Start() {
	address := fmt.Sprintf("%s:%s", server.Config.Host, server.Config.Port)
	server.Gin.Run(address)
}