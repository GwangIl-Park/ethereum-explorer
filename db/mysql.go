package db

import (
	"database/sql"
	"ethereum-explorer/logger"
)

func New() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		logger.Logger.Error("db error")
		return nil, err
	}
	
	logger.Logger.Info("db start")
	return db, nil
}