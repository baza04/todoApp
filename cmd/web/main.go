package main

import (
	"log"

	todoapp "github.com/baza04/todoApp"
	"github.com/baza04/todoApp/pkg/handler"
)

func main() {
	srv := new(todoapp.Server)
	handler := new(handler.Handler)

	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
