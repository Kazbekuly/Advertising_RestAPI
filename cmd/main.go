package main

import (
	"Advertising"
	"Advertising/configs"
	"Advertising/pkg/handler"
	"Advertising/pkg/logging"
	"Advertising/pkg/repository"
	"Advertising/pkg/service"
	"fmt"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Info("start application")
	srv := new(Advertising.Server)
	cfg := configs.GetConfig()
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logger.Fatal("error while connecting to database")
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	_ = fmt.Sprintf("Starting server...\nhttp://localhost%v/\n", ":"+cfg.Listen.Port)
	if err := srv.Run(cfg, handlers.InitRoutes()); err != nil {
		logger.Info("Error while starting application")
	}
}
