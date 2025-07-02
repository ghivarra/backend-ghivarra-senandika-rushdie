package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {
	dsn := os.Getenv("DB_DSN")

	// get timeout connect
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(context, dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to the database. Reason: %v", err)
	}

	DB = pool
	return nil
}
