package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeoTwins/go-clean-architecture/internal/interfaces/handler/dto"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/port/input"
	"github.com/labstack/echo/v4"
)

type IShoppingItemHandler interface {
	FindByOwnerID(c echo.Context) error
	Register(c echo.Context) error
	Remove(c echo.Context) error
}

type shoppingItemHandler struct {
	u input.IShoppingItemUsecase
}

// FindByOwnerID implements IShoppingItemHandler.
func (s *shoppingItemHandler) FindByOwnerID(c echo.Context) error {
	ownerID, err := strconv.Atoi(c.Param("owner_id"))
	if err != nil {
		// TODO : Middlewareで処理させるでいい
		return c.JSON(http.StatusBadRequest, err)
	}

	items, err := s.u.Find(c.Request().Context(), uint(ownerID))
	if err != nil {
		// TODO : Middlewareで処理させるでいい
		return c.JSON(http.StatusBadRequest, err)
	}

	responses := []*dto.ShoppingItemResponse{}
	for _, v := range items {
		fmt.Print(v)
		res := &dto.ShoppingItemResponse{
			ID:       v.ID,
			OwnerID:  v.OwnerID,
			Name:     v.Name,
			Category: v.Category,
			Picked:   v.Picked,
		}
		responses = append(responses, res)
	}
	return c.JSON(http.StatusOK, responses)
}

// Register implements IShoppingItemHandler.
func (s *shoppingItemHandler) Register(c echo.Context) error {
	req := dto.ShoppingItemRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := s.u.Register(c.Request().Context(), req.OwnerID, req.Name, req.Category); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.NoContent(http.StatusCreated)
}

// Remove implements IShoppingItemHandler.
func (s *shoppingItemHandler) Remove(c echo.Context) error {
	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		// TODO : Middlewareで処理させるでいい
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := s.u.Remove(c.Request().Context(), uint(itemID)); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.NoContent(http.StatusOK)
}

func NewShoppingItemHandler(u input.IShoppingItemUsecase) IShoppingItemHandler {
	return &shoppingItemHandler{
		u: u,
	}
}
