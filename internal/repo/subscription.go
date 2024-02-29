package repo

import (
	"context"
	"github.com/AzamatAbdranbayev/small.product/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ISubscription interface {
	CreateSub(userId, productId string) error
	GetSubByUserId(userId string) ([]models.SubScription, error)
	GetSubByProductId(productId string) ([]models.SubScription, error)
}

type subscription struct {
	db *pgxpool.Pool
}

func NewSubscriptionRepo(db *pgxpool.Pool) ISubscription {
	return &subscription{db}
}

func (s *subscription) CreateSub(userId, productId string) error {
	_, err := s.db.Exec(context.Background(), `INSERT INTO subscriptions (user_id, product_id) VALUES ($1,$2)`, userId, productId)
	return err
}

func (s *subscription) GetSubByUserId(userId string) ([]models.SubScription, error) {
	var sbs []models.SubScription

	rows, err := s.db.Query(context.Background(), `SELECT * FROM subscriptions WHERE user_id = $1`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var sb models.SubScription
		if err := rows.Scan(&sb.UserId, &sb.ProductId); err != nil {
			return nil, err
		}
		sbs = append(sbs, sb)
	}
	return sbs, err
}
func (s *subscription) GetSubByProductId(productId string) ([]models.SubScription, error) {
	var sbs []models.SubScription

	rows, err := s.db.Query(context.Background(), `SELECT * FROM subscriptions WHERE product_id = $1`, productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var sb models.SubScription
		if err := rows.Scan(&sb.UserId, &sb.ProductId); err != nil {
			return nil, err
		}
		sbs = append(sbs, sb)
	}
	return sbs, err
}
