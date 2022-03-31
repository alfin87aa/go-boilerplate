package models

import (
	"boilerplate/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primaryKey;"`
	Fullname  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);unique;not null" sql:"index"`
	Password  string    `gorm:"type:varchar(255)"`
	Active    bool      `gorm:"type:bool;default:true"`
	CreatedAt time.Time `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}

func (entity *User) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.Password = utils.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	entity.UpdatedAt = time.Now().Local()
	return nil
}

func (entity *User) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
