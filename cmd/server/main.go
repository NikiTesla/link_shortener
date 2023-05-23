package main

import (
	"log"

	"github.com/NikiTesla/link_shortener/pkg/service"
	"google.golang.org/grpc"
)

const (
	port = 50051
)

func main() {
	// TODO set environment

	server := grpc.NewServer()
	if err := service.RunShortenerServer(port, server); err != nil {
		log.Fatal(err)
	}
}
