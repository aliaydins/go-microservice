package entity

import (
	"gorm.io/gorm"
	"time"
)

type OrderType string

const (
	BUY  OrderType = "BUY"
	SELL OrderType = "SELL"
)

type Order struct {
	gorm.Model
	ID        int       `gorm:"primary_key;autoIncrement:true"`
	UserId    int       `gorm:"column:user_id"`
	Type      OrderType `gorm:"size:255;not null;"`
	Price     int
	Quantity  int
	CreatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
