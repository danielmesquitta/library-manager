package main

import (
	"github.com/danielmesquitta/library-manager/internal/database"
	"github.com/danielmesquitta/library-manager/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/gen"
)

func main() {
	_ = godotenv.Load()

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:       "internal/query",
		Mode:          gen.WithDefaultQuery,
		FieldNullable: true,
	})
	g.UseDB(db)

	// Generate basic type-safe DAO APIs for all models
	g.ApplyBasic(
		model.Author{},
		model.Book{},
		model.Borrower{},
		model.Loan{},
	)

	g.Execute()
}
