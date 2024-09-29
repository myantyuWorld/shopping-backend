//go:generate mockgen -source=$GOFILE -destination=mock/mock_$GOFILE -package=mock
package input

import (
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/dto"
)

type ITransactionUsecase interface {
	FindByID(id uint) (*dto.TransactionOutput, error)
	FindByAccountID(accountID uint) ([]*dto.TransactionOutput, error)
}
