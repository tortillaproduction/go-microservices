package main

import (
	"log"

	"github.com/tortillaproduction/go-microservices/apps/query-service/infra/gorm/handler"
)

func main() {
	log.Println("starting query-service...")

	db, err := handler.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	log.Println("connect to db successfully")
	_ = db

	select {}
}
