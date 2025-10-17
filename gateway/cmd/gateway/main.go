package main

import (
	"fmt"
	"log"

	"github.com/bijoaja/GoIQ-microservices.v2/gateway/internal/app"
)

func main() {
	srv, err := app.New()
	if err != nil {
		log.Fatalf("init failed: %v", err)
	}
	fmt.Println("gateway running on :8080")
	srv.Run()
}
