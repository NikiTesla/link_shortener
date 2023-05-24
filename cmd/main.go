package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "github.com/NikiTesla/link_shortener/api"
	"github.com/NikiTesla/link_shortener/pkg/environment"
	"github.com/NikiTesla/link_shortener/pkg/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	// rest proxy running on port 8081
	go runRest(fmt.Sprintf("%s:%d", env.Config.Host, env.Config.Port), env.Config.RestPort)

	server := grpc.NewServer()
	shortenerServer := service.NewShortenerServer(env)
	if err := service.RunShortenerServer(shortenerServer, server); err != nil {
		log.Fatal(err)
	}
}

// runRest function runs rest service on port restPort that proxies calls to grpc service on grpcAddress
func runRest(grpcAddress string, restPort int) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterShortenerHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		panic(err)
	}

	log.Printf("rest server listening at %d", restPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", restPort), mux); err != nil {
		panic(err)
	}
}
