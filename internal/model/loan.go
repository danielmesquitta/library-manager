package model

import "time"

type Loan struct {
    ID         uint      `gorm:"primaryKey"`
    BookID     uint
    Book       Book
    BorrowerID uint
    Borrower   Borrower
    DueDate    time.Time
    ReturnedAt *time.Time
    CreatedAt  time.Time
    UpdatedAt  time.Time
}
