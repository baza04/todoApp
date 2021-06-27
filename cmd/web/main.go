package main

import (
	"log"

	todoapp "github.com/baza04/todoApp"
	"github.com/baza04/todoApp/pkg/handler"
	"github.com/baza04/todoApp/pkg/repository"
	"github.com/baza04/todoApp/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializating configs: %s", err.Error())
	}

	repo := repository.NewRepository()
	services := service.NewService(repo)
	handler := handler.NewHandler(services)

	srv := new(todoapp.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
