package main

import (
	"log"
	"github.com/MaulanaAhmadSulami/juke_test.git/internal/config"
	"github.com/MaulanaAhmadSulami/juke_test.git/internal/db"
)

func testConnection(){
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("trying to connect with:", cfg.GetDBConnectionString())

	database, err := db.NewPostgresDB(cfg.GetDBConnectionString(), cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	log.Println("connected")
}