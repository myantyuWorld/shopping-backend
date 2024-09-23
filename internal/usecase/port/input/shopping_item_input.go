//go:generate mockgen -source=$GOFILE -destination=mock/mock_$GOFILE -package=mock
package input

import (
	"context"

	"github.com/LeoTwins/go-clean-architecture/internal/usecase/dto"
)

type IShoppingItemUsecase interface {
	Find(ctx context.Context, ownerID uint) ([]*dto.ShoppingItemOutput, error)
	Register(ctx context.Context, ownerID uint, name string, category string) error
	Remove(ctx context.Context, itemID uint) error
}
