package service

import (
	"github.com/AzamatAbdranbayev/small.product/internal/models"
	"github.com/AzamatAbdranbayev/small.product/internal/repo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ISubScription interface {
	CreateNewSub(sub models.SubScription) error
	GetSubByUserId(userId string) ([]models.SubScription, error)
	GetSubByProductId(productId string) ([]models.SubScription, error)
}

type subscription struct {
	logInfo        *zerolog.Event
	logErr         *zerolog.Event
	repo           repo.ISubscription
	userService    IUser
	productService IProduct
}

func NewSubScriptionService(repo repo.ISubscription, userService IUser, productService IProduct) ISubScription {
	return &subscription{
		logInfo:        log.Info().Str("component", "subscription"),
		logErr:         log.Error().Stack().Str("component", "subscription"),
		repo:           repo,
		userService:    userService,
		productService: productService,
	}
}

func (s *subscription) CreateNewSub(sub models.SubScription) error {
	if _, err := s.userService.GetUserById(sub.UserId); err != nil {
		return err
	}

	if _, err := s.productService.GetProductById(sub.ProductId); err != nil {
		return err
	}
	return s.repo.CreateSub(sub.UserId, sub.ProductId)
}

func (s *subscription) GetSubByUserId(userId string) ([]models.SubScription, error) {
	if _, err := s.userService.GetUserById(userId); err != nil {
		return nil, err
	}

	return s.repo.GetSubByUserId(userId)
}

func (s *subscription) GetSubByProductId(productId string) ([]models.SubScription, error) {
	if _, err := s.productService.GetProductById(productId); err != nil {
		return nil, err
	}

	return s.repo.GetSubByProductId(productId)
}
