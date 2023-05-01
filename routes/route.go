package routes

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(timeout *time.Duration, db *sql.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")

	NewBlockRouter(*timeout, db, publicRouter)
	NewTransactionRouter(*timeout, db, publicRouter)
}