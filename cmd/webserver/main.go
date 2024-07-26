package main

import (
	"log"
	"net/http"
	"receipt_processor/server"
	"receipt_processor/services"
)

func main() {
	inMemoryService, err := services.NewInMemoryReceiptService()

	if err != nil {
		log.Fatal(err)
	}

	s, err := server.NewReceiptServer(inMemoryService)

	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", s); err != nil {
		log.Fatalf("could not listen on port %d %v", "8080", err)
	}
}
