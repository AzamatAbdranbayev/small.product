package service

import (
	"github.com/AzamatAbdranbayev/small.product/config"
	"github.com/AzamatAbdranbayev/small.product/internal/repo"
)

type Service struct {
	UserService    IUser
	ProductService IProduct
	Subscription   ISubScription
	Smtp           ISmtp
	Facade         IFacade
}

func NewService(repo *repo.Repo, cfg *config.Config) *Service {

	s := &Service{}

	smtp := NewSmtpService(cfg.Smtp)
	s.Smtp = smtp

	userService := NewUserService(repo.User)
	s.UserService = userService

	productService := NewProductService(repo.Product)
	s.ProductService = productService

	subscriptionService := NewSubScriptionService(repo.Subscription, userService, productService)
	s.Subscription = subscriptionService

	facade := NewFacade(subscriptionService, productService, userService, smtp)
	s.Facade = facade
	return s
}
