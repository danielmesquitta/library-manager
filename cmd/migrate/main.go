package main

import (
	"log"

	"github.com/danielmesquitta/library-manager/internal/database"
	"github.com/danielmesquitta/library-manager/internal/model"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}

	if err := db.AutoMigrate(
		&model.Author{},
		&model.Book{},
		&model.Borrower{},
		&model.Loan{},
	); err != nil {
		log.Fatalf("auto migrate: %v", err)
	}

	log.Println("AutoMigrate complete")
}
