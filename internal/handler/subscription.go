package handler

import (
	"errors"
	"github.com/AzamatAbdranbayev/small.product/internal/models"
	"github.com/AzamatAbdranbayev/small.product/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HandlerSubscription struct {
	service service.ISubScription
}

func NewHandlerSubscription(service service.ISubScription) *HandlerSubscription {
	return &HandlerSubscription{service}
}

// CreateNewSub godoc
//
//	@Tags		Subscription
//	@Summary	создать новую подписку
//	@Accept		json
//	@Produce	json
//
//	@Param		Body body  models.SubScription	true	"Тело"
//	@Success	200		{string} ok
//	@Router		/subscription/new [post]
func (h *HandlerSubscription) CreateNewSub(c echo.Context) error {
	var sub models.SubScription
	if err := c.Bind(&sub); err != nil {
		return err
	}

	if err := h.service.CreateNewSub(sub); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, "ok")
}

// GetSubsByUserId godoc
//
//	@Tags		Subscription
//	@Summary	получить лист подписок по юзеру
//	@Accept		json
//	@Produce	json
//
//	@Param		user_id path string true "ID it's user"
//	@Success	200		{array} models.SubScription
//	@Router		/subscription/user/{user_id} [get]
func (h *HandlerSubscription) GetSubsByUserId(c echo.Context) error {
	userId := c.Param("user_id")
	if userId == "" {
		return errors.New("empty userId")
	}

	subs, err := h.service.GetSubByUserId(userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, subs)
}
