//go:generate mockgen -source=$GOFILE -destination=mock/mock_$GOFILE -package=mock
package output

import (
	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/dto"
)

type IShoppingItemPresenter interface {
	Output(domainItem model.ShoppingItem) dto.ShoppingItemOutput
}
