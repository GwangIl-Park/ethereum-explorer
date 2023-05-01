package routes

import (
	"ethereum-explorer/server"
)

func Setup(server *server.Server) {
	publicRouter := server.Gin.Group("")

	NewBlockRouter(server.Timeout, server.Db, publicRouter)
	NewTransactionRouter(server.Timeout, server.Db, publicRouter)
}