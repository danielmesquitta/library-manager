package model

import "time"

type Author struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:255;not null"`
    Books     []Book
    CreatedAt time.Time
    UpdatedAt time.Time
}
