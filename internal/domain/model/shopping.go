package model

import "errors"

type ShoppingCategory string

const (
	Food      ShoppingCategory = "Food"
	Necessity ShoppingCategory = "Necessity"
)

type ShoppingName string

type ShoppingItem struct {
	OwnerID  uint
	Category ShoppingCategory
	Name     ShoppingName
	Picked   bool
}

func NewShoppingName(name string) (*ShoppingName, error) {
	if len(name) > 30 {
		return nil, errors.New("30文字以上は指定できません")
	}
	shoppingName := ShoppingName(name)

	return &shoppingName, nil
}

func NewShoppingItem(ownerID uint, category ShoppingCategory, name string) (*ShoppingItem, error) {
	shoppingName, err := NewShoppingName(name)
	if err != nil {
		return nil, err
	}
	return &ShoppingItem{
		OwnerID:  ownerID,
		Category: category,
		Name:     *shoppingName,
		Picked:   false,
	}, nil
}
