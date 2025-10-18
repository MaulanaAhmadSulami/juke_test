package db

import (
	"database/sql"
	"fmt"
	"log"
	cfg "github.com/MaulanaAhmadSulami/juke_test.git/internal/config"

	_ "github.com/lib/pq"
)

func NewPostgresDB(connStr string, dbConf cfg.DbConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open datbase: %w", err);
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err);
	}

	db.SetMaxOpenConns(dbConf.MaxOpenConns)
	db.SetMaxIdleConns(dbConf.MaxIdleConns)

	log.Println("ping connedct");
	return db, nil
}