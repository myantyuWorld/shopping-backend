package interacter

import (
	"context"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/LeoTwins/go-clean-architecture/internal/domain/repository"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/dto"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/port/input"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/port/output"
)

type shoppingItemUsecase struct {
	shoppingItemRepo repository.IShoppingItemRepository
	presenter        output.IShoppingItemPresenter
}

// Find implements input.IShoppingItemUsecase.
func (s *shoppingItemUsecase) Find(ctx context.Context, ownerID uint) ([]*dto.ShoppingItemOutput, error) {
	items, err := s.shoppingItemRepo.FindByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	outputs := []*dto.ShoppingItemOutput{}
	for _, v := range items {
		output := s.presenter.Output(*v)
		outputs = append(outputs, &output)
	}
	return outputs, nil
}

// Register implements input.IShoppingItemUsecase.
func (s *shoppingItemUsecase) Register(ctx context.Context, ownerID uint, name string, category string) error {
	item, err := model.NewShoppingItem(ownerID, category, name)
	if err != nil {
		return err
	}

	err = s.shoppingItemRepo.Save(ctx, item)
	if err != nil {
		return err
	}
	return nil
}

// Remove implements input.IShoppingItemUsecase.
func (s *shoppingItemUsecase) Remove(ctx context.Context, itemID uint) error {
	err := s.shoppingItemRepo.LogicalDelete(ctx, itemID)
	if err != nil {
		return err
	}
	return nil
}

func NewShoppingItemUsecase(shoppingItemRepo repository.IShoppingItemRepository, presenter output.IShoppingItemPresenter) input.IShoppingItemUsecase {
	return &shoppingItemUsecase{
		shoppingItemRepo: shoppingItemRepo,
		presenter:        presenter,
	}
}
