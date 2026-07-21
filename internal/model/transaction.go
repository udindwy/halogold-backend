package model

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Type      string    `gorm:"type:varchar(10);not null" json:"type"`
	Amount    float64   `gorm:"type:numeric;not null" json:"amount"`
	Gram      float64   `gorm:"type:numeric;not null" json:"gram"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
