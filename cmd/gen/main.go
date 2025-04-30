package main

import (
	"github.com/danielmesquitta/library-manager/internal/database"
	"github.com/danielmesquitta/library-manager/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/gen"
)

type TopBorrowerQuerier interface {
	// ListTopBorrowersByAuthor returns the top borrowers for each author.
	// The limit parameter specifies the maximum number of borrowers to return for each author.
	//
	// SELECT sub.author_id,
	// 	a.name AS author_name,
	// 	sub.borrower_id,
	// 	br.name AS borrower_name,
	// 	sub.loan_count
	// FROM (
	// 		SELECT bo.author_id,
	// 			l.borrower_id,
	// 			COUNT(*) AS loan_count,
	// 			ROW_NUMBER() OVER (
	// 				PARTITION BY bo.author_id
	// 				ORDER BY COUNT(*) DESC
	// 			) AS rn
	// 		FROM loan l
	// 			JOIN book bo ON l.book_id = bo.id
	// 		GROUP BY bo.author_id,
	// 			l.borrower_id
	// 	) AS sub
	// 	JOIN author a ON a.id = sub.author_id
	// 	JOIN borrower br ON br.id = sub.borrower_id
	// WHERE sub.rn <= @limit;
	ListTopBorrowersByAuthor(limit int) ([]gen.T, error)
}

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
		model.TopBorrower{},
	)

	// Generate custom query methods
	g.ApplyInterface(func(TopBorrowerQuerier) {}, model.TopBorrower{})

	g.Execute()
}
