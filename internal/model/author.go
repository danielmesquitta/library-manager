package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Author struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id,omitzero"`
	Name      string         `gorm:"size:255;not null"                              json:"name,omitzero"`
	Books     []Book         `                                                      json:"books,omitzero"`
	CreatedAt time.Time      `gorm:"autoCreateTime;not null"                        json:"created_at,omitzero"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;not null"                        json:"updated_at,omitzero"`
	DeletedAt gorm.DeletedAt `gorm:"index"                                          json:"deleted_at,omitzero"`
}
