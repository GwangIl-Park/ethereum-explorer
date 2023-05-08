package routes

import (
	"ethereum-explorer/server"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "http://foo.com:3000")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")

			if c.Request.Method == "OPTIONS" {
					c.AbortWithStatus(204)
					return
			}

			c.Next()
	}
}

func Setup(server *server.Server) {
	publicRouter := server.Gin.Group("")

	NewBlockRouter(server.Timeout, server.Db, publicRouter)
	NewTransactionRouter(server.Timeout, server.Db, publicRouter)
}