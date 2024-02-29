package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/AzamatAbdranbayev/small.product/internal/errapp"
	"github.com/AzamatAbdranbayev/small.product/internal/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
	"time"
)

type IUser interface {
	GetUserById(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	CreateUser(user *models.User) error
	GetUsersByIds(ids []string) ([]models.User, error)
}
type user struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) IUser {
	return &user{db}
}

func (u *user) GetUserById(id string) (models.User, error) {
	var user models.User

	err := u.db.QueryRow(context.Background(), `
	SELECT users.id, users.email,  users.phone_number, users.created_at
	FROM users
	WHERE users.id = $1`, id).
		Scan(&user.Id, &user.Email, &user.PhoneNumber, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, errapp.ErrorNotFoundUser
		}
		return user, err
	}
	return user, nil
}

func (u *user) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := u.db.QueryRow(context.Background(), `
	SELECT users.id, users.email,  users.phone_number, users.created_at
	FROM users
	WHERE users.email = $1`, email).
		Scan(&user.Id, &user.Email, &user.PhoneNumber, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, errapp.ErrorNotFoundUser
		}
		return user, err
	}
	return user, nil
}

func (u *user) GetUsersByIds(ids []string) ([]models.User, error) {
	var users []models.User

	query := fmt.Sprintf("SELECT id, email, phone_number FROM users WHERE id IN ('%s')", strings.Join(ids, "', '"))
	rows, err := u.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.Id, &u.Email, &u.PhoneNumber); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
func (u *user) CreateUser(user *models.User) error {
	err := u.db.QueryRow(context.Background(), `insert into users
			( email, phone_number, created_at)
			VALUES ($1, $2, $3) returning (id)`,
		user.Email, user.PhoneNumber, time.Now().UTC()).Scan(&user.Id)
	return err
}
