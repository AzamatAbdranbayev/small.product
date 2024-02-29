package repo

import (
	"context"
	"errors"
	"github.com/AzamatAbdranbayev/small.product/internal/errapp"
	"github.com/AzamatAbdranbayev/small.product/internal/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type IProduct interface {
	CreateNewProduct(product *models.Product) error
	GetProductById(id string) (models.Product, error)
	UpdatePriceByProductId(price uint64, id string) error
}
type product struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) IProduct {
	return &product{db}
}

func (p *product) CreateNewProduct(product *models.Product) error {
	now := time.Now().UTC()
	err := p.db.QueryRow(context.Background(), `insert into products
			(name, price, created_at,updated_at)
			VALUES ($1, $2, $3,$4) returning (id)`,
		product.Name, product.Price, now, now).Scan(&product.Id)
	return err
}

func (p *product) GetProductById(id string) (models.Product, error) {
	var product models.Product

	err := p.db.QueryRow(context.Background(), `
	SELECT id, name, price, created_at,updated_at
	FROM products
	WHERE id = $1`, id).
		Scan(&product.Id, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return product, errapp.ErrorNotFoundProduct
		}
		return product, err
	}
	return product, nil
}

func (p *product) UpdatePriceByProductId(price uint64, id string) error {
	_, err := p.db.Exec(context.Background(),
		`update products SET price = $1 , updated_at= $2 where id =$3`,
		price, time.Now().UTC(), id)
	return err
}
