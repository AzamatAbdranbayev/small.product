package handler

import (
	"errors"
	"github.com/AzamatAbdranbayev/small.product/internal/models"
	"github.com/AzamatAbdranbayev/small.product/internal/service"
	"github.com/AzamatAbdranbayev/small.product/pkg/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HandlerUser struct {
	service service.IUser
}

func NewHandlerUser(service service.IUser) *HandlerUser {
	return &HandlerUser{service}
}

// GetUserById godoc
//
//	@Tags		User
//	@Summary	Получить пользователя по айди
//	@Accept		json
//	@Produce	json
//
//	@Param		id path string true "user's id"
//	@Success	200		{object} models.User
//	@Router		/user/id/{id} [get]
func (h *HandlerUser) GetUserById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return errors.New("undefined id")
	}
	if err := helpers.CheckValidUuid(id); err != nil {
		return err
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// GetUserByEmail godoc
//
//	@Tags		User
//	@Summary	Получить пользователя по почте
//	@Accept		json
//	@Produce	json
//
//	@Param		id path string true "user's email"
//	@Success	200		{object} models.User
//	@Router		/user/email/{email} [get]
func (h *HandlerUser) GetUserByEmail(c echo.Context) error {
	email := c.Param("email")
	if email == "" {
		return errors.New("undefined id")
	}

	user, err := h.service.GetUserByEmail(email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// CreateUser godoc
//
//	@Tags		User
//	@Summary	создать нового пользователя
//	@Accept		json
//	@Produce	json
//
//	@Param		Body body  models.User	true	"Тело"
//	@Success	200		{object} models.User
//	@Router		/user/new [post]
func (h *HandlerUser) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := h.checkValid(&user); err != nil {
		return err
	}

	if err := h.service.CreateUser(&user); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *HandlerUser) checkValid(user *models.User) error {
	if err := user.CheckEmail(); err != nil {
		return err
	}
	if err := user.CheckPhoneNumber(); err != nil {
		return err
	}
	return nil
}
