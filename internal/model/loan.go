package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Loan struct {
	ID         uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id,omitzero"`
	BookID     uuid.UUID      `gorm:"type:uuid;not null"                             json:"book_id,omitzero"`
	Book       Book           `                                                      json:"book,omitzero"`
	BorrowerID uuid.UUID      `gorm:"type:uuid;not null"                             json:"borrower_id,omitzero"`
	Borrower   Borrower       `                                                      json:"borrower,omitzero"`
	DueDate    time.Time      `gorm:"not null"                                       json:"due_date,omitzero"`
	ReturnedAt *time.Time     `                                                      json:"returned_at,omitzero"`
	CreatedAt  time.Time      `gorm:"autoCreateTime;not null"                        json:"created_at,omitzero"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime;not null"                        json:"updated_at,omitzero"`
	DeletedAt  gorm.DeletedAt `gorm:"index"                                          json:"deleted_at,omitzero"`
}
