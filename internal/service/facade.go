package service

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type IFacade interface {
	UpdatePriceProduct(price uint64, productId string) error
}
type facade struct {
	logInfo      *zerolog.Event
	logErr       *zerolog.Event
	subscription ISubScription
	product      IProduct
	user         IUser
	smtp         ISmtp
}

func NewFacade(subscription ISubScription, product IProduct, user IUser, smtp ISmtp) IFacade {
	return &facade{
		logInfo:      log.Info().Str("component", "facade"),
		logErr:       log.Error().Stack().Str("component", "facade"),
		subscription: subscription,
		product:      product,
		user:         user,
		smtp:         smtp,
	}
}

func (f *facade) UpdatePriceProduct(price uint64, productId string) error {
	if err := f.product.UpdateProductPrice(price, productId); err != nil {
		return err
	}

	//TODO: ошибка в логике отправки и поиска подписок никак не должна возвращать ошибку, так как уже цена записалась
	// по хорошему тут лучше асинхронить, например RabbitMq использовать
	sbs, err := f.subscription.GetSubByProductId(productId)
	if err != nil {
		f.logErr.Err(err)
	}

	var ids []string
	for _, s := range sbs {
		ids = append(ids, s.UserId)
	}

	users, err := f.user.GetUsersByIds(ids)
	if err != nil {
		f.logErr.Err(err)
	}

	var emails []string
	for _, u := range users {
		emails = append(emails, u.Email)
	}
	f.smtp.SendEmailSubscribersUpdatedPrice(emails, productId)
	return nil
}
