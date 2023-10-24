package test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ericoliveiras/gate-guard/internal/config"
	"github.com/ericoliveiras/gate-guard/internal/db"
	_ "github.com/lib/pq"
)

func TestDB(t *testing.T) {
	t.Run("should connect successfully", func(t *testing.T) {
		testDB, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Error creating database mock: %v", err)
		}
		defer testDB.Close()

		config := &config.Config{
			DB: config.DBConfig{
				Driver:   "postgres",
				Host:     "localhost",
				User:     "myuser",
				Password: "mypassword",
				Name:     "mydb",
				Port:     "5432",
			},
		}

		dbx, err := db.Init(config)
		if err != nil {
			t.Fatalf("Error initializing the database: %v", err)
		}
		defer dbx.Close()

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unmet expectations: %s", err)
		}
	})
}
