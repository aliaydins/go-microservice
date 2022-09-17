package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        int       `gorm:"primary_key;autoIncrement:true"`
	Email     string    `gorm:"size:255;not null;"`
	FirstName string    `gorm:"size:255;not null;"`
	LastName  string    `gorm:"size:255;not null;"`
	Password  string    `gorm:"size:255;not null;"`
	CreatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
