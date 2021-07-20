package main

import (
	"github.com/pArtour/todo"
	"github.com/pArtour/todo/package/handler"
	"github.com/pArtour/todo/package/repository"
	"github.com/pArtour/todo/package/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	server := new(todo.Server)
	if err := server.Run("8000", handlers.InitRouts()); err != nil {
		log.Fatalf("Error occred while running server %s", err.Error())
	}

}
