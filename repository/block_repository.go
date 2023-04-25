package repository

import "database/sql"

type blockRepository struct {
	db *sql.DB
}

func NewBlockRepository(db *sql.DB) 