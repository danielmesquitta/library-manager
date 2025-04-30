package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/danielmesquitta/library-manager/internal/database"
	"github.com/danielmesquitta/library-manager/internal/model"
	"github.com/danielmesquitta/library-manager/internal/query"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	ctx := context.Background()

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}

	// Initialize query over existing db connection
	query.SetDefault(db)

	jk := &model.Author{Name: "J. K. Rowling"}
	if err := query.Author.WithContext(ctx).Create(jk); err != nil {
		log.Fatalf("insert author: %v", err)
	}
	fmt.Println("new author id:", jk.ID)

	hp1 := &model.Book{
		ISBN:     "9780747532699",
		Title:    "Harry Potter and the Philosopher’s Stone",
		AuthorID: jk.ID,
		Copies:   3,
	}

	hp2 := &model.Book{
		ISBN:     "9780747532698",
		Title:    "Harry Potter and the Goblet of Fire",
		AuthorID: jk.ID,
		Copies:   2,
	}

	if err := query.Book.WithContext(ctx).Create(hp1, hp2); err != nil {
		log.Fatal(err)
	}

	books, _ := query.
		Book.
		WithContext(ctx).
		Preload(query.Book.Author). // eager-loads the Author relation
		Order(query.Book.Title).
		Find()

	for _, b := range books {
		fmt.Printf("%s — %s\n", b.Title, b.Author.Name)
	}

	dan := &model.Borrower{Name: "Daniel Mesquita", Email: "dan@example.com"}
	_ = query.Borrower.WithContext(ctx).Create(dan)

	due := time.Now().AddDate(0, 0, 14) // two-week loan

	err = query.Q.Transaction(func(tx *query.Query) error {
		// create loan record
		err := tx.Loan.WithContext(ctx).Create(&model.Loan{
			BookID:     hp1.ID,
			BorrowerID: dan.ID,
			DueDate:    due,
		})
		if err != nil {
			return err
		}

		// decrement available copies atomically
		_, err = tx.Book.WithContext(ctx).
			Where(tx.Book.ID.Eq(hp1.ID)).
			UpdateSimple(tx.Book.Copies.Sub(1))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatalf("transaction failed: %v", err)
	}

	now := time.Now()

	_ = query.Q.Transaction(func(tx *query.Query) error {
		_, err := tx.Loan.WithContext(ctx).
			Where(tx.Loan.BookID.Eq(hp1.ID), tx.Loan.BorrowerID.Eq(dan.ID), tx.Loan.ReturnedAt.IsNull()).
			Update(tx.Loan.ReturnedAt, now)
		if err != nil {
			return err
		}

		_, err = tx.Book.WithContext(ctx).
			Where(tx.Book.ID.Eq(hp1.ID)).
			UpdateSimple(tx.Book.Copies.Add(1))
		if err != nil {
			return err
		}

		return nil
	})

	overdues, _ := query.
		Loan.
		WithContext(ctx).
		Preload(query.Loan.Book, query.Loan.Borrower).
		Where(
			query.Loan.ReturnedAt.IsNull(),
			query.Loan.DueDate.Lt(time.Now()),
		).
		Find()

	for _, l := range overdues {
		fmt.Printf("⚠️  %s is overdue for %s (due %s)\n",
			l.Book.Title, l.Borrower.Name, l.DueDate.Format(time.DateOnly))
	}
}
