package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id,omitzero"`
	ISBN      string         `gorm:"uniqueIndex;size:13;not null"                   json:"isbn,omitzero"`
	Title     string         `gorm:"size:255;not null"                              json:"title,omitzero"`
	AuthorID  uuid.UUID      `gorm:"type:uuid;not null"                             json:"author_id,omitzero"`
	Author    Author         `                                                      json:"author,omitzero"`
	Copies    uint           `gorm:"default:1;not null"                             json:"copies,omitzero"`
	CreatedAt time.Time      `gorm:"autoCreateTime;not null"                        json:"created_at,omitzero"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime; not null"                       json:"updated_at,omitzero"`
	DeletedAt gorm.DeletedAt `gorm:"index"                                          json:"deleted_at,omitzero"`
}
