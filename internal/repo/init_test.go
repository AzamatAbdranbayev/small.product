package repo

import (
	"context"
	"fmt"
	"github.com/AzamatAbdranbayev/small.product/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"testing"
	"time"
)

func initDb() (db *pgxpool.Pool, err error) {
	ctx := context.Background()
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "test_user", "test_pass", "localhost", "5432", "test_db")
	poolConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		return nil, err
	}
	db, err = pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}
	return
}
func initRepo() (*Repo, error) {
	db, err := initDb()
	if err != nil {
		return nil, err
	}

	return &Repo{
		User:         NewUserRepo(db),
		Subscription: NewSubscriptionRepo(db),
		Product:      NewProductRepo(db),
	}, nil
}
func TestDbQueries(t *testing.T) {
	repo, err := initRepo()
	if err != nil {
		t.Errorf("init database " + err.Error())
	}

	testProducts := []struct {
		name    string
		args    *models.Product
		wantErr bool
	}{
		{
			name: "create valid product name",
			args: &models.Product{
				Name:      "2213",
				Price:     155,
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			wantErr: false,
		},
		{
			name: "create invalid product name",
			args: &models.Product{
				Name:      "",
				Price:     12255,
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			wantErr: true,
		},
	}
	for _, tt := range testProducts {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.Product.CreateNewProduct(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("create valid product() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	//	=============== User Repo
	testUsers := []struct {
		name    string
		args    *models.User
		wantErr bool
	}{
		{
			name: "create user 1 ",
			args: &models.User{
				Email:       "22173",
				PhoneNumber: "4",
			},
			wantErr: false,
		},
		{
			name: "create user 2",
			args: &models.User{
				Email:       "22513",
				PhoneNumber: "5",
			},
			wantErr: false,
		},
	}
	for _, tt := range testUsers {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.User.CreateUser(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("create valid user() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
