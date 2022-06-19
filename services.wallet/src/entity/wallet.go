package entity

import (
	"gorm.io/gorm"
	"time"
)

type Currency string

const (
	BTC Currency = "BTC"
	ETH Currency = "ETH"
)

type Wallet struct {
	gorm.Model
	ID        int `gorm:"primary_key;autoIncrement:true"`
	UserId    int `gorm:"column:user_id"`
	Currency  Currency
	Quantity  int
	CreatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
