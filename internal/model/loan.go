package model

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	BookID     uint       `gorm:"not null" json:"book_id,omitempty"`
	Book       Book       `                json:"book,omitempty"`
	BorrowerID uint       `gorm:"not null" json:"borrower_id,omitempty"`
	Borrower   Borrower   `                json:"borrower,omitempty"`
	DueDate    time.Time  `gorm:"not null" json:"due_date,omitempty"`
	ReturnedAt *time.Time `                json:"returned_at,omitempty"`
}
