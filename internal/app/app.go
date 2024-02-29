package app

import (
	"context"
	"github.com/AzamatAbdranbayev/small.product/config"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	starters []func(ctx context.Context) error
	stoppers []func(ctx context.Context) error
	ctx      context.Context
	errs     chan error
}

func NewApp(ctx context.Context) *App {
	return &App{
		ctx:  ctx,
		errs: make(chan error),
	}
}

func (app *App) withStart(starter func(ctx context.Context) error) *App {
	app.starters = append(app.starters, starter)
	return app
}

func (app *App) withStop(stopper func(ctx context.Context) error) *App {
	app.stoppers = append(app.stoppers, stopper)
	return app
}

func (app *App) Run() error {
	ctx, cancel := context.WithCancel(app.ctx)
	defer cancel()

	app.ctx = ctx

	for _, startFunc := range app.starters {
		go func(startFunc func(ctx context.Context) error) {
			if err := startFunc(app.ctx); err != nil {
				app.errs <- err
			}
		}(startFunc)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-quit:
		log.Info().Msg("sigterm or sigint received. Stopping app...")
		return app.stop()
	case <-app.ctx.Done():
		log.Info().Msg("app context was cancelled. Stopping app...")
		return app.stop()
	case err := <-app.errs:
		log.Error().Msg("error occurred on app initialization")
		return err
	}
}
func (app *App) stop() error {
	for _, stopFunc := range app.stoppers {
		if err := stopFunc(app.ctx); err != nil {
			return err
		}
	}
	log.Info().Msg("service stopped")
	return nil
}
func (app *App) WithHTTP(e *echo.Echo, cfg *config.HttpConfig) *App {
	//e.Logger.set
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.HideBanner = true

	return app.
		withStart(func(ctx context.Context) error {
			log.Info().Msgf("starting server on %v", cfg.Address)
			err := e.Start(cfg.Address)
			if err != nil {
				return errors.Wrap(err, "start httpclient")
			}

			return nil
		}).
		withStop(func(ctx context.Context) error {

			log.Info().Msg("stopping httpclient server...")
			err := e.Shutdown(ctx)
			if err != nil {
				return errors.Wrap(err, "httpclient server shutdown")
			}
			log.Info().Msg("httpclient server successfully stopped")
			return nil
		})
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusBadRequest
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	//TODO: тут поле message запишем на дефолтную ошибку, системную ошибку в STDOUT и в какое-нибудь хранилище
	c.JSON(code, map[string]interface{}{
		"error": map[string]interface{}{
			"code":    code,
			"message": "unknown error",
		},
	})
}
