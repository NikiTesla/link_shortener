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
		log.Fatal("can't load dotenv variables, err:", err)
	}

	// configuring port, host and database
	env, err := environment.NewEnvironment()
	if err != nil {
		log.Fatal("can't load environment, err:", err)
	}

	// rest proxy running on port REST_PORT
	go runRest(fmt.Sprintf("%s:%s", env.Host, env.Port), os.Getenv("REST_PORT"))

	server := grpc.NewServer()
	shortenerServer := service.NewShortenerServer(env)
	if err := service.RunShortenerServer(shortenerServer, server); err != nil {
		log.Fatal(err)
	}
}

// runRest function runs rest service on port restPort that proxies calls to grpc service on grpcAddress
func runRest(grpcAddress string, restPort string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterShortenerHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		panic(err)
	}

	log.Printf("rest server listening at %s", restPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", restPort), mux); err != nil {
		panic(err)
	}
}
