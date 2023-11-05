package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestDbConnection(t *testing.T) {
	t.Run("should connect to database successfully", func(t *testing.T) {
		driver, dsn, err := setupDbTest()
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}

		db, err := sqlx.Open(driver, dsn)
		if err != nil {
			t.Fatalf("Failed to connect to the database. %v", err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			t.Fatalf("Error when pinging the database: %v", err)
		}
	})
}

func setupDbTest() (string, string, error) {
	err := godotenv.Load("../.env.test")
	if err != nil {
		return "", "", err
	}

	password := os.Getenv("DB_TEST_PASS")
	driver := os.Getenv("DB_TEST_DRIVER")
	user := os.Getenv("DB_TEST_USER")
	name := os.Getenv("DB_TEST_NAME")
	host := os.Getenv("DB_TEST_HOST")
	port := os.Getenv("DB_TEST_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, name, port)

	return driver, dsn, nil
}
