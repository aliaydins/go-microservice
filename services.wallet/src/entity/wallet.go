package entity

import (
	"gorm.io/gorm"
	"time"
)

type Wallet struct {
	gorm.Model
	ID        int `gorm:"primary_key;autoIncrement:true"`
	UserId    int `gorm:"column:user_id"`
	USD       int
	BTC       int
	CreatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
