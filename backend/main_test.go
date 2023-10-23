package main

import (
	"testing"
	"github.com/kbergstrom78/Swatchify/backend/database"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
	}
	_ = db
}
