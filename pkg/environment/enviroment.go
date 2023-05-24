package environment

import (
	"fmt"
	"log"

	"github.com/NikiTesla/link_shortener/pkg/repository"
)

type Environment struct {
	Config *Config
	DB     repository.Repo
}

func NewEnvironment(configFile string) (*Environment, error) {
	log.Println("Setting environment")

	cfg, err := NewConfig(configFile)
	if err != nil {
		return nil, fmt.Errorf("can't create environment because of config: %s", err.Error())
	}

	var db repository.Repo
	if cfg.DBType == "in-memory" {
		db, err = NewInMemoryDataBase()
	} else if cfg.DBType == "postgres" {
		db, err = NewPostgersDataBase(cfg.PostgresDBConfig)
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
