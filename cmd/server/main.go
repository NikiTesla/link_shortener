package main

import (
	"log"
	"os"

	"github.com/NikiTesla/link_shortener/pkg/environment"
	"github.com/NikiTesla/link_shortener/pkg/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// configuring passwords, configs source and data storage type
	if err := godotenv.Load(); err != nil {
		log.Fatal("can't load env variables, err:", err)
	}

	// configuring port, host and database
	configFile := os.Getenv("CONFIGFILE")
	env, err := environment.NewEnvironment(configFile)
	if err != nil {
		log.Fatal("can't load environment, err:", err)
	}

	server := grpc.NewServer()
	shortenerServer := service.NewShortenerServer(env)
	if err := service.RunShortenerServer(shortenerServer, server); err != nil {
		log.Fatal(err)
	}
}
