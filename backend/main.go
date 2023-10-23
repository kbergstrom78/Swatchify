package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/kbergstrom78/Swatchify/backend/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	_ = db

	r := gin.Default()
	// set up routes
	r.Run(":8080")
}