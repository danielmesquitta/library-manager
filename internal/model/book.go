package model

import "time"

type Book struct {
    ID        uint      `gorm:"primaryKey"`
    ISBN      string    `gorm:"uniqueIndex;size:13;not null"`
    Title     string    `gorm:"size:255;not null"`
    AuthorID  uint
    Author    Author
    Copies    uint      `gorm:"default:1"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
