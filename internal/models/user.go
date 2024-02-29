package models

import (
	"errors"
	"github.com/AzamatAbdranbayev/small.product/internal/constant"
	"github.com/AzamatAbdranbayev/small.product/pkg/helpers"
	"net/mail"
	"time"
)

type User struct {
	Id          string    `json:"id" swaggerignore:"true"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at" swaggerignore:"true"`
}

func (u *User) CheckId() error {
	return helpers.CheckValidUuid(u.Id)
}
func (u *User) CheckEmail() error {
	_, err := mail.ParseAddress(u.Email)
	return err
}

func (u *User) CheckPhoneNumber() error {
	if len([]rune(u.PhoneNumber)) > constant.MinPhoneNumberChar {
		return errors.New("phone muber above maximum")
	}
	return nil
}
