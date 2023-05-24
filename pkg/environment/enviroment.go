package environment

import (
	"fmt"
	"log"
	"os"

	"github.com/NikiTesla/link_shortener/pkg/repository"
)

type Environment struct {
	Config *Config
	DB     repository.Repo
}

// NewEnvironment creates environment according to configuration file.
// Stores port, host and database abstraction
func NewEnvironment(configFile string) (*Environment, error) {
	log.Println("Setting environment")

	cfg, err := NewConfig(configFile)
	if err != nil {
		return nil, fmt.Errorf("can't create environment because of config: %s", err.Error())
	}

	var db repository.Repo
	if os.Getenv("DBTYPE") == "in-memory" {
		db, err = NewInMemoryDataBase()
	} else if os.Getenv("DBTYPE") == "postgres" {
		db, err = NewPostgersDataBase(cfg.PostgresDBConfig)
	} else {
		log.Fatalf("Cannot define database type")
	}

	if err != nil {
		return nil, fmt.Errorf("cant create environment bacause of database: %s", err.Error())
	}

	log.Printf("Host is %s\n", cfg.Host)
	log.Printf("Port is %d\n", cfg.Port)
	log.Printf("Database config is %+v\n", cfg.DBType)

	return &Environment{
		Config: cfg,
		DB:     db,
	}, nil
}
