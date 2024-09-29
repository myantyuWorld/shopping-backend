package dto

type ShoppingItemOutput struct {
	ID       uint
	OwnerID  uint
	Name     string
	Category string
	Picked   bool
}
