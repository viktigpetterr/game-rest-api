package mysql

import (
	"github.com/joho/godotenv"
	"testing"
)

func TestOpen(t *testing.T) {
	conn := Connection{}
	err := godotenv.Load("../../../../.env.local")
	if err != nil {
		t.Error("failed to load dotenv")
	}
	db, err := conn.Open()
	if err != nil {
		t.Error("failed to open db session")
	}
	if db == nil {
		t.Error("failed to assert db instance")
	}
}

func TestMigrate(t *testing.T) {
	conn := Connection{}
	err := godotenv.Load("../../../../.env.local")
	if err != nil {
		t.Error("failed to load dotenv")
	}
	err = conn.Migrate()
	if err != nil {
		t.Error("failed to assert that Migration run without errors")
	}
}
