package model

import "time"

type ShoppingResult struct {
	ID          uint
	OwnerID     uint
	Category    ShoppingCategory
	Date        time.Time
	TotalAmount Money
	Picture     string
}

func NewShoppingResult(ownerID uint, category ShoppingCategory, date time.Time, totalAmount Money, picture string) (*ShoppingResult, error) {
	return &ShoppingResult{
		OwnerID:     ownerID,
		Category:    category,
		Date:        date,
		TotalAmount: totalAmount,
		Picture:     picture,
	}, nil
}
