package main

import (
	"log"
	"microservicesGRPC/authentication/internal/handler"
)

func main() {
	handler := handler.New(5)
	if err := handler.Run(":3000"); err != nil {
		log.Println(err)
	}
}
