package repository

import (
	"context"
	"errors"

	domain "github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/LeoTwins/go-clean-architecture/internal/domain/repository"
	dbModel "github.com/LeoTwins/go-clean-architecture/internal/infrastructure/repository/model"
	"gorm.io/gorm"
)

type shoppingItemRepository struct {
	db *gorm.DB
}

// Find implements repository.IShoppingItemRepository.
func (r *shoppingItemRepository) Find(ctx context.Context) ([]*domain.ShoppingItem, error) {
	var dbItems []*dbModel.ShoppingItem
	result := r.db.Find(&dbItems)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	var shoppingItems []*domain.ShoppingItem
	for _, v := range dbItems {
		shoppingItems = append(shoppingItems, domain.NewShoppingItemDb(
			v.ID,
			v.OwnerID,
			v.Category,
			v.Name,
			v.Picked,
		))
	}

	return shoppingItems, nil
}

// LogicalDelete implements repository.IShoppingItemRepository.
func (r *shoppingItemRepository) LogicalDelete(ctx context.Context, item *domain.ShoppingItem) error {
	dbItem := dbModel.ShoppingItem{
		ID:     item.ID,
		Picked: item.Picked,
	}
	result := r.db.Save(&dbItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update implements repository.IShoppingItemRepository.
func (r *shoppingItemRepository) Update(ctx context.Context, item *domain.ShoppingItem) error {
	dbItem := dbModel.ShoppingItem{
		ID:       item.ID,
		OwnerID:  item.OwnerID,
		Category: string(item.Category),
		Name:     string(item.Name),
		Picked:   item.Picked,
	}

	result := r.db.Save(&dbItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Save implements repository.IShoppingItemRepository.
func (r *shoppingItemRepository) Save(ctx context.Context, item *domain.ShoppingItem) error {
	dbItem := dbModel.ShoppingItem{
		OwnerID:  item.OwnerID,
		Category: string(item.Category),
		Name:     string(item.Name),
		Picked:   item.Picked,
	}

	result := r.db.Create(&dbItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewShoppingItemRepository(db *gorm.DB) repository.IShoppingItemRepository {
	return &shoppingItemRepository{
		db: db,
	}
}
