package handler

import (
	"errors"
	"github.com/AzamatAbdranbayev/small.product/internal/models"
	"github.com/AzamatAbdranbayev/small.product/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HandlerProduct struct {
	service       service.IProduct
	serviceFacade service.IFacade
}

func NewHandlerProduct(service service.IProduct, serviceFacade service.IFacade) *HandlerProduct {
	return &HandlerProduct{service: service, serviceFacade: serviceFacade}
}

// CreateNewProduct godoc
//
//	@Tags		Product
//	@Summary	Создать новый продукт
//	@Accept		json
//	@Produce	json
//
//	@Param		Body body  models.Product	true	"Тело"
//	@Success	200		{object} models.Product
//	@Router		/product/new [post]
func (h *HandlerProduct) CreateNewProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return err
	}
	if err := h.checkValid(&product); err != nil {
		return err
	}
	if err := h.service.CreateNewProduct(&product); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, product)
}

// GetProductById godoc
//
//	@Tags		Product
//	@Summary	Получить новый продукт по айди
//	@Accept		json
//	@Produce	json
//
//	@Param		id path string true "ID it's product"
//	@Success	200		{object} models.Product
//	@Router		/product/{id} [get]
func (h *HandlerProduct) GetProductById(c echo.Context) error {
	productId := c.Param("id")
	if productId == "" {
		return errors.New("undefined product id")
	}
	product, err := h.service.GetProductById(productId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, product)
}

// UpdatePriceById godoc
//
//	@Tags		Product
//	@Summary	Изменить цену продукта
//	@Accept		json
//	@Produce	json
//
//	@Param		id path string true "ID it's product"
//	@Success	200		{string} ok
//	@Router		/product/{id} [patch]
func (h *HandlerProduct) UpdatePriceById(c echo.Context) error {
	productId := c.Param("id")
	if productId == "" {
		return errors.New("undefined product id")
	}

	var product models.ProductUpdatePriceReq
	if err := c.Bind(&product); err != nil {
		return err
	}

	if err := h.serviceFacade.UpdatePriceProduct(product.Price, productId); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}

func (h *HandlerProduct) checkValid(product *models.Product) error {
	if err := product.CheckValidName(); err != nil {
		return err
	}
	if err := product.CheckMaxPrice(); err != nil {
		return err
	}
	return nil
}
