package main

import (
	"database/sql"
	"ethereum-explorer/config"
)

type Server struct {
	Db *sql.DB
	Config *config.Config
}

func NewServer(cfg config.Config, db sql.DB) {

}