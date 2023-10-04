package connect

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"time"
)

func Connect(ctx context.Context) *pgxpool.Pool {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbUrl := os.Getenv("POSTGRES_URL")
	if dbUrl == "" {
		panic("missing environment variable POSTGRES_URL")
	}
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		panic(fmt.Sprintf("error configuring the database: %s", err))
	}
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic(fmt.Sprint("Unable to connect to database:", err))
	}
	var now time.Time
	err = pool.QueryRow(context.Background(), "SELECT NOW()").Scan(&now)
	if err != nil {
		panic(fmt.Sprint("failed to execute query", err))
	}
	log.Println("CONNECTION WITH DATABASE ESTABLISHED!")

	return pool
}
