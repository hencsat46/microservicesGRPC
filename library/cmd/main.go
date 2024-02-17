package main

import (
	"log"
	"microservicesGRPC/library/internal/controller"
	"microservicesGRPC/library/internal/handler"
	"microservicesGRPC/library/internal/repository"
)

func main() {
	port := ":3500"
	repo := repository.New()
	usecase := controller.NewUsecase(repo)
	handler := handler.NewHandler(usecase)
	log.Printf("Server started on %s...", port)
	if err := handler.Run(port); err != nil {
		log.Println(err)
	}
}
