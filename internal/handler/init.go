package handler

import (
	_ "github.com/AzamatAbdranbayev/small.product/docs"
	"github.com/AzamatAbdranbayev/small.product/internal/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

type Handlers struct {
	HandlerUser         *HandlerUser
	HandlerProduct      *HandlerProduct
	HandlerSubscription *HandlerSubscription
	router              *echo.Echo
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func NewHandlers(service *service.Service, e *echo.Echo) *Handlers {
	userH := NewHandlerUser(service.UserService)

	productH := NewHandlerProduct(service.ProductService, service.Facade)

	subH := NewHandlerSubscription(service.Subscription)
	return &Handlers{
		HandlerUser:         userH,
		HandlerProduct:      productH,
		HandlerSubscription: subH,
		router:              e,
	}
}

func (h *Handlers) InitRoutes() {
	h.router.GET("/swagger/*", echoSwagger.WrapHandler)
	h.router.GET("/health/check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	users := h.router.Group("/user")
	users.GET("/id/:id", h.HandlerUser.GetUserById)
	users.GET("/email/:email", h.HandlerUser.GetUserByEmail)
	users.POST("/new", h.HandlerUser.CreateUser)

	product := h.router.Group("/product")
	product.GET("/:id", h.HandlerProduct.GetProductById)
	product.POST("/new", h.HandlerProduct.CreateNewProduct)
	product.PATCH("/:id", h.HandlerProduct.UpdatePriceById)

	subscription := h.router.Group("/subscription")
	subscription.POST("/new", h.HandlerSubscription.CreateNewSub)
	subscription.GET("/user/:user_id", h.HandlerSubscription.GetSubsByUserId)
}
