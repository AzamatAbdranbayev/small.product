package service

import (
	"errors"
	"github.com/AzamatAbdranbayev/small.product/internal/models"
	"github.com/AzamatAbdranbayev/small.product/internal/repo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type IUser interface {
	GetUserById(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	CreateUser(user *models.User) error
	GetUsersByIds(ids []string) ([]models.User, error)
}

type user struct {
	logInfo *zerolog.Event
	logErr  *zerolog.Event
	repo    repo.IUser
}

func NewUserService(repo repo.IUser) IUser {
	return &user{
		logInfo: log.Info().Str("component", "user"),
		logErr:  log.Error().Stack().Str("component", "user"),
		repo:    repo,
	}
}

func (u *user) GetUserById(id string) (models.User, error) {
	return u.repo.GetUserById(id)
}
func (u *user) GetUserByEmail(email string) (models.User, error) {
	return u.repo.GetUserByEmail(email)
}

func (u *user) CreateUser(user *models.User) error {
	return u.repo.CreateUser(user)
}

func (u *user) GetUsersByIds(ids []string) ([]models.User, error) {
	if len(ids) == 0 {
		return nil, errors.New("empty ids")
	}
	return u.repo.GetUsersByIds(ids)
}
