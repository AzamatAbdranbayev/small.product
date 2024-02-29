package service

import (
	"github.com/AzamatAbdranbayev/small.product/internal/models"
	"github.com/AzamatAbdranbayev/small.product/internal/repo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type IProduct interface {
	CreateNewProduct(product *models.Product) error
	GetProductById(id string) (models.Product, error)
	UpdateProductPrice(price uint64, productId string) error
}

type product struct {
	logInfo *zerolog.Event
	logErr  *zerolog.Event
	repo    repo.IProduct
}

func NewProductService(repo repo.IProduct) IProduct {
	return &product{
		logInfo: log.Info().Str("component", "product"),
		logErr:  log.Error().Stack().Str("component", "product"),
		repo:    repo,
	}
}

func (p *product) CreateNewProduct(product *models.Product) error {
	return p.repo.CreateNewProduct(product)
}
func (p *product) GetProductById(id string) (models.Product, error) {
	return p.repo.GetProductById(id)
}

func (p *product) UpdateProductPrice(price uint64, productId string) error {
	return p.repo.UpdatePriceByProductId(price, productId)
}
