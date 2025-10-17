package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bijoaja/GoIQ-microservices.v2/user-service/internal/app"
)

func main() {
	migrate := flag.Bool("migrate", false, "run migrations")
	flag.Parse()

	srv, err := app.New()
	if err != nil {
		log.Fatalf("failed init app: %v", err)
	}

	if *migrate {
		if err := srv.Migrate(); err != nil {
			log.Fatalf("migrate failed: %v", err)
		}
		fmt.Println("migrations done")
		return
	}

	srv.Run()
}
