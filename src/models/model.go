package models

import "time"

type Coin struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiryDate time.Time `json:"expiry_date"`
}

type ExpiredCoinLogs struct {
	Name      string
	ExpiredAt time.Time
}
