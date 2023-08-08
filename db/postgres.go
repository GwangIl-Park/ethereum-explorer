package db

import (
	"context"
	"fmt"

	"ethereum-explorer/config"
	"ethereum-explorer/logger"

	"database/sql"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	Client *sql.DB
	
	Collections map[string]*mongo.Collection
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

	return &DB{db, table}, nil
}

func (db *DB) InsertOneDocument(colName string, document Document) error {
	_, err := db.Collections[colName].InsertOne(context.TODO(), document)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) InsertManyDocument(colName string, documents Documents) error {
	_, err := db.Collections[colName].InsertMany(context.TODO(), documents)
	if err != nil {
		return err
	}

	return nil
}