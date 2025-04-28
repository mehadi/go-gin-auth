package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title   string `gorm:"size:255" json:"title"`
	Content string `json:"content"`
}
