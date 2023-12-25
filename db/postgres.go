package db

import (
	"context"
	"fmt"

	"ethereum-explorer/config"
	"ethereum-explorer/logger"

	"database/sql"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

type DB struct {
	Client *sql.DB
}

type Document interface{}
type Documents []interface{}

func NewDB(ctx context.Context, cfg config.Config, colNames []string) (*DB, error) {
	dbConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		panic(err)
	}	
	
	logger.Logger.WithFields(log.Fields{
		"uri": cfg.DbHost,
		"db": cfg.DbName,
	}).Info("postgresql connecting")

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}	

	logger.Logger.WithFields(log.Fields{
		"uri": cfg.DbHost,
		"db": cfg.DbName,
	}).Info("postgresql Connected")

	return &DB{db}, nil
}

func (db *DB) Close() {
	db.Client.Close()
}