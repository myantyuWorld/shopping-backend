package presenter

import (
	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/dto"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/port/output"
)

type shoppingItemPresenter struct{}

// Output implements output.IShoppingItemPresenter.
func (s *shoppingItemPresenter) Output(domainItem model.ShoppingItem) dto.ShoppingItemOutput {
	return dto.ShoppingItemOutput{
		ID:       domainItem.ID,
		OwnerID:  domainItem.OwnerID,
		Name:     string(domainItem.Name),
		Category: string(domainItem.Category),
		Picked:   domainItem.Picked,
	}
}

func NewShoppingItemPresenter() output.IShoppingItemPresenter {
	return &shoppingItemPresenter{}
}
