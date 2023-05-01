package db

import (
	"database/sql"
	"ethereum-explorer/config"
	"ethereum-explorer/logger"
	"fmt"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s/%s)", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		logger.Logger.Error("db error")
		return nil, err
	}
	
	logger.Logger.Info("db start")
	return db, nil
}