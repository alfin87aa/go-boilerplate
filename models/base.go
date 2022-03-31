package models

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

//Model is sample of common table structure
type Model struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index"`
}
