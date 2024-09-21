package model

import "time"

type ShoppingItem struct {
	ID        uint      `gorm:"primaryKey autoIncrement"`
	OwnerID   uint      `gorm:"not null"`
	Category  string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Picked    bool      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Tabler interface {
	TableName() string
}

func (s ShoppingItem) TableName() string {
	return "shopping_items"
}
