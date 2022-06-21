package entity

import (
	"gorm.io/gorm"
	"time"
)

type OrderType string
type Currency string

const (
	BUY  OrderType = "BUY"
	SELL OrderType = "SELL"
)

const (
	BTC Currency = "BTC"
	ETH Currency = "ETH"
)

type Order struct {
	gorm.Model
	ID              int `gorm:"primary_key;autoIncrement:true"`
	UserId          int
	MatchingOrderId int
	Type            OrderType `gorm:"size:4;not null;"`
	Currency        Currency  `gorm:"size:3;not null;"`
	Quantity        int       `gorm:"size:32;not null;"`
	OrderPrice      int       `gorm:"size:32;not null;"`
	CreatedAt       time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
