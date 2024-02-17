package main

import (
	"log"
	"microservicesGRPC/library/internal/controller"
	"microservicesGRPC/library/internal/handler"
	"microservicesGRPC/library/internal/repository"
)

func main() {
	repo := repository.New()
	usecase := controller.NewUsecase(repo)
	handler := handler.NewHandler(usecase)
	log.Println("Server is starting...")
	if err := handler.Run(":3500"); err != nil {
		log.Println(err)
	}
}
