package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ISBN     string `gorm:"uniqueIndex;size:13;not null" json:"isbn,omitempty"`
	Title    string `gorm:"size:255;not null"            json:"title,omitempty"`
	AuthorID uint   `gorm:"not null"                     json:"author_id,omitempty"`
	Author   Author `                                    json:"author,omitempty"`
	Copies   uint   `gorm:"default:1;not null"           json:"copies,omitempty"`
}
