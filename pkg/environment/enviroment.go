package environment

import (
	"fmt"
	"log"
	"os"

	"github.com/NikiTesla/link_shortener/pkg/repository"
)

type Environment struct {
	Host string
	Port string
	DB   repository.Repo
}

// NewEnvironment creates environment according to configuration file.
// Stores port, host and database abstraction
func NewEnvironment() (*Environment, error) {
	log.Println("Setting environment")

	var db repository.Repo
	var err error

	if os.Getenv("DBTYPE") == "in-memory" {
		db, err = NewInMemoryDataBase()
	} else if os.Getenv("DBTYPE") == "postgres" {
		db, err = NewPostgersDataBase()
	} else {
		log.Fatalf("Cannot define database type")
	}

	if err != nil {
		return nil, fmt.Errorf("can't create environment bacause of database: %s", err.Error())
	}

	log.Printf("Host is %s\n", os.Getenv("HOST"))
	log.Printf("Port is %s\n", os.Getenv("PORT"))
	log.Printf("Database config is %s\n", os.Getenv("POSTGRES_URL"))

	return &Environment{
		DB:   db,
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}, nil
}
