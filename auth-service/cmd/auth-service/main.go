package main

import (
	"fmt"
	"log"

	"github.com/example/micro/auth-service/internal/app"
)

func main() {
	srv, err := app.New()
	if err != nil {
		log.Fatalf("init failed: %v", err)
	}
	fmt.Println("auth-service running on :8082")
	srv.Run()
}
