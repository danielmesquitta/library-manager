package model

import (
	"gorm.io/gorm"
)

type Borrower struct {
	gorm.Model
	Name  string `gorm:"size:255;not null"    json:"name,omitempty"`
	Email string `gorm:"size:255;uniqueIndex" json:"email,omitempty"`
	Loans []Loan `                            json:"loans,omitempty"`
}
