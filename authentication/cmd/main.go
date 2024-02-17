package main

import (
	"log"
	"microservicesGRPC/authentication/internal/controller"
	"microservicesGRPC/authentication/internal/handler"
	"microservicesGRPC/authentication/internal/repository"
)

func main() {
	port := ":3000"
	repo := repository.New()
	usecase := controller.New(repo)
	handler := handler.New(usecase)
	log.Printf("Server started on %s...", port)
	if err := handler.Run(port); err != nil {
		log.Println(err)
	}
}
