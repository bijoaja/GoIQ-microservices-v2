package main

import (
	"fmt"
	"log"

	"github.com/example/micro/gateway/internal/app"
)

func main() {
	srv, err := app.New()
	if err != nil {
		log.Fatalf("init failed: %v", err)
	}
	fmt.Println("gateway running on :8080")
	srv.Run()
}
