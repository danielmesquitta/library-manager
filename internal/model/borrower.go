package model

import "time"

type Borrower struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:255;not null"`
    Email     string    `gorm:"size:255;uniqueIndex"`
    Loans     []Loan
    CreatedAt time.Time
    UpdatedAt time.Time
}
