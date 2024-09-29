//go:generate mockgen -source=$GOFILE -destination=mock/mock_$GOFILE -package=mock
package repository

import (
	"context"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
)

type IShoppingItemRepository interface {
	FindByOwnerID(ctx context.Context, ownerID uint) ([]*model.ShoppingItem, error)
	Save(ctx context.Context, item *model.ShoppingItem) error
	LogicalDelete(ctx context.Context, itemID uint) error
}
