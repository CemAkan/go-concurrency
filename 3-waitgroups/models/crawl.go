package models

import "gorm.io/gorm"

type CrawlPage struct {
	gorm.Model
	URL      string    `gorm:"uniqueIndex"`
	Elements []Element `gorm:"foreignKey:PageID"`
}

type Element struct {
	gorm.Model
	PageID      uint
	ElementType string
	Content     string
	Attribute   string
	Position    int
}
