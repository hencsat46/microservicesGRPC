package main

import (
	"log"
	"microservicesGRPC/authentication/internal/controller"
	"microservicesGRPC/authentication/internal/handler"
	"microservicesGRPC/authentication/internal/repository"
)

func main() {
	repo := repository.New()
	usecase := controller.New(repo)
	handler := handler.New(usecase)
	log.Println("Server is starting...")
	if err := handler.Run(":3000"); err != nil {
		log.Println(err)
	}
}
