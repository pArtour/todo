package main

import (
	"github.com/pArtour/todo"
	"github.com/pArtour/todo/package/handler"
	"log"
)

func main() {
	server := new(todo.Server)
	handlers := new(handler.Handler)

	if err := server.Run("8000", handlers.InitRouts()); err != nil {
		log.Fatalf("Error occred while running server %s", err.Error())
	}

}
