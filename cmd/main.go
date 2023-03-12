package main

import (
	"log"
	todo "todo-app"
	"todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Serever)
	if err := srv.Run("8000", handlers.InnitRouter()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
