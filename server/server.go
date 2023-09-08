package server

import (
	"ethereum-explorer/config"
	"ethereum-explorer/db"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/subscriber"
	"net/http"
	"time"

	"github.com/rs/cors"
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
	handler := cors.Default().Handler(router)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://foo.com", "http://foo.com:8080"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler = corsMiddleware.Handler(handler)

	http.ListenAndServe(server.Config.Url, handler)
}
