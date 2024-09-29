package router

import (
	"log"
	"net/http"

	"github.com/LeoTwins/go-clean-architecture/internal/infrastructure/repository"
	"github.com/LeoTwins/go-clean-architecture/internal/infrastructure/service"
	"github.com/LeoTwins/go-clean-architecture/internal/interfaces/handler"
	"github.com/LeoTwins/go-clean-architecture/internal/interfaces/middleware"
	"github.com/LeoTwins/go-clean-architecture/internal/interfaces/presenter"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/interacter"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewRouter(e *echo.Echo, db *gorm.DB) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Print(err)
	}

	e.Use(middleware.Logger(logger))

	accountRepo := repository.NewAccountRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	shoppingItemRepo := repository.NewShoppingItemRepository(db)
	transactionManager := service.NewTransactionManager(db)

	accountUsecase := interacter.NewAccountUsecase(accountRepo, transactionRepo, transactionManager)
	transactionPresenter := presenter.NewTransactionPresenter()
	transactionUsecase := interacter.NewTransactionUsecase(transactionRepo, transactionPresenter)
	shoppingItemPresenter := presenter.NewShoppingItemPresenter()
	shoppingItemUsecase := interacter.NewShoppingItemUsecase(shoppingItemRepo, shoppingItemPresenter)

	accountHandler := handler.NewAccountHandler(accountUsecase)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)
	shoppingItemHandler := handler.NewShoppingItemHandler(shoppingItemUsecase)

	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.POST("/open-account", accountHandler.OpenAccount)
	e.POST("/deposit", accountHandler.Deposit)
	e.POST("/withdraw", accountHandler.Withdraw)
	e.POST("/transfer", accountHandler.Transfer)

	e.GET("/:id", transactionHandler.FindByID)
	e.GET("/:account-id", transactionHandler.FindByAccountID)

	shoppingItemApi := e.Group("/shopping")
	shoppingItemApi.GET("/item/:owner_id", shoppingItemHandler.FindByOwnerID)
	shoppingItemApi.POST("/item", shoppingItemHandler.Register)
	shoppingItemApi.POST("/item/:item_id", shoppingItemHandler.Remove)
}
