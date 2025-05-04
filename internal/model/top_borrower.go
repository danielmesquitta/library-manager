package model

import "github.com/google/uuid"

type TopBorrower struct {
	AuthorID     uuid.UUID `json:"author_id"`
	AuthorName   string    `json:"author_name"`
	BorrowerID   uuid.UUID `json:"borrower_id"`
	BorrowerName string    `json:"borrower_name"`
	LoanCount    int       `json:"loan_count"`
}
