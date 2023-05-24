package main

import (
	"log"
	"os"

	"github.com/NikiTesla/link_shortener/pkg/environment"
	"github.com/NikiTesla/link_shortener/pkg/service"
	"google.golang.org/grpc"
)

func main() {
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
