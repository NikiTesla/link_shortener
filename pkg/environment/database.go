package environment

import (
	"context"
	"log"
	"os"

	"github.com/NikiTesla/link_shortener/pkg/repository"
	"github.com/jackc/pgx"
)

// NewPostgersDataBase creates connection to postgres databse in case of using it
func NewPostgersDataBase(cfg PostgresDBConfig) (*repository.PostgresDB, error) {
	config := pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     uint16(cfg.Port),
		Database: cfg.DBname,
		User:     cfg.Username,
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	db, err := pgx.Connect(config)
	if err != nil {
		log.Fatalf("can't connect database, err: %s\n", err)
	}
	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("Cannot connect to database, error: %s\n", err)
	}

	return &repository.PostgresDB{
		DB: db,
	}, nil
}

// NewInMemoryDataBase creates in-memory database in case of using it
func NewInMemoryDataBase() (*repository.InMemoryDB, error) {
	return &repository.InMemoryDB{DB: make(map[string]string)}, nil
}
