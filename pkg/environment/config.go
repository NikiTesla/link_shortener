package environment

import (
	"encoding/json"
	"log"
	"os"
)

// PostgresDBConfig is a struct for database configuration parse
type PostgresDBConfig struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Username string `json:"username"`
	DBname   string `json:"dbname"`
	SSlmode  string `json:"sslmode"`
}

// Config is a struct for general configuration parse
type Config struct {
	Port             int              `json:"port"`
	RestPort         int              `json:"rest-port"`
	Host             string           `json:"host"`
	DBType           string           `json:"dbtype"`
	PostgresDBConfig PostgresDBConfig `json:"db-config"`
}

// NewConfig unmarshal configuration from configFile
func NewConfig(configFile string) (*Config, error) {
	rawJSON, err := os.ReadFile(configFile)
	if err != nil {
		log.Printf("Can't read config file, err: %s\n", err.Error())
		return nil, err
	}

	var config Config
	err = json.Unmarshal(rawJSON, &config)
	if err != nil {
		log.Printf("Can't unmarshall config json, err: %s\n", err.Error())
		return nil, err
	}

	return &config, nil
}
