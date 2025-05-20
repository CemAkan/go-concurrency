package db

import "gorm.io/gorm"

type ShortURL struct {
	gorm.Model
	OriginalURL string `gorm:"not null"`
	ShortCode   string `gorm:"uniqueIndex;not null"`
	ClickCount  int64  `gorm:"default:0"`
}
