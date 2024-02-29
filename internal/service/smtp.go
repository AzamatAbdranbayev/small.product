package service

import (
	"errors"
	"github.com/AzamatAbdranbayev/small.product/config"
	"log"
	"net/smtp"
)

type ISmtp interface {
	SendEmailSubscribersUpdatedPrice(emails []string, productId string)
}
type smtpS struct {
	cfg  *config.SmtpConfig
	auth smtp.Auth
}
type loginAuth struct {
	username, password string
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}
func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}
func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func NewSmtpService(cfg *config.SmtpConfig) ISmtp {
	s := &smtpS{cfg: cfg}

	auth := LoginAuth(s.cfg.ProducerEmail, s.cfg.Token)
	s.auth = auth

	return s
}

func (s *smtpS) SendEmailSubscribersUpdatedPrice(emails []string, productId string) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subJect := "Информация об изменении цены товаров"

	msg := []byte("Subject: " + subJect + "\r\n" +
		"From: " + s.cfg.ProducerEmail + "\r\n" +
		mime +
		"Цена изменилась у продукта с идентификатором : " + productId)

	for _, email := range emails {
		if err := s.sendEmail(email, msg); err != nil {
			log.Println(err.Error())
		}
	}
}
func (s *smtpS) sendEmail(email string, msg []byte) error {
	from := s.cfg.ProducerEmail
	to := []string{email}
	err := smtp.SendMail(s.cfg.SmtpHost, s.auth, from, to, msg)
	if err != nil {
		return err
	}
	return nil
}
