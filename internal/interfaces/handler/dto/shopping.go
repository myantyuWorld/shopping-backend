package dto

type ShoppingItemResponse struct {
	ID       uint   `json:"id"`
	OwnerID  uint   `json:"owner_id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Picked   bool   `json:"picked"`
}

type ShoppingItemRequest struct {
	OwnerID  uint   `json:"owner_id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
