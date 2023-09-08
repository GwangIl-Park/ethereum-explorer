package router

import (
	"ethereum-explorer/server"
	"net/http"
)

func Setup(server *server.Server, router *http.ServeMux) {
	NewBlockRouter(server.Timeout, server.Db, router)
	NewTransactionRouter(server.Timeout, server.Db, router)
}
