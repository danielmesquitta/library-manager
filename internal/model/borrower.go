package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Borrower struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id,omitzero"`
	Name      string         `gorm:"size:255;not null"                              json:"name,omitzero"`
	Email     string         `gorm:"size:255;uniqueIndex"                           json:"email,omitzero"`
	Loans     []Loan         `                                                      json:"loans,omitzero"`
	CreatedAt time.Time      `gorm:"autoCreateTime;not null"                        json:"created_at,omitzero"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;not null"                        json:"updated_at,omitzero"`
	DeletedAt gorm.DeletedAt `gorm:"index"                                          json:"deleted_at,omitzero"`
}
