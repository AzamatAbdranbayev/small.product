package models

import (
	"errors"
	"github.com/AzamatAbdranbayev/small.product/internal/constant"
	"github.com/AzamatAbdranbayev/small.product/pkg/helpers"
	"time"
)

type ProductUpdatePriceReq struct {
	Price uint64 `json:"price"`
}
type Product struct {
	Id        string    `json:"id" swaggerignore:"true"`
	Name      string    `json:"name"`
	Price     uint64    `json:"price"`
	CreatedAt time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" swaggerignore:"true"`
}

func (p *Product) CheckId() error {
	return helpers.CheckValidUuid(p.Id)
}
func (p *Product) CheckValidName() error {
	ln := len([]rune(p.Name))
	if ln > constant.MaxProductChar || ln <= constant.MinProductChar {
		return errors.New("not enough characters")
	}
	return nil
}

func (p *Product) CheckMaxPrice() error {
	if p.Price > constant.MaxProductPrice {
		return errors.New("price above maximum")
	}
	return nil
}
