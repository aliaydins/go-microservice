package entity

import (
	"gorm.io/gorm"
	"time"
)

type History struct {
	gorm.Model
	UserId    int    `gorm:"column:user_id"`
	Type      string `gorm:"size:255;not null;"`
	USD       int
	BTC       int
	Amount    int
	CreatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
