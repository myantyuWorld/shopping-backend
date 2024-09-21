package model

import "time"

type ShoppingResult struct {
	ID          uint      `gorm:"primaryKey autoIncrement"`
	OwnerID     uint      `gorm:"not null"`
	Category    string    `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
	TotalAmount uint      `gorm:"not null"`
	Picture     string    `gorm:"not null"`
	Picked      bool      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
