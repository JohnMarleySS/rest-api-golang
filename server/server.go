package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JohnMarleySS/rest-api-golang/prisma/db"
	"github.com/JohnMarleySS/rest-api-golang/router"
)

func RunServer() {
	if err := run(); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}
}

func run() error {

	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return fmt.Errorf("failed to connect to Prisma: %v", err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			log.Printf("Failed to disconnect Prisma: %v", err)
		}
	}()

	router.RunRoutes(client)

	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	return nil
}
