package main

import (
	"log"
	todo "todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Serever)
	if err := srv.Run("8000", handlers.InnitRouter()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
