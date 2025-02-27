package main

import (
	"log"

	"github.com/akproger/todo-app"
	"github.com/akproger/todo-app/pkg/handler"
	"github.com/akproger/todo-app/pkg/repository"
	"github.com/akproger/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
