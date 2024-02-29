package main

import (
	"context"
	"github.com/AzamatAbdranbayev/small.product/config"
	app2 "github.com/AzamatAbdranbayev/small.product/internal/app"
	"github.com/AzamatAbdranbayev/small.product/internal/handler"
	"github.com/AzamatAbdranbayev/small.product/internal/repo"
	"github.com/AzamatAbdranbayev/small.product/internal/service"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	e := echo.New()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	repo, err := repo.NewRepo(cfg.Db, ctx)
	if err != nil {
		log.Fatalln(err)
	}
	service := service.NewService(repo, cfg)
	handlers := handler.NewHandlers(service, e)
	handlers.InitRoutes()

	app := app2.NewApp(ctx).WithHTTP(e, cfg.Http)
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}

}
