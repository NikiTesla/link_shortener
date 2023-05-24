package environment

import (
	"context"
	"log"
	"os"

	"github.com/NikiTesla/link_shortener/pkg/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPostgersDataBase creates connection to postgres databse in case of using it
func NewPostgersDataBase() (*repository.PostgresDB, error) {
	config, err := pgxpool.ParseConfig(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("cannot parse database config: %s", err)
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("can't connect to database, err: %s\n", err)
	}

	if err := dbpool.Ping(context.Background()); err != nil {
		log.Fatalf("can't ping database, error: %s\n", err)
	}

	return &repository.PostgresDB{
		DB: dbpool,
	}, nil
}

// NewInMemoryDataBase creates in-memory database in case of using it
func NewInMemoryDataBase() (*repository.InMemoryDB, error) {
	return &repository.InMemoryDB{DB: make(map[string]string)}, nil
}
