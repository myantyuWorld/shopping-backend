package model

import (
	"errors"
)

type ShoppingCategory string

const (
	Food      ShoppingCategory = "food"
	Necessity ShoppingCategory = "necessity"
)

func NewShoppingCategory(category string) ShoppingCategory {
	return ShoppingCategory(category)
}

type ShoppingName string

func NewShoppingName(name string) (*ShoppingName, error) {
	if len(name) > 30 {
		return nil, errors.New("30文字以上は指定できません")
	}
	shoppingName := ShoppingName(name)

	return &shoppingName, nil
}

type ShoppingItem struct {
	ID       uint
	OwnerID  uint
	Category ShoppingCategory
	Name     ShoppingName
	Picked   bool
}

func NewShoppingItem(ownerID uint, category string, name string) (*ShoppingItem, error) {
	shoppingName, err := NewShoppingName(name)
	if err != nil {
		return nil, err
	}
	return &ShoppingItem{
		OwnerID:  ownerID,
		Category: NewShoppingCategory(category),
		Name:     *shoppingName,
		Picked:   false,
	}, nil
}

func NewShoppingItemDb(ID uint, ownerID uint, category string, name string, picked bool) *ShoppingItem {
	shoppingName, _ := NewShoppingName(name)
	return &ShoppingItem{
		ID:       ID,
		OwnerID:  ownerID,
		Category: NewShoppingCategory(category),
		Name:     *shoppingName,
		Picked:   picked,
	}

}
