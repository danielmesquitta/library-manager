package model

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  string `gorm:"size:255;not null" json:"name,omitempty"`
	Books []Book `                         json:"books,omitempty"`
}
