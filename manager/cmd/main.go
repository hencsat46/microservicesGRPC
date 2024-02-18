package main

import (
	"log"
	"microservicesGRPC/manager/internal/handler"
)

func main() {
	port := ":4000"
	handler := handler.New()
	log.Printf("Server started on %s...", port)
	if err := handler.Run(port); err != nil {
		log.Println(err)
	}

}
