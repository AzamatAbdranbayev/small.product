package repo

import (
	"context"
	"fmt"
	"github.com/AzamatAbdranbayev/small.product/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repo struct {
	User         IUser
	Product      IProduct
	Subscription ISubscription
}

func NewRepo(cfg *config.DbConfig, ctx context.Context) (*Repo, error) {
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DbName)
	poolConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return &Repo{
		User:         NewUserRepo(db),
		Subscription: NewSubscriptionRepo(db),
		Product:      NewProductRepo(db),
	}, nil
}
